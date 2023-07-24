package repository_test

import (
	"github.com/stretchr/testify/assert"
	"gitlab.com/jeelabs/learnings/goex-rest-unit-test/entity"
	. "gitlab.com/jeelabs/learnings/goex-rest-unit-test/repository"
	"testing"
)

var repo = NewCompanyRepository()

func Test_companyRepository_Get(t *testing.T) {
	companies := repo.Get()
	assert.NotNil(t, companies)
	assert.Equal(t, len(entity.Companies), len(*companies))
}

func Test_companyRepository_ById(t *testing.T) {
	t.Run("should success", func(t *testing.T) {
		company, err := repo.ById("1")
		assert.Nil(t, err)
		assert.NotNil(t, company)
		assert.Equal(t, "1", company.ID)
		assert.Equal(t, "DELL", company.Name)
	})

	t.Run("should data not found", func(t *testing.T) {
		company, err := repo.ById("10")
		assert.Nil(t, company)
		assert.NotNil(t, err)
		assert.Equal(t, "data not found", err.Error())
	})
}

func Test_companyRepository_Add(t *testing.T) {
	newCompany := entity.Company{
		ID:      "5",
		Name:    "Apple",
		CEO:     "Steve Job",
		Revenue: "5000",
	}
	comp, err := repo.Add(newCompany)
	assert.Nil(t, err)
	assert.NotNil(t, comp)
	assert.Equal(t, newCompany.ID, comp.ID)
	assert.Equal(t, 5, len(entity.Companies))
}

func Test_companyRepository_Update(t *testing.T) {

	t.Run("should success update", func(t *testing.T) {
		newCompany := entity.Company{
			Name:    "Apple",
			CEO:     "Steve Job",
			Revenue: "5000",
		}
		comp, err := repo.Update("4", newCompany)
		assert.Nil(t, err)
		assert.NotNil(t, comp)

		selectComp, err := repo.ById("4")
		assert.Nil(t, err)
		assert.Equal(t, selectComp.ID, comp.ID)
	})

	t.Run("should failed or data not found", func(t *testing.T) {
		newCompany := entity.Company{
			Name:    "Apple",
			CEO:     "Steve Job",
			Revenue: "5000",
		}
		comp, err := repo.Update("10", newCompany)
		assert.NotNil(t, err)
		assert.Nil(t, comp)
		assert.Equal(t, "data not found", err.Error())
	})

}

func Test_companyRepository_Delete(t *testing.T) {
	oldTotal := len(entity.Companies)
	err := repo.Delete("1")
	assert.Nil(t, err)
	assert.Equal(t, oldTotal-1, len(entity.Companies))

}
