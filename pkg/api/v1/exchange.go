package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/soub4i/giftsxchanger/pkg/datastore"
)

func GetExchange(c *gin.Context) {
	ds := datastore.GetDS()
	m := ds.GetExchanges()
	c.IndentedJSON(http.StatusOK, m)
}

func Shuffle(c *gin.Context) {
	ds := datastore.GetDS()

	if err := ds.AssignRecipients(); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, ds.GetExchanges())
}
