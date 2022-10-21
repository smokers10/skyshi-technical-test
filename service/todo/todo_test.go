package todo

import (
	"errors"
	"skyshi-technical-test/infrastructure/lib"
	"skyshi-technical-test/model/domain/todo"

	"testing"

	body_request "skyshi-technical-test/model/web/request_body/todo"

	"github.com/stretchr/testify/mock"
)

var (
	repository = todo.MockRepository{Mock: mock.Mock{}}
	service    = todoService{
		repository: &repository,
	}
)

func TestCreate(t *testing.T) {
	t.Run("error when create", func(t *testing.T) {
		expected := lib.Expected{
			Status:  "Failed",
			Message: "Server Error",
		}

		repository.Mock.On("Create", mock.Anything).Return(0, errors.New(mock.Anything)).Once()

		actual := service.Create(&body_request.Create{
			ActivityGroupID: 0,
			Title:           mock.Anything,
			IsActive:        mock.Anything,
			Priority:        mock.Anything,
		})

		lib.UnitTesting().CommonAssertion(t, &expected, &actual, lib.DefaultOption)
	})

	t.Run("success create operation", func(t *testing.T) {
		expected := lib.Expected{
			Status:  "Success",
			Message: "Success",
		}

		repository.Mock.On("Create", mock.Anything).Return(1, nil).Once()

		actual := service.Create(&body_request.Create{
			ActivityGroupID: 0,
			Title:           mock.Anything,
			IsActive:        mock.Anything,
			Priority:        mock.Anything,
		})

		lib.UnitTesting().CommonAssertion(t, &expected, &actual, lib.UnitTesting().DataChecking(true))
	})
}

func TestDelete(t *testing.T) {
	t.Run("todo not exists", func(t *testing.T) {
		expected := lib.Expected{
			Status:  "Failed",
			Message: "todo is not found",
		}

		repository.Mock.On("GetOne", mock.Anything).Return(&todo.Todo{
			ID:              0,
			ActivityGroupID: 0,
			Title:           mock.Anything,
			IsActive:        mock.Anything,
			Priority:        mock.Anything,
			CreatedAt:       mock.Anything,
			UpdatedAt:       mock.Anything,
			DeletedAt:       mock.Anything,
		}).Once()

		actual := service.Delete(&body_request.Delete{ID: 1})

		lib.UnitTesting().CommonAssertion(t, &expected, &actual, lib.DefaultOption)
	})

	t.Run("todo already deleted", func(t *testing.T) {
		expected := lib.Expected{
			Status:  "Failed",
			Message: "data already deleted",
		}

		repository.Mock.On("GetOne", mock.Anything).Return(&todo.Todo{
			ID:              1,
			ActivityGroupID: 0,
			Title:           mock.Anything,
			IsActive:        mock.Anything,
			Priority:        mock.Anything,
			CreatedAt:       mock.Anything,
			UpdatedAt:       mock.Anything,
			DeletedAt:       mock.Anything,
		}).Once()

		actual := service.Delete(&body_request.Delete{ID: 1})

		lib.UnitTesting().CommonAssertion(t, &expected, &actual, lib.DefaultOption)
	})

	t.Run("error when deleting todo", func(t *testing.T) {
		expected := lib.Expected{
			Status:  "Failed",
			Message: "Server Error",
		}

		repository.Mock.On("GetOne", mock.Anything).Return(&todo.Todo{
			ID:              1,
			ActivityGroupID: 0,
			Title:           mock.Anything,
			IsActive:        mock.Anything,
			Priority:        mock.Anything,
			CreatedAt:       mock.Anything,
			UpdatedAt:       mock.Anything,
			DeletedAt:       "",
		}).Once()

		repository.Mock.On("Delete", 1).Return(errors.New(mock.Anything)).Once()

		actual := service.Delete(&body_request.Delete{ID: 1})

		lib.UnitTesting().CommonAssertion(t, &expected, &actual, lib.DefaultOption)
	})

	t.Run("success delete operation", func(t *testing.T) {
		expected := lib.Expected{
			Status:  "success",
			Message: "success",
		}

		repository.Mock.On("GetOne", mock.Anything).Return(&todo.Todo{
			ID:              1,
			ActivityGroupID: 0,
			Title:           mock.Anything,
			IsActive:        mock.Anything,
			Priority:        mock.Anything,
			CreatedAt:       mock.Anything,
			UpdatedAt:       mock.Anything,
			DeletedAt:       "",
		}).Once()

		repository.Mock.On("Delete", 1).Return(nil).Once()

		actual := service.Delete(&body_request.Delete{ID: 1})

		lib.UnitTesting().CommonAssertion(t, &expected, &actual, lib.DefaultOption)
	})
}

func TestGetAll(t *testing.T) {
	t.Run("empty todo", func(t *testing.T) {
		expected := lib.Expected{
			Status:  "Failed",
			Message: "Todo is empty",
		}

		repository.Mock.On("GetAll").Return([]todo.Todo{}).Once()

		actual := service.GetAll()

		lib.UnitTesting().CommonAssertion(t, &expected, &actual, lib.DefaultOption)
	})

	t.Run("todo data exists", func(t *testing.T) {
		expected := lib.Expected{
			Status:  "Success",
			Message: "Success",
		}

		repository.Mock.On("GetAll").Return([]todo.Todo{
			{
				ID:              1,
				ActivityGroupID: 2,
				Title:           mock.Anything,
				IsActive:        mock.Anything,
				Priority:        mock.Anything,
				CreatedAt:       mock.Anything,
				UpdatedAt:       mock.Anything,
				DeletedAt:       mock.Anything,
			},
		}).Once()

		actual := service.GetAll()

		lib.UnitTesting().CommonAssertion(t, &expected, &actual, lib.UnitTesting().DataChecking(true))
	})
}

func TestGetOne(t *testing.T) {
	t.Run("todo not found", func(t *testing.T) {
		expected := lib.Expected{
			Status:  "Failed",
			Message: "Todo not found",
		}

		repository.Mock.On("GetOne", 1).Return(&todo.Todo{ID: 0}).Once()

		actual := service.GetOne(&body_request.GetOne{ID: 1})

		lib.UnitTesting().CommonAssertion(t, &expected, &actual, lib.DefaultOption)
	})

	t.Run("todo found", func(t *testing.T) {
		expected := lib.Expected{
			Status:  "Success",
			Message: "Success",
		}

		repository.Mock.On("GetOne", 1).Return(&todo.Todo{ID: 1}).Once()

		actual := service.GetOne(&body_request.GetOne{ID: 1})

		lib.UnitTesting().CommonAssertion(t, &expected, &actual, lib.UnitTesting().DataChecking(true))
	})
}

func TestUpdate(t *testing.T) {
	t.Run("failed to update", func(t *testing.T) {
		expected := lib.Expected{
			Status:  "Failed",
			Message: "Server Error",
		}

		repository.Mock.On("Update", mock.Anything).Return(errors.New(mock.Anything)).Once()

		actual := service.Update(&body_request.Update{
			ID:              1,
			ActivityGroupID: 1,
			Title:           mock.Anything,
			IsActive:        mock.Anything,
			Priority:        mock.Anything,
			UpdatedAt:       mock.Anything,
		})

		lib.UnitTesting().CommonAssertion(t, &expected, &actual, lib.DefaultOption)
	})

	t.Run("success update operation", func(t *testing.T) {
		expected := lib.Expected{
			Status:  "Success",
			Message: "Success",
		}

		repository.Mock.On("Update", mock.Anything).Return(nil).Once()

		repository.Mock.On("GetOne", 1).Return(&todo.Todo{ID: 1}).Once()

		actual := service.Update(&body_request.Update{
			ID:              1,
			ActivityGroupID: 1,
			Title:           mock.Anything,
			IsActive:        mock.Anything,
			Priority:        mock.Anything,
			UpdatedAt:       mock.Anything,
		})

		lib.UnitTesting().CommonAssertion(t, &expected, &actual, lib.DefaultOption)
	})
}
