package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/jamisonwilliams99/GO_techcompanies_api/models"
	"github.com/jamisonwilliams99/GO_techcompanies_api/storage"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

// struct containing Company data
type Company struct {
	Name           string `json:"name"`
	Industry       string `json:"industry"`
	Funding        string `json:"funding"`
	Employees      int    `json:"employees"`
	EmployeeGrowth string `json:"employeegrowth"`
	Revenue        string `json:"revenue"`
}

// struct wrapping the gorm database interface
// - used to interface with the API context handler functions
type Repository struct {
	DB *gorm.DB
}

// context handler for POST api call
func (r *Repository) CreateCompany(context *fiber.Ctx) error {
	company := Company{}

	err := context.BodyParser(&company)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}

	err = r.DB.Create(&company).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create company"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "company has been added"})
	return nil
}

// context handler for DELETE api call
func (r *Repository) DeleteCompany(context *fiber.Ctx) error {
	companyModel := models.Companies{}
	id := context.Params("id")

	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}

	err := r.DB.Delete(companyModel, id)

	if err.Error != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not delete company",
		})
		return err.Error
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "company deleted successfully",
	})
	return nil
}

// context handler for general GET api call
func (r *Repository) GetCompanies(context *fiber.Ctx) error {
	companyModels := &[]models.Companies{}

	err := r.DB.Find(companyModels).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get companies"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "companies fetched successfully",
		"data":    companyModels,
	})
	return nil
}

// context handler for GET (by id) api call
func (r *Repository) GetCompanyByID(context *fiber.Ctx) error {
	id := context.Params("id") // get company id from http request
	companyModel := &models.Companies{}

	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}

	err := r.DB.Where("id = ?", id).First(companyModel).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get the company"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "company id fetched successfully",
		"data":    companyModel,
	})
	return nil
}

// maps api url paths to corresponding context handler functions
func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create_company", r.CreateCompany)
	api.Delete("/delete_company/:id", r.DeleteCompany)
	api.Get("/get_company/:id", r.GetCompanyByID)
	api.Get("/tech_companies", r.GetCompanies)
}

func main() {
	err := godotenv.Load(".env") // load environment variables from .env file

	if err != nil {
		log.Fatal(err)
	}

	// load data base configuration struct with environment variables
	// from the .env file
	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}

	db, err := storage.NewConnection(config) // make database connection

	if err != nil {
		log.Fatal("could not load the database")
	}

	err = models.MigrateCompanies(db) // populate database instance with fields to match Company struct

	if err != nil {
		log.Fatal("could not load the database")
	}

	// initialize Repository instance to interface with api context handler functions
	// with the database
	r := Repository{
		DB: db,
	}

	app := fiber.New()  // initialize fiber instance
	r.SetupRoutes(app)  // maps context handler functions to the api paths
	app.Listen(":8080") // configure fiber instance to operate on port 8080

}
