package book

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"solid-example-go/application/book"
)

const HEADER = "text/html; charset=utf-8"

type BookPrinterController struct {
	htmlPrinter book.HtmlBookPrinter
}

func (bpc *BookPrinterController) PrintBook(c *gin.Context) {
	title := c.Query("title")
	author := c.Query("author")
	response := bpc.htmlPrinter.ServeBook(title, author)

	c.Data(http.StatusOK, HEADER, []byte(response))
}
