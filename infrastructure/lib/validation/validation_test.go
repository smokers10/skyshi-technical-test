package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmail(t *testing.T) {
	t.Run("empty email", func(t *testing.T) {
		r := Validation().Email.EmailRequired("")

		assert.Equal(t, true, r)
	})

	t.Run("invalid email", func(t *testing.T) {
		r := Validation().Email.EmailValidity("angkasa@spaceKill.com")

		assert.Equal(t, true, r)
	})

	t.Run("exists email", func(t *testing.T) {
		r := Validation().Email.EmailRequired("johndoe@gmail.com")

		assert.Equal(t, false, r)
	})

	t.Run("valid email", func(t *testing.T) {
		r := Validation().Email.EmailValidity("nadzarmutaqin@gmail.com")

		assert.Equal(t, false, r)
	})
}

func TestString(t *testing.T) {
	t.Run("empty string", func(t *testing.T) {
		r := Validation().String.Required("")

		assert.Equal(t, true, r)
	})

	t.Run("length not meet", func(t *testing.T) {
		r := Validation().String.Length("nadzar", 7, 10)

		assert.Equal(t, true, r)
	})

	t.Run("string not match", func(t *testing.T) {
		r := Validation().String.Match("ajun", "arjuna")

		assert.Equal(t, true, r)
	})

	t.Run("exists string", func(t *testing.T) {
		r := Validation().String.Required("john doe")

		assert.Equal(t, false, r)
	})

	t.Run("length meet", func(t *testing.T) {
		r := Validation().String.Length("alexander I", 7, 10)

		assert.Equal(t, true, r)
	})

	t.Run("string match", func(t *testing.T) {
		r := Validation().String.Match("arjuna", "arjuna")

		assert.Equal(t, false, r)
	})
}

func TestArray(t *testing.T) {
	a1 := []string{}
	a2 := []string{"a", "b", "c"}

	t.Run("empty array", func(t *testing.T) {
		r := Validation().Array.NotEmpty(len(a1))

		assert.Equal(t, true, r)
	})

	t.Run("filled array", func(t *testing.T) {
		r := Validation().Array.NotEmpty(len(a2))

		assert.Equal(t, false, r)
	})
}

func TestResponse(t *testing.T) {
	t.Run("empty response", func(t *testing.T) {
		r := Validation().Response.SendResponse("")

		assert.Equal(t, "", r.Message)
	})

	t.Run("empty response", func(t *testing.T) {
		m := "hello world"
		r := Validation().Response.SendResponse(m)

		assert.Equal(t, m, r.Message)
	})
}
