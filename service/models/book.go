package models

import (
 "fmt"
 "validator/config"
)

type Book struct {
 ID     uint   `gorm:"primaryKey"`
 Title  string `gorm:"size:255"`
 Author string `gorm:"size:100"`
 Year   int
 ISBN   string `gorm:"size:20"`
}

// Create a new book record
func CreateBook(title, author string, year int, isbn string) {
 book := Book{Title: title, Author: author, Year: year, ISBN: isbn}
 if err := config.DB.Create(&book).Error; err != nil {
  fmt.Println("Failed to create book:", err)
  return
 }
 fmt.Println("Book created successfully:", book)
}

// Get all books
func GetBooks() {
 var books []Book
 if err := config.DB.Find(&books).Error; err != nil {
  fmt.Println("Failed to retrieve books:", err)
  return
 }
 fmt.Println("Books:", books)
}

// Get a book by ID
func GetBookByID(id uint) {
 var book Book
 if err := config.DB.First(&book, id).Error; err != nil {
  fmt.Println("Book not found:", err)
  return
 }
 fmt.Println("Book:", book)
}

// Update a book by ID
func UpdateBook(id uint, newTitle string) {
 var book Book
 if err := config.DB.First(&book, id).Error; err != nil {
  fmt.Println("Book not found:", err)
  return
 }
 book.Title = newTitle
 if err := config.DB.Save(&book).Error; err != nil {
  fmt.Println("Failed to update book:", err)
  return
 }
 fmt.Println("Book updated:", book)
}

// Delete a book by ID
func DeleteBook(id uint) {
 if err := config.DB.Delete(&Book{}, id).Error; err != nil {
  fmt.Println("Failed to delete book:", err)
  return
 }
 fmt.Println("Book deleted with ID:", id)
}