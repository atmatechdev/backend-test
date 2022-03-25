package api

import (
	"errors"
	"fmt"
	"log"
	"os"
	"technical-test-atmatech/models"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

type customClaims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

type TokenData struct {
	UserId uint      `json:"user_id"`
	Exp    time.Time `json:"exp"`
}

var tokenIssuer = "Atmatech"
var tokenDuration = int64(24 * 3600) // 24 hours

func getTokenSecret() string {
	// err := godotenv.Load(".env") // LOCAL DEVELOPMENT
	err := godotenv.Load("public.env") // FOR PRESENTATION / PUBLIC REPO
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("TOKEN_SECRET")
}

func GenerateUserToken(u models.User) (signedToken string) {
	now := time.Now().Unix()
	secret := getTokenSecret()
	expAt := now + int64(tokenDuration)
	claims := customClaims{
		UserID: u.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expAt,
			Issuer:    tokenIssuer,
		},
	}

	// Token itself
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// The signed token
	signed, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Fatal("Error while generating JWT: ", err)
		return
	}
	// fmt.Println("Signed Token: ", signed)
	return signed
}

func ParseUserToken(bearerToken string) (*customClaims, error) {
	var errMsg string
	secret := getTokenSecret()
	token, err := jwt.ParseWithClaims(
		bearerToken,
		&customClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		},
	)
	if err != nil {
		errMsg = "JWT-Error-Validating token"
		fmt.Println(errMsg, err)
		return &customClaims{}, errors.New(errMsg)
	}

	claims, ok := token.Claims.(*customClaims)
	if !ok {
		errMsg = "JWT-Error-Parsing claims"
		fmt.Println(errMsg, err)
		return &customClaims{}, errors.New(errMsg)
	}

	if claims.ExpiresAt < time.Now().UTC().Unix() {
		errMsg = "JWT-Error-Expired"
		return &customClaims{}, errors.New(errMsg)
	}
	// fmt.Println("Claims: ", claims)
	return claims, nil
}
