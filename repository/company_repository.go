package repository

import (
	"errors"
	"gitlab.com/jeelabs/learnings/goex-rest-unit-test/entity"
)

type CompanyRepository interface {
	Get() *[]entity.Company
	ById(id string) (*entity.Company, error)
	Add(company entity.Company) (*entity.Company, error)
	Update(id string, company entity.Company) (*entity.Company, error)
	Delete(id string) error
}

type companyRepository struct {
}

func (c companyRepository) Get() *[]entity.Company {
	return &entity.Companies
}

func (c companyRepository) ById(id string) (*entity.Company, error) {
	index := -1

	for i, v := range entity.Companies {
		if v.ID == id {
			index = i
		}
	}

	if index == -1 {
		return nil, errors.New("data not found")
	}

	return &entity.Companies[index], nil
}

func (c companyRepository) Add(company entity.Company) (*entity.Company, error) {
	entity.Companies = append(entity.Companies, company)
	return &company, nil
}

func (c companyRepository) Update(id string, company entity.Company) (*entity.Company, error) {
	index := -1

	for i, v := range entity.Companies {
		if v.ID == id {
			index = i
		}
	}

	if index == -1 {
		return nil, errors.New("data not found")
	}

	comp := entity.Companies[index]
	comp.Name = company.Name
	comp.CEO = company.CEO
	comp.Revenue = company.Revenue

	return &comp, nil
}

func (c companyRepository) Delete(id string) error {
	index := -1

	for i, v := range entity.Companies {
		if v.ID == id {
			index = i
		}
	}

	if index == -1 {
		return errors.New("data not found")
	}

	entity.Companies = append(entity.Companies[:index], entity.Companies[index+1:]...)

	return nil
}

func NewCompanyRepository() CompanyRepository {
	return companyRepository{}
}
