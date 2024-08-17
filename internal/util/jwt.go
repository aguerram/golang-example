package util

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"mostafa/learn_go/internal/model"
	"strconv"
	"time"
)

func GenerateUserJWT(user *model.User, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": strconv.Itoa(int(user.Id)),
		"exp": time.Now().Add(time.Second * 10).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	return token.SignedString([]byte(secret))
}

func ValidateUserJWT(tokenString string, secret string) (uint32, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		expClaim := claims["exp"]
		//check if claim is not expired
		if exp, ok := expClaim.(float64); ok {
			if int64(exp) < time.Now().Unix() {
				return 0, fmt.Errorf("token is expired")
			}
		} else {
			return 0, fmt.Errorf("exp claim is not valid")
		}

		if userId, err := strconv.ParseInt(claims["sub"].(string), 10, 32); err != nil {
			return 0, err
		} else {
			return uint32(userId), nil
		}
	} else {
		return 0, err
	}
}
