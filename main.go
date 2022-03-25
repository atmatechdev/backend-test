package main

import (
	"log"
	"technical-test-atmatech/api"
	"technical-test-atmatech/database"

	"github.com/gofiber/fiber/v2"
)

func testServer(c *fiber.Ctx) error {
	res := api.Response{
		Status:  200,
		Message: "Server Running",
		Data:    nil,
	}
	// fmt.Println(res)
	return c.JSON(res)
}

func bookRoutes(app *fiber.App) {

	books := app.Group("/books", api.CheckUserAuth)
	books.Get("/", api.GetBooks)
	books.Get("/:id", api.GetBook)
	books.Post("/", api.CreateBook)
	books.Put("/:id", api.UpdateBook)
	books.Delete("/:id", api.DeleteBook)
}

func authRoutes(app *fiber.App) {
	app.Post("/auth/login", api.UserLogin)
	app.Post("/auth/register", api.UserRegister)
}

func main() {
	app := fiber.New()
	app.Get("/", testServer)

	bookRoutes(app)
	authRoutes(app)

	// Connect to Database
	db, err := database.DBConnect()
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	pgDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer pgDB.Close()

	app.Listen(":3000")
}
