package lib

import (
	"testing"

	response "skyshi-technical-test/model/web"

	"github.com/stretchr/testify/assert"
)

type Expected struct {
	Message string
	Status  string
	Data    interface{}
}

type options struct {
	DataChecking bool
}

var DefaultOption = &options{DataChecking: false}

type unitTesting struct{}

func UnitTesting() *unitTesting {
	return &unitTesting{}
}

func (ut *unitTesting) CommonAssertion(t *testing.T, expected *Expected, actual *response.Response, option *options) {
	// response should not empty
	assert.NotEmpty(t, actual.Message)
	assert.NotEmpty(t, actual.Status)

	// response  expected detail must match w/ actual response
	assert.Equal(t, expected.Message, actual.Message)
	assert.Equal(t, expected.Status, actual.Status)

	// if data must checked
	if option.DataChecking {
		assert.NotEmpty(t, actual.Data)
	}
}

func (ut *unitTesting) DataChecking(data_checking bool) *options {
	return &options{DataChecking: data_checking}
}
