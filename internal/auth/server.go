package auth

import (
	"context"
	"time"

	"github.com/NavroO/adhub/proto/authpb"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("supersecret")

type AuthServer struct {
	authpb.UnimplementedAuthServiceServer
}

func New() *AuthServer {
	return &AuthServer{}
}

func (s *AuthServer) ValidateToken(ctx context.Context, req *authpb.ValidateTokenRequest) (*authpb.ValidateTokenResponse, error) {
	token, err := jwt.Parse(req.Token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrInvalidKey
		}
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		return &authpb.ValidateTokenResponse{Valid: false, Error: "invalid token"}, nil
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return &authpb.ValidateTokenResponse{Valid: false, Error: "invalid claims"}, nil
	}

	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		return &authpb.ValidateTokenResponse{Valid: false, Error: "token expired"}, nil
	}

	return &authpb.ValidateTokenResponse{
		Valid:  true,
		UserId: claims["sub"].(string),
	}, nil
}
