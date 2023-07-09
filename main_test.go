package main_test

import (
	"bytes"
	"github.com/goccy/go-json"
	"github.com/rs/xid"
	"github.com/stretchr/testify/assert"
	. "gitlab.com/jeelabs/learnings/go-crud"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {

	t.Run("test welcome page test", func(t *testing.T) {
		router := SetUpRouter()

		mockResponse := `{"message":"welcome to simple go crud"}`
		router.GET("/", WelcomeHandler)
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		writer := httptest.NewRecorder()
		router.ServeHTTP(writer, request)

		responseData, _ := ioutil.ReadAll(writer.Body)

		assert.Equal(t, mockResponse, string(responseData))
		assert.Equal(t, http.StatusOK, writer.Code)
	})

	t.Run("get Companies", func(t *testing.T) {
		router := SetUpRouter()

		router.GET("/Companies", GetCompanies)
		request, _ := http.NewRequest(http.MethodGet, "/Companies", nil)
		writer := httptest.NewRecorder()
		router.ServeHTTP(writer, request)

		var companies []Company
		err := json.Unmarshal(writer.Body.Bytes(), &companies)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, writer.Code)
		assert.NotEmpty(t, companies)
	})

	t.Run("create new company", func(t *testing.T) {
		router := SetUpRouter()

		router.POST("/Companies", CreateCompany)

		id := xid.New().String()

		company := Company{
			ID:      id,
			Name:    "New Company",
			CEO:     "New CEO",
			Revenue: "2000",
		}
		jsonValue, _ := json.Marshal(company)

		request, _ := http.NewRequest(http.MethodPost, "/Companies", bytes.NewBuffer(jsonValue))

		writer := httptest.NewRecorder()
		router.ServeHTTP(writer, request)

		assert.Equal(t, http.StatusCreated, writer.Code)

		index := -1
		for i := 0; i < len(Companies); i++ {
			if Companies[i].ID == id {
				index = i
			}
		}

		assert.NotEqual(t, -1, index)
		assert.Equal(t, id, Companies[index].ID)

	})

	t.Run("update company", func(t *testing.T) {
		router := SetUpRouter()

		router.PUT("/Companies/:id", UpdateCompany)

		company := Company{
			Name:    "Update Company",
			CEO:     "update CEO",
			Revenue: "3000",
		}

		jsonValue, _ := json.Marshal(&company)

		t.Run("success update", func(t *testing.T) {
			request, _ := http.NewRequest(http.MethodPut, "/Companies/1", bytes.NewBuffer(jsonValue))
			writer := httptest.NewRecorder()
			router.ServeHTTP(writer, request)
			assert.Equal(t, http.StatusCreated, writer.Code)
		})

		t.Run("failed update", func(t *testing.T) {
			request, _ := http.NewRequest(http.MethodPut, "/Companies/12", bytes.NewBuffer(jsonValue))
			writer := httptest.NewRecorder()
			router.ServeHTTP(writer, request)
			assert.Equal(t, http.StatusNotFound, writer.Code)
		})

	})

	t.Run("delete company", func(t *testing.T) {
		router := SetUpRouter()
		router.DELETE("/Companies/:id", DeleteCompany)

		t.Run("success delete", func(t *testing.T) {
			request, _ := http.NewRequest(http.MethodDelete, "/Companies/2", nil)
			writer := httptest.NewRecorder()
			router.ServeHTTP(writer, request)
			assert.Equal(t, http.StatusOK, writer.Code)
		})

		t.Run("notfound delete", func(t *testing.T) {
			request, _ := http.NewRequest(http.MethodDelete, "/Companies/30",
				nil)
			writer := httptest.NewRecorder()
			router.ServeHTTP(writer, request)
			assert.Equal(t, http.StatusNotFound, writer.Code)
		})

	})
}
