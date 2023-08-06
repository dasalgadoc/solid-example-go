package ping

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const pong = "pong"

type Ping struct{}

func (p *Ping) Pong(c *gin.Context) {
	c.String(http.StatusOK, pong)
}
