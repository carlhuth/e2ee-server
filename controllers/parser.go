package controllers

import (
        jwt "github.com/dgrijalva/jwt-go"
        "github.com/mancabizjak/e2ee-server/core/authentication"
        "net/http"
	"log"
)

func ExtractTokenInfo(r *http.Request) (string, string, uint) {
 	authBackend := authentication.InitJWTAuthenticationBackend()
        tokenRequest, err := jwt.ParseFromRequest(r, func(token *jwt.Token) (interface{}, error) {
                return authBackend.PublicKey, nil
        })

        if err != nil {
                log.Println(err)
        }
        var uuid string
        uuid = tokenRequest.Claims["sub"].(string)
        var username string
        username = tokenRequest.Claims["name"].(string)
	accountId := tokenRequest.Claims["accountId"].(float64)
	return uuid, username, uint(accountId)
}

