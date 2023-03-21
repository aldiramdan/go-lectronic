package libs

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var mySecret = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	UserID string
	Role string
	jwt.RegisteredClaims
}

func NewToken(id string, role string) *Claims {
	return &Claims{
		UserID: id,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 20)),},
	}	
}

func (c *Claims) CreateToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(mySecret)
}

func CheckToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(mySecret), nil
	})

	if err != nil {
		return nil, err
	}

	claims := token.Claims.(*Claims)
	return claims, nil
}