package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUserFound(t *testing.T) {
	// Initialization

	// Execution

	user, err := UserDao.GetUser(0)

	// Validation
	assert.Nil(t, user, "We were not expecting a user with id 0")
	assert.NotNil(t, err)
	assert.EqualValues(t, err.StatusCode, http.StatusNotFound)
	assert.EqualValues(t, err.Code, "not_found")
	assert.EqualValues(t, err.Message, "user 0 does not exists")
}

func TestGetUserNoError(t *testing.T) {

	user, err := UserDao.GetUser(123)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 1, user.ID)
	assert.EqualValues(t, "Dave", user.FirstName)
	assert.EqualValues(t, "Augustus", user.LastName)
	assert.EqualValues(t, "email@email.com", user.Email)
	//
}
