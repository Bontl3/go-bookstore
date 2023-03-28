package models

import (
	"github.com/Bontl3/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

// an instance of gorm.DB which will be initialised in the init function
var db *gorm.DB

// Defining the fields of a book record in the database
type Book struct {
	// `gorm.Model` adds standard fields to the record such as `ID`, `CreatedAt` and `UpdatedAt`
	gorm.Model
	Name        string `gorm: ""json:"name"`
	Author      string `json:"name"`
	Publication string `json:"publication"`
}

// Function initialises the database connection and performs the
// neccesary connections
func init() {
	// calling `config.Connect()` to establish a database connection
	config.Connect()
	// Retrieve a reference to the database
	db = config.GetDB()
	// Creates the neccesary table in the database
	db.AutoMigrate(&Book{})
}

// Creates a new book record in the database
func (b *Book) CreateBook() *Book {
	// Creates new record
	db.NewRecord(b)
	// Inserts the record into the database
	db.Create(&b)
	return b
}

// Returns a list of all books in the database
func GetAllBooks() ([]Book, error) {
	// Declare a variable to hold the book records
	var books []Book
	// Use the 'Find' method of the 'db' object to retrieve all book records
	if err := db.Find(&books).Error; err != nil {
		// If an error occurs, return a nil slice and the error object
		return nil, err
	}
	// If the query is successful, return the book records slice and a nil error object
	return books, nil
}

// Returns a pointer to a book record in the database with the given ID
func GetBookById(Id int64) (*Book, *gorm.DB, error) {
	// Create a variable that wil be used to hold the book record that is retrieved from the database
	var book Book
	// Query the database and store the result in the 'db'variable
	db := db.Where("ID=?", Id).Find(&book)
	// Check for errors
	if err := db.Error; err != nil {
		return nil, db, db.Error
	}
	// Return the book record and a nil error object if the query was successful
	return &book, db, nil
}

func DeleteBook(ID int64) (*Book, error) {
	var book Book
	// Use 'First' method to retrieve the book record by ID
	if err := db.First(&book, ID).Error; err != nil {
		// If the book record was not found, return a nil pointer and the error object
		if gorm.IsRecordNotFoundError(err) {
			return nil, err
		}
		// Otherwise, return the error object
		return nil, err
	}
	// Delete the book record from the database
	if err := db.Delete(&book).Error; err != nil {
		return nil, err
	}
	// Return a pointer to the deleted book record and a nil error object
	return &book, nil
}
