package todo

import (
	response "skyshi-technical-test/model/web"
	body_request "skyshi-technical-test/model/web/request_body/todo"
)

type Todo struct {
	ID              int
	ActivityGroupID int
	Title           string
	IsActive        string
	Priority        string
	CreatedAt       string
	UpdatedAt       string
	DeletedAt       string
}

type Service interface {
	GetAll() response.Response
	GetOne(body *body_request.GetOne) response.Response
	Create(body *body_request.Create) response.Response
	Delete(body *body_request.Delete) response.Response
	Update(body *body_request.Update) response.Response
}

type Repository interface {
	GetAll() []Todo
	GetOne(id int) *Todo
	Create(data *body_request.Create) (created_id int, err error)
	Delete(id int) error
	Update(data *body_request.Update) error
}
