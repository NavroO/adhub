package shared

import (
	"os"
	"strings"
)

type Config struct {
	Port        string
	CorsOrigins []string
	LogPayloads bool
}

func LoadConfig() Config {
	origins := os.Getenv("CORS_ORIGINS")
	corsList := []string{"*"}
	if origins != "" {
		corsList = strings.Split(origins, ",")
	}

	return Config{
		Port:        os.Getenv("PORT"),
		CorsOrigins: corsList,
		LogPayloads: os.Getenv("LOG_PAYLOADS") == "true",
	}
}
