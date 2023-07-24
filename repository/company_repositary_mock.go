package repository

import (
	"github.com/stretchr/testify/mock"
	"gitlab.com/jeelabs/learnings/goex-rest-unit-test/entity"
)

type CompanyRepositoryMock struct {
	Mock mock.Mock
}

func (m *CompanyRepositoryMock) Get() *[]entity.Company {
	args := m.Mock.Called()
	return args.Get(0).(*[]entity.Company)
}

func (m *CompanyRepositoryMock) ById(id string) (*entity.Company, error) {
	args := m.Mock.Called(id)
	return args.Get(0).(*entity.Company), args.Error(1)
}

func (m *CompanyRepositoryMock) Add(company entity.Company) (*entity.Company, error) {
	args := m.Mock.Called(company)
	//comp := args.Get(0).(entity.Company)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	} else {
		return args.Get(0).(*entity.Company), nil
	}
}

func (m *CompanyRepositoryMock) Update(id string, company entity.Company) (*entity.Company, error) {
	args := m.Mock.Called(id, company)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	} else {
		return args.Get(0).(*entity.Company), nil
	}
}

func (m *CompanyRepositoryMock) Delete(id string) error {
	args := m.Mock.Called(id)
	if args.Get(0) == nil {
		return nil
	} else {
		return args.Error(0)
	}
}
