package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp/v3"
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
	pp.Printf("ID is: ",id)
	book, err := bookById(id)
	pp.Printf("ERROR: ",err)
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

	pp.Println(books)
	return nil, errors.New("book not found")
}

func getBooks(c *gin.Context){
	c.IndentedJSON(http.StatusOK, books)
}

func createBook(c *gin.Context) {
	var newBook book
	fmt.Println("HERE:")	
	fmt.Println(newBook)
	if err := c.BindJSON(&newBook); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid book object."} )
		return
	}

	books = append(books, newBook)
	pp.Println(books)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func main(){
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", createBook)
	router.GET("/books/:id", getBookById)
	router.Run("localhost:3000")

}