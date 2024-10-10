package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/soub4i/giftsxchanger/pkg/datastore"
)

// postAlbums adds an album from JSON received in the request body.
func Create(c *gin.Context) {
	var new datastore.Member
	ds := datastore.GetDS()

	if err := c.BindJSON(&new); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "id not found", "error": err.Error()})
		return
	}
	new.ID = strconv.Itoa(len(ds.Members))
	m := ds.AddMember(new)
	c.IndentedJSON(http.StatusCreated, m)
}

func Update(c *gin.Context) {
	var new datastore.Member
	id := c.Param("id")
	ds := datastore.GetDS()
	if id == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "id not found"})
		return
	}

	m := ds.GetMember(id)
	if m == nil {
		c.IndentedJSON(http.StatusNotFound, nil)
		return
	}

	if err := c.BindJSON(&new); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "id not found", "error": err.Error()})
		return
	}
	new.ID = id

	res := ds.UpdateMember(new)
	c.IndentedJSON(http.StatusOK, res)
}

func Fetch(c *gin.Context) {
	id := c.Param("id")
	ds := datastore.GetDS()

	if id == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "id not found"})
		return
	}

	m := ds.GetMember(id)
	if m != nil {
		c.IndentedJSON(http.StatusOK, m)
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "member not found"})
}

func Get(c *gin.Context) {
	ds := datastore.GetDS()
	m := ds.GetMembers()
	c.IndentedJSON(http.StatusOK, m)
}

func Delete(c *gin.Context) {
	id := c.Param("id")
	ds := datastore.GetDS()
	if id == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "id not found"})
		return
	}

	m := ds.GetMember(id)
	if m == nil {
		c.IndentedJSON(http.StatusNotFound, m)
		return
	}

	ds.DeleteMember(id)
	c.IndentedJSON(http.StatusNoContent, nil)
}
