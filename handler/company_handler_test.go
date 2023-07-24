package handler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/rs/xid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gitlab.com/jeelabs/learnings/goex-rest-unit-test/entity"
	. "gitlab.com/jeelabs/learnings/goex-rest-unit-test/handler"
	"gitlab.com/jeelabs/learnings/goex-rest-unit-test/helper"
	routers "gitlab.com/jeelabs/learnings/goex-rest-unit-test/router"
	"gitlab.com/jeelabs/learnings/goex-rest-unit-test/service"
	"net/http"
	"net/http/httptest"
	"testing"
)

var companyService = service.CompanyServiceMock{Mock: mock.Mock{}}

var handler = NewCompanyHandler(&companyService)

var router = routers.SetupRouter()

func Test_companyHandler_ById(t *testing.T) {
	router.GET("/companies/:id", handler.ById)

	t.Run("get company by id should success", func(t *testing.T) {
		var company = entity.Companies[0]
		companyService.Mock.On("ById", "1").Return(company, nil)

		request, _ := http.NewRequest(http.MethodGet, "/companies/1", nil)
		writer := httptest.NewRecorder()
		router.ServeHTTP(writer, request)

		var response helper.BodyResponse
		err := json.Unmarshal(writer.Body.Bytes(), &response)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, writer.Code)
		assert.Equal(t, http.StatusOK, response.Code)

		data := response.Data.(map[string]interface{})
		assert.Equal(t, company.ID, data["id"])

	})

	t.Run("get company by unavailable id should failed", func(t *testing.T) {
		companyService.Mock.On("ById", "99").Return(nil, errors.New("data not found"))

		request, _ := http.NewRequest(http.MethodGet, "/companies/99", nil)
		writer := httptest.NewRecorder()
		router.ServeHTTP(writer, request)

		var response helper.BodyResponse
		err := json.Unmarshal(writer.Body.Bytes(), &response)

		assert.NotNil(t, err)
		assert.Equal(t, http.StatusBadRequest, writer.Code)

	})
}

func Test_companyHandler_Get(t *testing.T) {
	router.GET("/companies", handler.Get)

	t.Run("get company should success", func(t *testing.T) {
		var companies = entity.Companies
		companyService.Mock.On("Get", mock.Anything).Return(companies, nil)

		request, _ := http.NewRequest(http.MethodGet, "/companies", nil)
		writer := httptest.NewRecorder()
		router.ServeHTTP(writer, request)

		var response helper.BodyResponse

		err := json.Unmarshal(writer.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, writer.Code)
		assert.Equal(t, http.StatusOK, response.Code)

		data := response.Data.([]interface{})
		assert.Equal(t, len(companies), len(data))
	})
}

func Test_companyHandler_Add(t *testing.T) {
	router.POST("/companies", handler.Add)

	t.Run("add new company should success", func(t *testing.T) {
		var company = entity.Company{
			ID:      xid.New().String(),
			Name:    "Toshiba",
			CEO:     "Mr. Toshiba",
			Revenue: "600",
		}
		companyService.Mock.On("Add", company).Return(company, nil)

		jsonCompanyParam, _ := json.Marshal(company)

		request, _ := http.NewRequest(http.MethodPost, "/companies", bytes.NewBuffer(jsonCompanyParam))
		writer := httptest.NewRecorder()
		router.ServeHTTP(writer, request)

		var response helper.BodyResponse

		err := json.Unmarshal(writer.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, writer.Code)
		assert.Equal(t, http.StatusCreated, response.Code)

		data := response.Data.(map[string]interface{})
		assert.Equal(t, company.ID, data["id"])

	})
}

func Test_companyHandler_Update(t *testing.T) {
	router.PUT("/companies/:id", handler.Update)

	t.Run("Update new company should success", func(t *testing.T) {
		var company = entity.Company{
			ID:      "1",
			Name:    "Update",
			CEO:     "Mr. update",
			Revenue: "600",
		}
		companyService.Mock.On("Update", "1", company).Return(company, nil)

		jsonCompanyParam, _ := json.Marshal(company)

		request, _ := http.NewRequest(http.MethodPut, "/companies/1", bytes.NewBuffer(jsonCompanyParam))
		writer := httptest.NewRecorder()
		router.ServeHTTP(writer, request)

		var response helper.BodyResponse

		err := json.Unmarshal(writer.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, writer.Code)
		assert.Equal(t, http.StatusCreated, response.Code)

		data := response.Data.(map[string]interface{})
		assert.Equal(t, company.ID, data["id"])

	})
}

func Test_companyHandler_Delete(t *testing.T) {
	router.DELETE("/companies/:id", handler.Delete)

	t.Run("delete company should success", func(t *testing.T) {

		companyService.Mock.On("Delete", "1").Return(nil)

		request, _ := http.NewRequest(http.MethodDelete, "/companies/1", nil)
		writer := httptest.NewRecorder()
		router.ServeHTTP(writer, request)

		var response helper.BodyResponse

		err := json.Unmarshal(writer.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, writer.Code)
		assert.Equal(t, http.StatusOK, response.Code)

	})
}
