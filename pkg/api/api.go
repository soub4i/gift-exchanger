package api

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/soub4i/giftsxchanger/pkg/api/v1"
)

func Register(version string) *gin.Engine {
	r := gin.Default()
	router := r.Group(version)
	router.GET("/health", v1.Healthcheck)

	router.GET("/members", v1.Get)
	router.POST("/members", v1.Create)
	router.GET("/members/:id", v1.Fetch)
	router.PUT("/members/:id", v1.Update)
	router.DELETE("/members/:id", v1.Delete)

	router.GET("/gift-exchange", v1.GetExchange)
	router.POST("/gift-exchange", v1.Shuffle)

	return r
}
