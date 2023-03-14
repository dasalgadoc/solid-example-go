package entrypoints

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"solid-example-go/application/book"
)

type HttpBookDeleter struct {
	bookDeleter book.BookDeleter
}

func NewHttpBookDeleter(bookDeleterUC book.BookDeleter) HttpBookDeleter {
	return HttpBookDeleter{
		bookDeleter: bookDeleterUC,
	}
}

func (hbd *HttpBookDeleter) DeleteBook(c *gin.Context) {
	var isbn string
	isbn = c.Query("isbn")
	if isbn == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	err := hbd.bookDeleter.Delete(isbn)
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
