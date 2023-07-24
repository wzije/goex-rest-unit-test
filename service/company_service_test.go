package service_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gitlab.com/jeelabs/learnings/goex-rest-unit-test/entity"
	"gitlab.com/jeelabs/learnings/goex-rest-unit-test/repository"
	. "gitlab.com/jeelabs/learnings/goex-rest-unit-test/service"
	"testing"
)

var repo = &repository.CompanyRepositoryMock{Mock: mock.Mock{}}
var service = NewCompanyService(repo)

func Test_companyService_Get(t *testing.T) {
	repo.Mock.On("Get", mock.Anything).Return(&[]entity.Company{})

	company, err := service.Get()
	assert.Nil(t, err)
	assert.NotNil(t, company)
}

func Test_companyService_ById(t *testing.T) {
	repo.Mock.On("ById", "1").Return(&entity.Company{}, nil)

	result, err := service.ById("1")
	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func Test_companyService_Add(t *testing.T) {

	t.Run("add company should success", func(t *testing.T) {
		comp := entity.Company{
			ID:      "10",
			Name:    "Axio",
			CEO:     "Mr. Axio",
			Revenue: "500",
		}
		repo.Mock.On("Add", comp).Return(&comp, nil)

		result, err := service.Add(comp)
		assert.NotNil(t, result)
		assert.Nil(t, err)
		assert.Equal(t, comp.ID, result.ID)
	})

	t.Run("Fail to add company should failed", func(t *testing.T) {
		comp := entity.Company{}
		err := errors.New("data not found")

		repo.Mock.On("Add", comp).Return(nil, err)

		result, err := service.Add(comp)
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})

}

func Test_companyService_Update(t *testing.T) {
	t.Run("update company should success", func(t *testing.T) {
		comp := entity.Company{
			ID:      "1",
			Name:    "Advance",
			CEO:     "Mr. E",
			Revenue: "900",
		}
		repo.Mock.On("Update", "1", comp).Return(&comp, nil)

		result, err := service.Update("1", comp)
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, comp.Name, result.Name)
	})

	t.Run("update empty company should failed", func(t *testing.T) {
		id := "99"
		comp := entity.Company{}
		repo.Mock.On("Update", id, comp).Return(nil, errors.New("data not found"))

		result, err := service.Update(id, comp)
		assert.Nil(t, result)
		assert.NotNil(t, err)
	})
}

func Test_companyService_Delete(t *testing.T) {
	t.Run("delete company should success", func(t *testing.T) {
		repo.Mock.On("Delete", "1").Return(nil)
		err := service.Delete("1")
		assert.Nil(t, err)
	})
	t.Run("delete unavailable company should failed", func(t *testing.T) {
		id := "99"
		repo.Mock.On("Delete", id).Return(errors.New("data not found"))
		err := service.Delete(id)
		assert.NotNil(t, err)
		assert.Equal(t, "data not found", err.Error())
	})
}
