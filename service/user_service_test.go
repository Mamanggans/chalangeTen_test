package service

import (
	"go-jwt/dto"
	"go-jwt/entity"
	"go-jwt/package/errs"
	"go-jwt/repository/user_repository"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserService_CreateNewUser_Success(t *testing.T) {
	userRepo := user_repository.NewUserRepoMock()

	userService := NewUserService(userRepo)

	payload := dto.NewUserRequest{
		Email:    "test@mail.com",
		Password: "123456",
	}

	user_repository.CreateNewUser = func(user entity.User) errs.MessageErr {
		return nil
	}

	response, err := userService.CreateNewUser(payload)

	assert.Nil(t, err)

	assert.Equal(t, http.StatusCreated, response.StatusCode)

	assert.Equal(t, "success", response.Result)

	assert.Equal(t, "user registered successfully", response.Message)

}

func TestUserService_CreateNewUser_InvalidRequestBodyError(t *testing.T) {
	userService := NewUserService(nil)

	tt := []struct {
		name        string
		payload     dto.NewUserRequest
		expectation errs.MessageErr
	}{
		{
			name: "invalid email",
			payload: dto.NewUserRequest{
				Email:    "",
				Password: "123456",
			},
			expectation: errs.NewBadRequest("email cannot be empty"),
		},
		{
			name: "invalid password",
			payload: dto.NewUserRequest{
				Email:    "test@mail.com",
				Password: "",
			},
			expectation: errs.NewBadRequest("password cannot be empty"),
		},
	}

	for _, eachTest := range tt {
		t.Run(eachTest.name, func(t *testing.T) {

			response, err := userService.CreateNewUser(eachTest.payload)

			assert.NotNil(t, err)

			assert.Nil(t, response)

			assert.Equal(t, eachTest.expectation.Status(), err.Status())

			assert.Equal(t, eachTest.expectation.Message(), err.Message())

			assert.Equal(t, eachTest.expectation.Error(), err.Error())
		})
	}

}

func TestUserService_CreateNewUser_HashPasswordError(t *testing.T) {
	userService := NewUserService(nil)

	longChar := strings.Repeat("a", 73)

	payload := dto.NewUserRequest{
		Email:    "test@mail.com",
		Password: longChar,
	}

	response, err := userService.CreateNewUser(payload)

	assert.Nil(t, response)

	assert.NotNil(t, err)

	assert.Equal(t, http.StatusInternalServerError, err.Status())

	assert.Equal(t, "something went wrong", err.Message())

	assert.Equal(t, "INTERNAL_SERVER_ERROR", err.Error())
}

func TestUserService_CreateNewUser_UserRepoError(t *testing.T) {
	userRepo := user_repository.NewUserRepoMock()

	userService := NewUserService(userRepo)

	payload := dto.NewUserRequest{
		Email:    "test@mail.com",
		Password: "123456",
	}

	user_repository.CreateNewUser = func(user entity.User) errs.MessageErr {
		return errs.NewInternalServerError("something went wrong")
	}

	response, err := userService.CreateNewUser(payload)

	assert.Nil(t, response)

	assert.NotNil(t, err)

	assert.Equal(t, http.StatusInternalServerError, err.Status())

	assert.Equal(t, "something went wrong", err.Message())

	assert.Equal(t, "INTERNAL_SERVER_ERROR", err.Error())
}
