package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	. "gitlab.com/jeelabs/learnings/goex-rest-unit-test/entity"
	"gitlab.com/jeelabs/learnings/goex-rest-unit-test/helper"
	"gitlab.com/jeelabs/learnings/goex-rest-unit-test/service"
	"net/http"
)

type CompanyHandler interface {
	Get(ctx *gin.Context)
	ById(ctx *gin.Context)
	Add(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type companyHandler struct {
	service service.CompanyService
}

func (c *companyHandler) Get(ctx *gin.Context) {

	result, err := c.service.Get()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, helper.BodyResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    &result,
	})
}

func (c *companyHandler) ById(ctx *gin.Context) {
	id := ctx.Param("id")

	result, err := c.service.ById(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	ctx.JSON(http.StatusOK, helper.BodyResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    &result,
	})

}

func (c *companyHandler) Add(ctx *gin.Context) {
	var newCompany Company

	if err := ctx.ShouldBindJSON(&newCompany); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if newCompany.ID == "" {
		newCompany.ID = xid.New().String()
	}

	result, err := c.service.Add(newCompany)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, helper.BodyResponse{
		Code:    http.StatusCreated,
		Message: "Success Add",
		Data:    &result,
	})
}

func (c *companyHandler) Update(ctx *gin.Context) {
	id := ctx.Param("id")

	var company Company
	if err := ctx.ShouldBindJSON(&company); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := c.service.Update(id, company)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, helper.BodyResponse{
		Code:    http.StatusCreated,
		Message: "Success Update",
		Data:    &result,
	})
}

func (c *companyHandler) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	err := c.service.Delete(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, helper.BodyResponse{
		Code:    http.StatusOK,
		Message: "Success delete",
	})
}

func NewCompanyHandler(service service.CompanyService) CompanyHandler {
	return &companyHandler{service: service}
}

//func CompanyRouter(router *gin.Engine) {
//	h := NewCompanyHandler(service.NewCompanyService(repository.NewCompanyRepository()))
//	router.GET("/companies", h.Get)
//	router.GET("/companies/:id", h.ById)
//	router.POST("/companies", h.Add)
//	router.PUT("/companies/:id", h.Update)
//	router.DELETE("/companies/:id", h.Delete)
//}
