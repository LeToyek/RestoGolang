package helper

import (
	"errors"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var JwtKey = []byte("secret_key")

type Claims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}

func GenerateAllTokens(userID string) (string, time.Time) {
	expirationTime := time.Now().Add(time.Minute * 5)

	claims := &Claims{
		ID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)

	if err != nil {
		panic(err)
	}
	return tokenString, expirationTime
}
func ValidateToken(c *gin.Context) (*Claims, error) {
	cookie, err := c.Request.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			return nil, err
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return nil, err
	}
	tokenStr := cookie.Value
	claims := &Claims{}

	//token is invalid
	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return JwtKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, err
		}
		return nil, err
	}

	if !tkn.Valid {
		return nil, errors.New("token is invalid")
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, errors.New("token is expired")
	}

	return claims, err
}
