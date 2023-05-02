package controller

import (
	"books-app/dto"
	"books-app/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	repository.Repository
}

func (m *Controller) GetAllBooks(c echo.Context) error {
	fetchedBooks, _ := m.Repository.FindAll()
	var books []dto.Book
	for _, fetchedBook := range fetchedBooks {
		book := dto.Book{
			Isbn:   fetchedBook.Isbn,
			Title:  fetchedBook.Title,
			Author: fetchedBook.Author,
		}
		books = append(books, book)
	}

	return c.JSON(http.StatusOK, books)
}
