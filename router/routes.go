package router

import (
	"github.com/eovinicius/go-jobs/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {

	handler.InitializeHandler()
	basePath := "/api/v1"
	v1 := router.Group(basePath)

	v1.POST("/opening", handler.CreateOpeningHandler)

	v1.GET("/opening", handler.ShowOpeningHandler)

	v1.DELETE("/opening", handler.DeleteOpeningHandler)

	v1.PUT("/opening", handler.UpdateOpeningHandler)

	v1.GET("/openings", handler.ListOpeningHandler)

}
