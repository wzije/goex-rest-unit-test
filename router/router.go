package router

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/jeelabs/learnings/goex-rest-unit-test/handler"
	"gitlab.com/jeelabs/learnings/goex-rest-unit-test/repository"
	"gitlab.com/jeelabs/learnings/goex-rest-unit-test/service"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	return router
}

func RegisterRouter() *gin.Engine {
	router := SetupRouter()
	v1 := router.Group("/api/v1")
	h := handler.NewCompanyHandler(service.NewCompanyService(repository.NewCompanyRepository()))
	v1.GET("/companies", h.Get)
	v1.GET("/companies/:id", h.ById)
	v1.POST("/companies", h.Add)
	v1.PUT("/companies/:id", h.Update)
	v1.DELETE("/companies/:id", h.Delete)
	return router
}
