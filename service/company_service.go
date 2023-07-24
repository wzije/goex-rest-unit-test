package service

import (
	"errors"
	"gitlab.com/jeelabs/learnings/goex-rest-unit-test/entity"
	"gitlab.com/jeelabs/learnings/goex-rest-unit-test/repository"
)

type CompanyService interface {
	Get() (*[]entity.Company, error)
	ById(id string) (*entity.Company, error)
	Add(company entity.Company) (*entity.Company, error)
	Update(id string, company entity.Company) (*entity.Company, error)
	Delete(id string) error
}

type companyService struct {
	repository repository.CompanyRepository
}

func (c companyService) Get() (*[]entity.Company, error) {
	companies := c.repository.Get()
	if companies == nil {
		return nil, errors.New("data not found")
	}
	return companies, nil
}

func (c companyService) ById(id string) (*entity.Company, error) {
	company, err := c.repository.ById(id)
	return company, err
}

func (c companyService) Add(comp entity.Company) (*entity.Company, error) {

	company, err := c.repository.Add(comp)

	if err != nil {
		return nil, err
	}

	return company, nil
}

func (c companyService) Update(id string, comp entity.Company) (*entity.Company, error) {
	company, err := c.repository.Update(id, comp)
	if err != nil {
		return nil, err
	}
	return company, nil
}

func (c companyService) Delete(id string) error {
	return c.repository.Delete(id)
}

func NewCompanyService(companyRepository repository.CompanyRepository) CompanyService {
	return companyService{repository: companyRepository}
}
