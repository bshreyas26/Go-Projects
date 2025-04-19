package main

import (
	"Go-Fiber-PostgreSQL/models"
	"Go-Fiber-PostgreSQL/storage"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

type Book struct {
	Author    string `json:"author"`
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
}

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) CreateBook(c *fiber.Ctx) error {
	var book Book
	if err := c.BodyParser(&book); err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{"message": "Request body is invalid"})
		return err
	}
	err := r.DB.Create(&book).Error
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Could not create book"})
		return err
	}
	c.Status(http.StatusCreated).JSON(fiber.Map{"message": "Book created successfully"})
	return nil
}

func (r *Repository) DeleteBook(c *fiber.Ctx) error {
	bookModel := &models.Book{}
	id := c.Params("id")
	if id == "" {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid book id"})
		return nil
	}
	err := r.DB.Delete(bookModel, id).Error
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Could not delete book"})
		return err
	}
	c.Status(http.StatusOK).JSON(fiber.Map{"message": "Book deleted successfully"})
	return nil
}

func (r *Repository) GetBookbyID(c *fiber.Ctx) error {
	id := c.Params("id")
	bookModel := &models.Book{}
	if id == "" {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid book id"})
		return nil
	}
	err := r.DB.Where("id = ?", id).First(bookModel).Error
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Could not get book"})
	}
	c.Status(http.StatusOK).JSON(&fiber.Map{"message": "Book fetched successfully", "data": bookModel, "status": "success"})
	return nil
}

func (r *Repository) GetBooks(c *fiber.Ctx) error {
	bookModels := &[]models.Book{}
	err := r.DB.Find(&bookModels).Error
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Could not get books"})
	}
	c.Status(http.StatusOK).JSON(&fiber.Map{"message": "Books fetched successfully", "data": bookModels, "status": "success"})
	return nil
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create_books", r.CreateBook)
	api.Delete("/delete_book/:id", r.DeleteBook)
	api.Get("/get_book/:id", r.GetBookbyID)
	api.Get("/books", r.GetBooks)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	config := &storage.Config{
		Host:    os.Getenv("DB_HOST"),
		Port:    os.Getenv("DB_PORT"),
		User:    os.Getenv("DB_USER"),
		Pass:    os.Getenv("DB_PASS"),
		DBName:  os.Getenv("DB_NAME"),
		SSLMode: os.Getenv("DB_SSL_MODE"),
	}

	db, err := storage.NewConnection(config)

	if err != nil {
		log.Fatal(err)
	}

	err = models.MigrateBooks(db)
	if err != nil {
		log.Fatal(err)
	}
	r := Repository{
		DB: db,
	}
	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":8080")
}
