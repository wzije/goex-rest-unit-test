package service

import (
	"github.com/stretchr/testify/mock"
	"gitlab.com/jeelabs/learnings/goex-rest-unit-test/entity"
)

type CompanyServiceMock struct {
	Mock mock.Mock
}

func (c *CompanyServiceMock) Get() (*[]entity.Company, error) {
	args := c.Mock.Called()

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	result := args.Get(0).([]entity.Company)

	return &result, nil
}

func (c *CompanyServiceMock) ById(id string) (*entity.Company, error) {

	args := c.Mock.Called(id)

	if id == "" {
		return nil, args.Error(1)
	}

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	data := args.Get(0).(entity.Company)

	return &data, nil
}

func (c *CompanyServiceMock) Add(company entity.Company) (*entity.Company, error) {
	args := c.Mock.Called(company)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	data := args.Get(0).(entity.Company)

	return &data, nil
}

func (c *CompanyServiceMock) Update(id string, company entity.Company) (*entity.Company, error) {
	args := c.Mock.Called(id, company)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	data := args.Get(0).(entity.Company)

	return &data, nil
}

func (c *CompanyServiceMock) Delete(id string) error {
	args := c.Mock.Called(id)
	if args.Get(0) == nil {
		return args.Error(0)
	}
	return nil
}
