package auth

import (
	"fmt"
	"os"
	"time"

	user_migrations "github.com/YasserRABIE/QUIZFYv2/migrations/account_migrations"
	"github.com/YasserRABIE/QUIZFYv2/models/user"
	"github.com/golang-jwt/jwt/v5"
)

func ValidateToken(tokenString string) (*user.Account, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		return nil, fmt.Errorf("token expired")
	}

	id := claims["id"].(uint)
	user, err := user_migrations.GetById(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
