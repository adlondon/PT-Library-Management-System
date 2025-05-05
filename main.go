package main

import (
	"encoding/json"
	"io"
	"net/http"
	"slices"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

type PostBookBody struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	ISBN        string `json:"isbn"`
	Description string `json:"description"`
}

type Book struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	ISBN        string `json:"isbn"`
	Description string `json:"description"`
}

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8888", "http://localhost:5173"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	e.Static("/", "ui/dist")
	e.File("/", "ui/dist/index.html")

	books := []Book{
		{
			Title:       "The Great Gatsby",
			Author:      "F. Scott Fitzgerald",
			ISBN:        "9780743273565",
			Description: "A novel about the American Dream",
		},
		{
			Title:       "To Kill a Mockingbird",
			Author:      "Harper Lee",
			ISBN:        "9780446310789",
			Description: "A story of racial injustice and moral growth",
		},
	}

	e.GET("/books", func(c echo.Context) error {
		return c.JSON(202, books)
	})

	e.POST("/book", func(c echo.Context) error {
		r := c.Request()
		b, err := io.ReadAll(r.Body)
		if err != nil {
			log.Error("error in POST", err)
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid Body Request"})
		}

		newBook := PostBookBody{}
		err = json.Unmarshal(b, &newBook)
		if err != nil {
			log.Error(err)
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON"})
		}

		book := Book{
			Title:       newBook.Title,
			Author:      newBook.Author,
			ISBN:        newBook.ISBN,
			Description: newBook.Description,
		}

		books = append(books, book)

		return c.JSON(http.StatusCreated, book)
	})

	e.DELETE("/book/:isbn", func(c echo.Context) error {
		isbn := c.Param("isbn")
		books = removeBook(books, isbn)
		return c.JSON(http.StatusNoContent, books)
	})

	e.Logger.Fatal(e.Start(":8888"))
}

func removeBook(books []Book, isbn string) []Book {
	for i, book := range books {
		if book.ISBN == isbn {
			return slices.Delete(books, i, i+1)
		}
	}
	return books
}


