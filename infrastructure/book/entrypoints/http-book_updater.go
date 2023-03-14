package entrypoints

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"solid-example-go/application/book"
	"solid-example-go/infrastructure/book/dto"
)

type HttpBookUpdater struct {
	bookUpdater book.BookUpdater
}

func NewBookUpdater(bookUpdaterUC book.BookUpdater) HttpBookUpdater {
	return HttpBookUpdater{
		bookUpdater: bookUpdaterUC,
	}
}

func (hbu *HttpBookUpdater) UpdateBook(c *gin.Context) {
	var (
		err     error
		request dto.Book
	)
	if err = c.BindJSON(&request); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	err = hbu.bookUpdater.Update(request.ISBN, request.Author, request.Title)
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
