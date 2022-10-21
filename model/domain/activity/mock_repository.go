package activity

import (
	request_body "skyshi-technical-test/model/web/request_body/activity"

	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	Mock mock.Mock
}

func (mr *MockRepository) GetAll() []Activity {
	args := mr.Mock.Called()
	return args.Get(0).([]Activity)
}

func (mr *MockRepository) GetOne(id int) *Activity {
	args := mr.Mock.Called(id)
	return args.Get(0).(*Activity)
}

func (mr *MockRepository) Create(data *request_body.Create) (*Activity, error) {
	args := mr.Mock.Called(data)
	return args.Get(0).(*Activity), args.Error(1)
}

func (mr *MockRepository) Delete(id int) error {
	args := mr.Mock.Called(id)
	return args.Error(0)
}
