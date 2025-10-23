package main

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/NavroO/adhub/internal/ads"
	"github.com/NavroO/adhub/internal/shared"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func main() {
	shared.SetupLogger()
	log.Info().Msg("📦 Logger initialized")

	if err := godotenv.Load(".env"); err != nil {
		log.Warn().Msg("⚠️ .env file not found, using system environment")
	}
	cfg := shared.LoadConfig()
	if cfg.Port == "" {
		log.Fatal().Msg("❌ PORT is not set in .env")
	}

	db, err := shared.ConnectDB()
	if err != nil {
		log.Fatal().Msgf("cannot connect to db: %v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Error().Err(err).Msg("failed to close database")
		}
	}()

	r := chi.NewRouter()
	r.Use(shared.RequestLogger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   cfg.CorsOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
	}))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		if err := db.Ping(); err != nil {
			http.Error(w, "DB not reachable", http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(map[string]string{"status": "ok"}); err != nil {
			log.Error().Err(err).Msg("failed to encode health response")
		}
	})

	repo := ads.NewRepository(db)
	svc := ads.NewService(repo)
	h := ads.NewHandler(svc)

	r.Mount("/ads", h.Routes())

	srv := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Info().Msgf("🚀 Starting HTTP server on :%s", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Msgf("❌ Server failed: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Info().Msg("🧹 Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal().Msgf("❌ Server forced to shutdown: %v", err)
	}

	log.Info().Msg("✅ Server exited gracefully")
}
