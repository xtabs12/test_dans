package jwtx

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strings"
)

type JwtHmac256 struct {
}

func NewJwtHmac256() *JwtHmac256 {
	return &JwtHmac256{}
}

type PayloadToken struct {
	UserId     int    `json:"id"`
	Username   string `json:"username"`
	Name       string `json:"name"`
	TokenValue string `json:"token_value"`
}

func (j *JwtHmac256) ValidateTokenHmac256(r *http.Request) (payloadToken PayloadToken, err error) {
	auth := r.Header.Get("Authorization")
	splitAuth := strings.Split(auth, " ")
	if len(splitAuth) != 2 {
		return payloadToken, fmt.Errorf("header authorization is empty")
	} else if !strings.EqualFold(splitAuth[0], "Bearer") {
		return payloadToken, fmt.Errorf("header authorization is not Bearer")
	}
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(splitAuth[1], claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if !token.Valid {
		return payloadToken, errors.New("token not valid")
	}

	data, _ := json.Marshal(claims["data"])
	_ = json.Unmarshal(data, &payloadToken)
	payloadToken.TokenValue = auth
	return
}
