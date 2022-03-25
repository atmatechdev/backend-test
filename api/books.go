package api

import (
	"fmt"
	"strconv"
	"technical-test-atmatech/database"
	"technical-test-atmatech/models"

	"github.com/gofiber/fiber/v2"
)

type BookInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
}

type DeleteBookRes struct {
	Deleted    bool   `json:"deleted"`
	DeleteType string `json:"delete_type"`
}

func checkBookInput(book models.Book) (bool, string) {
	if book.Title == "" {
		return false, "INVALID INPUT: Book's Title cannot be empty"
	} else if book.Content == "" {
		return false, "INVALID INPUT: Book's Content cannot be empty"
	} else if book.Description == "" {
		return false, "INVALID INPUT: Book's Description cannot be empty"
	} else {
		return true, ""
	}
}

func GetBooks(c *fiber.Ctx) error {
	//  Obtain query strings, determine offset
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page < 1 {
		page = 1
	}
	offset := int(limit * (page - 1))

	// Get books
	books := []models.Book{}
	database.DB.Limit(limit).Offset(offset).Find(&books)
	// fmt.Println("GetBooks: ", books)
	res := SuccessResponse(books, "GET ALL BOOKS")
	return c.Status(200).JSON(res)
}

func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book models.Book

	database.DB.Find(&book, "id = ?", id)
	if book.Title == "" {
		errRes := ErrorResponse(404, "Book with given ID NOT FOUND")
		return c.Status(404).JSON(errRes)
	}

	res := SuccessResponse(book, "GET a book successful")
	return c.Status(200).JSON(res)
}

func CreateBook(c *fiber.Ctx) error {
	// Obtain user data from middleware
	var user models.User = c.Locals("user").(models.User)
	fmt.Println("User Data from Middleware: ", user)

	var book models.Book
	if err := c.BodyParser(&book); err != nil {
		errRes := ErrorResponse(400, "Book input format is invalid")
		return c.Status(400).JSON(errRes)
	}

	ok, err := checkBookInput(book)
	if !ok {
		errRes := ErrorResponse(400, err)
		return c.Status(400).JSON(errRes)
	}

	// Customize the CreatedById by data from middleware
	book.CreatedById = user.ID

	// Create data to DB
	database.DB.Create(&book)
	res := SuccessResponse(book, "CREATE book successful")
	return c.Status(200).JSON(res)
}

func UpdateBook(c *fiber.Ctx) error {
	// Obtain the book from DB
	id := c.Params("id")
	var book models.Book
	database.DB.Find(&book, "id = ?", id)
	if book.ID == 0 {
		errRes := ErrorResponse(404, "Book with given ID NOT FOUND")
		return c.Status(404).JSON(errRes)
	}

	// Parse request body to struct
	var input BookInput
	if err := c.BodyParser(&input); err != nil {
		errRes := ErrorResponse(400, "Requst body format for Updating Book is invalid")
		return c.Status(400).JSON(errRes)
	}

	// Updating book
	book.Title = input.Title
	book.Description = input.Description
	book.Content = input.Content
	ok, err := checkBookInput(book)
	if !ok {
		errRes := ErrorResponse(400, err)
		return c.Status(400).JSON(errRes)
	}
	database.DB.Save(&book)

	res := SuccessResponse(book, "UPDATE book successful")
	return c.Status(200).JSON(res)
}

func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book models.Book

	// Find book, if any
	database.DB.Find(&book, "id = ?", id)
	if book.ID == 0 {
		errRes := ErrorResponse(404, "Book not found")
		return c.Status(404).JSON(errRes)
	}

	// Delete book
	if err := database.DB.Delete(&book).Error; err != nil {
		// If failed to delete, then it's server's error
		errRes := ErrorResponse(500, "Failed to delete book")
		return c.Status(500).JSON(errRes)
	}

	deleteData := DeleteBookRes{
		Deleted:    true,
		DeleteType: "soft",
	}
	res := SuccessResponse(deleteData, "DELETE book successful")

	return c.Status(200).JSON(res)
}
