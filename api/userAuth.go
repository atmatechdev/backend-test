package api

import (
	"log"
	"technical-test-atmatech/database"
	"technical-test-atmatech/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type AuthResData struct {
	IsAuthenticated bool   `json:"is_authenticated"`
	AccessToken     string `json:"access_token"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func matchPasswordHash(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func isUsernameExist(u models.User) bool {
	var foundUser models.User
	database.DB.Find(&foundUser, "username = ?", u.Username)
	return foundUser.ID != 0
}

func UserRegister(c *fiber.Ctx) error {
	// Parse User model from Ctx
	var register Login
	err := c.BodyParser(&register)
	if err != nil {
		log.Fatal(err)
		return c.Status(400).JSON(err.Error())
	}
	// fmt.Println("User from Register: ", register)
	if register.Username == "" || register.Password == "" {
		errRes := ErrorResponse(400, "Incorrect Authentication input")
		return c.Status(400).JSON(errRes)
	}

	user := models.User{
		Username: register.Username,
		Password: register.Password,
	}
	if usernameExist := isUsernameExist(user); usernameExist {
		errRes := ErrorResponse(400, "Username already exsts")
		return c.Status(400).JSON(errRes)
	}

	// HASH user password
	hash, err := hashPassword(user.Password)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	user.Password = hash
	database.DB.Save(&user)

	// Generate token for user auth
	token := GenerateUserToken(user)

	data := AuthResData{
		IsAuthenticated: true,
		AccessToken:     token,
	}
	registerResponse := SuccessResponse(data, "Successfully registered")

	return c.Status(200).JSON(registerResponse)
}

func UserLogin(c *fiber.Ctx) error {
	// Parse User model from Ctx
	var user models.User
	var login Login
	err := c.BodyParser(&login)
	if err != nil {
		errRes := ErrorResponse(400, "Login data is in invalid structure")
		return c.Status(400).JSON(errRes)
	}
	// Find user with given username
	database.DB.Find(&user, "username = ?", login.Username)
	if user.ID == 0 {
		errRes := ErrorResponse(403, "User not found")
		return c.Status(403).JSON(errRes)
	}

	// Match hashed password
	isPasswordMatch := matchPasswordHash(user.Password, login.Password)
	if !isPasswordMatch {
		errRes := ErrorResponse(403, "Password is wrong")
		return c.Status(403).JSON(errRes)
	}

	// Return token if successful
	token := GenerateUserToken(user)

	data := AuthResData{
		IsAuthenticated: true,
		AccessToken:     token,
	}
	res := SuccessResponse(data, "Successfully logged in")
	return c.Status(200).JSON(res)
}
