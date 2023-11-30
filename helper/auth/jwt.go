package auth

import (
	"context"
	"fmt"
	"grpctest/internal/app/model"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc/metadata"
)

type JwtCustomClaims struct {
	UserId string
	jwt.StandardClaims
}

const SECRET_KEY = "SCnJhUUVYTSmDx0ye293Rd90qdLdRRRo"

func GenToken(user model.User) (string, error) {
	claims := JwtCustomClaims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	result, err := token.SignedString([]byte(SECRET_KEY))

	if err != nil {
		return "", err
	}

	return result, nil
}

func GetUserIdFromContext(ctx context.Context) (*JwtCustomClaims, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("metadata is not provided")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return nil, fmt.Errorf("authorization token is not provided")
	}

	accessToken := values[0]
	accessToken = strings.TrimPrefix(accessToken, "Bearer ")
	token, err := jwt.ParseWithClaims(
		accessToken,
		&JwtCustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected token signing method")
			}
			
			return []byte(SECRET_KEY), nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}
	claims, ok := token.Claims.(*JwtCustomClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
