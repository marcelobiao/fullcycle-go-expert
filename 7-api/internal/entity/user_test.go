package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Marcelo", "marcelo@gmail.com", "123456")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.Equal(t, "Marcelo", user.Name)
	assert.Equal(t, "marcelo@gmail.com", user.Email)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("Marcelo", "marcelo@gmail.com", "123456")
	assert.Nil(t, err)
	assert.Nil(t, user.ValidatePassword("123456"))
	assert.NotNil(t, user.ValidatePassword("1234567"))
	assert.NotEqual(t, "123456", user.Password)
}
