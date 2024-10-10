package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Healthcheck(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"healthy": true})
}
