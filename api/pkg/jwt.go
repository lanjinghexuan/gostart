package pkg

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
)

func GetToken(userId int32) (string, error) {
	hmacSampleSecret := []byte("9999")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(hmacSampleSecret)

	return tokenString, err
}

func VieyToken(tokenString string) (int32, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		hmacSampleSecret := []byte("9999")
		return hmacSampleSecret, nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		log.Fatal(err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// 使用类型断言和类型转换来获取 userId
		switch v := claims["userId"].(type) {
		case float64:
			return int32(v), nil // 将 float64 转换为 int32
		case int:
			return int32(v), nil // 将 int 转换为 int32
		default:
			return 0, fmt.Errorf("claims['userId'] is not a number type")
		}
	} else {
		return 0, fmt.Errorf("invalid claims type")
	}
}
