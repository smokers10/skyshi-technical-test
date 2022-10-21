package activity

import (
	response "skyshi-technical-test/model/web"
	request_body "skyshi-technical-test/model/web/request_body/activity"
)

type Activity struct {
	ID        int    `json:"id,omitempty"`
	Email     string `json:"email,omitempty"`
	Title     string `json:"title,omitempty"`
	DeletedAt string `json:"deleted_at,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

type Service interface {
	GetAll() response.Response
	GetOne(body *request_body.GetOne) response.Response
	Create(body *request_body.Create) response.Response
	Delete(body *request_body.Delete) response.Response
}

type Repository interface {
	GetAll() []Activity
	GetOne(id int) *Activity
	Create(data *request_body.Create) (*Activity, error)
	Delete(is int) error
}
