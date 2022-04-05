package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)


type book struct {
	ID			string  `json:"id"`
	Title		string  `json:"title"`
	Author 		string  `json:"author"`
	Quantity	uint    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

func getBookById(c *gin.Context){
	id := c.Param("id")
	book, err := bookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No book found with that ID"} )
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}


func bookById(id string) (*book, error ) {
	for i,b := range books{
		if b.ID == id {
			return &books[i], nil
		}
	}

	return nil, errors.New("book not found")
}

func getBooks(c *gin.Context){
	c.IndentedJSON(http.StatusOK, books)
}

func checkoutBook (c *gin.Context) {
	id, ok := c.GetQuery("id")
	
	if !ok {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Invalid query"} )
		return
	}

	newBook, err := bookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No book with that ID"})
		return 
	}

	if newBook.Quantity < 1 {
		c.IndentedJSON(http.StatusConflict, gin.H{"message": "This book is not in stock."})
		return
	}
	newBook.Quantity -= 1

	c.IndentedJSON(http.StatusOK, newBook)

}

func returnBook(c *gin.Context) {
	id, ok := c.GetQuery("id")
	
	if !ok {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Invalid query"} )
		return
	}

	newBook, err := bookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No book with that ID"})
		return 
	}

	newBook.Quantity += 1

	c.IndentedJSON(http.StatusOK, newBook)

}

func createBook(c *gin.Context) {
	var newBook book
	if err := c.BindJSON(&newBook); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid book object."} )
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func main(){
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", createBook)
	router.GET("/books/:id", getBookById)
	router.PATCH("/checkout", checkoutBook)
	router.PATCH("/return", returnBook)
	router.Run("localhost:3000")

}