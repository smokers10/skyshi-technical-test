package todo

import (
	"fmt"
	"skyshi-technical-test/model/domain/todo"
	response "skyshi-technical-test/model/web"
	body_request "skyshi-technical-test/model/web/request_body/todo"
	"time"
)

type todoService struct {
	repository todo.Repository
}

// Create implements todo.Service
func (ts *todoService) Create(body *body_request.Create) response.Response {
	// call repository to presist todo data
	inserted_id, err := ts.repository.Create(body)

	// if create has an error
	if err != nil {
		return response.Response{
			Status:  "Failed",
			Message: "Server Error",
		}
	}

	// set up required response
	todo := todo.Todo{
		ID:              inserted_id,
		ActivityGroupID: body.ActivityGroupID,
		Title:           body.Title,
		IsActive:        body.IsActive,
		Priority:        body.Priority,
		CreatedAt:       time.Now().String(),
		UpdatedAt:       time.Now().String(),
	}

	return response.Response{
		Status:  "Success",
		Message: "Success",
		Data:    todo,
	}
}

// Delete implements todo.Service
func (ts *todoService) Delete(body *body_request.Delete) response.Response {
	// fetch selected todo item
	selected := ts.repository.GetOne(body.ID)

	// if todo not exists
	if selected.ID == 0 {
		return response.Response{
			Status:  "Failed",
			Message: "todo is not found",
		}
	}

	// if todo already deleted
	if selected.DeletedAt != "" && selected.ID != 0 {
		return response.Response{
			Status:  "Failed",
			Message: "data already deleted",
		}
	}

	// then if selected todo is not deleted and exists
	if err := ts.repository.Delete(body.ID); err != nil {
		fmt.Print(err)
		return response.Response{
			Status:  "Failed",
			Message: "Server Error",
		}
	}

	return response.Response{
		Status:  "success",
		Message: "success",
	}
}

// GetAll implements todo.Service
func (ts *todoService) GetAll() response.Response {
	// call related repository method to query all todo data
	todos := ts.repository.GetAll()

	// if todo is empty
	if len(todos) == 0 {
		return response.Response{
			Status:  "Failed",
			Message: "Todo is empty",
		}
	}

	return response.Response{
		Status:  "Success",
		Message: "Success",
		Data:    todos,
	}
}

// GetOne implements todo.Service
func (ts *todoService) GetOne(body *body_request.GetOne) response.Response {
	// fetch selected todo item
	todo := ts.repository.GetOne(body.ID)

	// if todo is empty
	if todo.ID == 0 {
		return response.Response{
			Status:  "Failed",
			Message: "Todo not found",
		}
	}

	return response.Response{
		Status:  "Success",
		Message: "Success",
		Data:    todo,
	}
}

// Update implements todo.Service
func (ts *todoService) Update(body *body_request.Update) response.Response {
	// set current date for update_at field
	now := time.Now().String()

	body.UpdatedAt = now

	// call related repository method for update todo
	if err := ts.repository.Update(body); err != nil {
		fmt.Println(err)
		return response.Response{
			Status:  "Failed",
			Message: "Server Error",
		}
	}

	// fetch updated todo item
	todo := ts.repository.GetOne(body.ID)

	return response.Response{
		Status:  "Success",
		Message: "Success",
		Data:    todo,
	}
}

func TodoService(repo *todo.Repository) todo.Service {
	return &todoService{repository: *repo}
}
