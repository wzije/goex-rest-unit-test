package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"net/http"
)

func WelcomeHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "welcome to simple go crud"})
}

func GetCompanies(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Companies)
}

func CreateCompany(ctx *gin.Context) {
	var newCompany Company

	if err := ctx.ShouldBindJSON(&newCompany); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if newCompany.ID == "" {
		newCompany.ID = xid.New().String()
	}
	Companies = append(Companies, newCompany)
	ctx.JSON(http.StatusCreated, newCompany)
}

func UpdateCompany(ctx *gin.Context) {
	id := ctx.Param("id")

	var company Company
	if err := ctx.ShouldBindJSON(&company); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	index := -1

	for i := 0; i < len(Companies); i++ {
		if Companies[i].ID == id {
			index = i
		}
	}

	if index == -1 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "data not found"})
		return
	}

	Companies[index] = company
	ctx.JSON(http.StatusCreated, Companies[index])
}

func DeleteCompany(ctx *gin.Context) {
	id := ctx.Param("id")

	index := -1

	for i := 0; i < len(Companies); i++ {
		if Companies[i].ID == id {
			index = i
		}
	}

	if index == -1 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "data not found"})
		return
	}

	Companies = append(Companies[:index], Companies[index+1:]...)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Company has been deleted",
	})
}

func main() {
	router := SetUpRouter()
	router.GET("/", WelcomeHandler)
	router.GET("/companies", GetCompanies)
	router.POST("/companies", CreateCompany)
	router.PUT("/companies/:id", UpdateCompany)
	router.DELETE("/companies/:id", DeleteCompany)
	err := router.Run()
	if err != nil {
		panic(err)
	}
}

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}
