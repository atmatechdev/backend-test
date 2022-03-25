package api

import (
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func tokenInitialValidation(token string) (string, error) {
	// Check whether it includes "Bearer "
	// Space between, and no more spaces
	tokenSlice := strings.Split(token, " ")
	check := len(tokenSlice) == 2 && tokenSlice[0] == "Bearer"
	if !check {
		return "", errors.New("invalid token format was provided")
	}
	return tokenSlice[1], nil
}

func CheckUserAuth(c *fiber.Ctx) error {
	// Get bearer token from "Authorization"
	bearer := c.GetReqHeaders()["Authorization"]
	if bearer == "" {
		errRes := ErrorResponse(401, "UNAUTHORIZED: Please provide an auth token")
		return c.Status(401).JSON(errRes)
	}

	// Validate - Initial
	token, err := tokenInitialValidation(bearer)
	if err != nil {
		errRes := ErrorResponse(401, "UNAUTHORIZED: Token invalid")
		return c.Status(401).JSON(errRes)
	}

	// Parse to get claims
	claims, err := ParseUserToken(token)
	if err != nil {
		errRes := ErrorResponse(401, "UNAUTHORIZED: Token invalid")
		return c.Status(401).JSON(errRes)
	}
	// fmt.Println("Claims from check auth: ", claims)

	// If there is claims, then get user data into context
	user, err := GetUserById(claims.UserID)
	if err != nil {
		errRes := ErrorResponse(404, "you are not authorized - user not found")
		return c.Status(404).JSON(errRes)
	}
	c.Locals("user", user)

	return c.Next()
}
