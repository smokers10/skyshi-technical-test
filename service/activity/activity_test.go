package activity

import (
	"errors"
	"skyshi-technical-test/infrastructure/lib"
	"skyshi-technical-test/model/domain/activity"
	request_body "skyshi-technical-test/model/web/request_body/activity"
	"testing"

	"github.com/stretchr/testify/mock"
)

var (
	repository = activity.MockRepository{Mock: mock.Mock{}}
	service    = activityService{repository: &repository}
)

func TestCreate(t *testing.T) {
	t.Run("failed to presist", func(t *testing.T) {
		expected := lib.Expected{
			Status:  "Failed",
			Message: "Server Error",
		}

		repository.Mock.On("Create", mock.Anything).Return(&activity.Activity{}, errors.New(mock.Anything)).Once()

		actual := service.Create(&request_body.Create{})

		lib.UnitTesting().CommonAssertion(t, &expected, &actual, lib.DefaultOption)
	})

	t.Run("failed to presist", func(t *testing.T) {
		expected := lib.Expected{
			Status:  "Success",
			Message: "Success",
		}

		repository.Mock.On("Create", mock.Anything).Return(&activity.Activity{
			ID:        1,
			Email:     mock.Anything,
			Title:     mock.Anything,
			DeletedAt: mock.Anything,
			CreatedAt: mock.Anything,
			UpdatedAt: mock.Anything,
		}, nil).Once()

		actual := service.Create(&request_body.Create{})

		lib.UnitTesting().CommonAssertion(t, &expected, &actual, lib.UnitTesting().DataChecking(true))
	})
}

func TestDelete(t *testing.T) {
	t.Run("activity not found", func(t *testing.T) {
		expected := lib.Expected{
			Status:  "Failed",
			Message: "Activity not found",
		}

		repository.Mock.On("GetOne", mock.Anything).Return(&activity.Activity{
			ID:        0,
			Email:     mock.Anything,
			Title:     mock.Anything,
			DeletedAt: mock.Anything,
			CreatedAt: mock.Anything,
			UpdatedAt: mock.Anything,
		}).Once()

		actual := service.Delete(&request_body.Delete{})

		lib.UnitTesting().CommonAssertion(t, &expected, &actual, lib.DefaultOption)
	})

	t.Run("activity already deleted", func(t *testing.T) {
		expected := lib.Expected{
			Status:  "Failed",
			Message: "Activity already deleted",
		}

		repository.Mock.On("GetOne", mock.Anything).Return(&activity.Activity{
			ID:        1,
			Email:     mock.Anything,
			Title:     mock.Anything,
			DeletedAt: mock.Anything,
			CreatedAt: mock.Anything,
			UpdatedAt: mock.Anything,
		}).Once()

		actual := service.Delete(&request_body.Delete{})

		lib.UnitTesting().CommonAssertion(t, &expected, &actual, lib.DefaultOption)
	})

	t.Run("failed to delete", func(t *testing.T) {
		expected := lib.Expected{
			Message: "Server Error",
			Status:  "Failed",
		}

		repository.Mock.On("GetOne", mock.Anything).Return(&activity.Activity{
			ID:        1,
			Email:     mock.Anything,
			Title:     mock.Anything,
			DeletedAt: "",
			CreatedAt: mock.Anything,
			UpdatedAt: mock.Anything,
		}).Once()

		repository.Mock.On("Delete", mock.Anything).Return(errors.New(mock.Anything)).Once()

		actual := service.Delete(&request_body.Delete{})

		lib.UnitTesting().CommonAssertion(t, &expected, &actual, lib.DefaultOption)
	})

	t.Run("success", func(t *testing.T) {
		expected := lib.Expected{
			Status:  "Success",
			Message: "Success",
		}

		repository.Mock.On("GetOne", mock.Anything).Return(&activity.Activity{
			ID:        1,
			Email:     mock.Anything,
			Title:     mock.Anything,
			DeletedAt: "",
			CreatedAt: mock.Anything,
			UpdatedAt: mock.Anything,
		}).Once()

		repository.Mock.On("Delete", mock.Anything).Return(nil).Once()

		actual := service.Delete(&request_body.Delete{})

		lib.UnitTesting().CommonAssertion(t, &expected, &actual, lib.DefaultOption)
	})
}

func TestGetAll(t *testing.T) {
	t.Run("activity is empty", func(t *testing.T) {
		expected := lib.Expected{
			Status:  "Success",
			Message: "There is no activity yet",
		}

		repository.Mock.On("GetAll").Return([]activity.Activity{}).Once()

		actual := service.GetAll()

		lib.UnitTesting().CommonAssertion(t, &expected, &actual, lib.DefaultOption)
	})

	t.Run("activity found", func(t *testing.T) {
		expected := lib.Expected{
			Status:  "Success",
			Message: "Success",
		}

		repository.Mock.On("GetAll").Return([]activity.Activity{
			{
				ID:        1,
				Email:     mock.Anything,
				Title:     mock.Anything,
				DeletedAt: mock.Anything,
				CreatedAt: mock.Anything,
				UpdatedAt: mock.Anything,
			},
		}).Once()

		actual := service.GetAll()

		lib.UnitTesting().CommonAssertion(t, &expected, &actual, lib.DefaultOption)
	})
}

func TestGetOne(t *testing.T) {
	t.Run("not found", func(t *testing.T) {
		expected := lib.Expected{
			Status:  "Success",
			Message: "Activity not found",
		}

		repository.Mock.On("GetOne", mock.Anything).Return(&activity.Activity{
			ID:        0,
			Email:     "",
			Title:     "",
			DeletedAt: "",
			CreatedAt: "",
			UpdatedAt: "",
		}).Once()

		actual := service.GetOne(&request_body.GetOne{ID: 1})

		lib.UnitTesting().CommonAssertion(t, &expected, &actual, lib.DefaultOption)
	})

	t.Run("found", func(t *testing.T) {
		expected := lib.Expected{
			Status:  "Success",
			Message: "Success",
		}

		repository.Mock.On("GetOne", mock.Anything).Return(&activity.Activity{
			ID:        1,
			Email:     mock.Anything,
			Title:     mock.Anything,
			DeletedAt: mock.Anything,
			CreatedAt: mock.Anything,
			UpdatedAt: mock.Anything,
		}).Once()

		actual := service.GetOne(&request_body.GetOne{ID: 1})

		lib.UnitTesting().CommonAssertion(t, &expected, &actual, lib.UnitTesting().DataChecking(true))
	})
}
