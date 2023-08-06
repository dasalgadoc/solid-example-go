package entrypoints

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"solid-example-go/internal/application/book"
	"solid-example-go/internal/infrastructure/book/dto"
)

type HttpBookCreator struct {
	bookCreator book.BookCreator
}

func NewHttpBookCreator(bookCreatorUC book.BookCreator) HttpBookCreator {
	return HttpBookCreator{
		bookCreator: bookCreatorUC,
	}
}

func (hbc *HttpBookCreator) CreateBook(c *gin.Context) {
	var (
		err     error
		request dto.Book
	)
	if err = c.BindJSON(&request); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if !request.HasISBN() {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = hbc.bookCreator.Create(request.ISBN, request.Author, request.Title)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
