package middleware

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
	envConfig "zhaoxin-api/config"
)

// JWT Header Payload Signature
// JWT also obeys the OAuth2.0 Protocol
// When we request the server the http header should have the 'Bearer ${tokenValue}'

type TokenPayload struct {
	ID uint
}

func Generate(payload *TokenPayload) (string,error) {
	tDuration,err := time.ParseDuration(envConfig.Config("JWT_TOKEN_EXP"))
	if err != nil {
		return "" ,err
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"expiration": time.Now().Add(tDuration).Unix(),
		"ID": payload.ID,
	})
	tokenValue,err := jwtToken.SignedString([]byte(envConfig.Config("JWT_TOKEN_KEY")))
	if err != nil {
		return "",err
	}
	return tokenValue,nil
}

func parse(tokenValue string) (*jwt.Token,error) {
	return jwt.Parse(tokenValue,func(jwtToken *jwt.Token) (interface{},error) {
		if _,ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil,fmt.Errorf("Unexpected signing method: %v ",jwtToken.Header["alg"])
		}
		return []byte(envConfig.Config("JWT_TOKEN_KEY")),nil
	})
}

func Verify(tokenValue string) (*TokenPayload,error) {
	parsedValue, err := parse(tokenValue)
	if err != nil {
		return nil, err
	}
	tokenClaims, ok := parsedValue.Claims.(jwt.MapClaims)
	if !ok {
		return nil, err
	}
	id,ok := tokenClaims["ID"].(float64)
	if !ok {
		return nil, errors.New("Something errors ")
	}
	return &TokenPayload{ID: uint(id)},nil
}