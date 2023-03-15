package entrypoints

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"solid-example-go/application/book"
	"solid-example-go/infrastructure/book/dto"
)

type HttpBookFinder struct {
	bookFinder book.BookGetter
}

func NewHttpBookFinder(bookGetterUC book.BookGetter) HttpBookFinder {
	return HttpBookFinder{
		bookFinder: bookGetterUC,
	}
}

func (hbf *HttpBookFinder) FindBook(c *gin.Context) {
	isbn := c.Query("isbn")
	if isbn == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	book, err := hbf.bookFinder.Find(isbn)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			})
		return
	}
	var response dto.Book
	response.ISBN = book.ISBN
	response.Author = book.Author
	response.Title = book.Title
	c.JSON(http.StatusOK, response)
}
