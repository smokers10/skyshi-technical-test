package todo

import (
	body_request "skyshi-technical-test/model/web/request_body/todo"

	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	Mock mock.Mock
}

func (mr *MockRepository) GetAll() []Todo {
	args := mr.Mock.Called()
	return args.Get(0).([]Todo)
}

func (mr *MockRepository) GetOne(id int) *Todo {
	args := mr.Mock.Called(id)
	return args.Get(0).(*Todo)
}

func (mr *MockRepository) Create(data *body_request.Create) (created_id int, err error) {
	args := mr.Mock.Called(data)
	return args.Int(0), args.Error(1)
}

func (mr *MockRepository) Delete(id int) error {
	args := mr.Mock.Called(id)
	return args.Error(0)
}

func (mr *MockRepository) Update(data *body_request.Update) error {
	args := mr.Mock.Called(data)
	return args.Error(0)
}
