package service

import (
	"go-jwt/dto"
	"go-jwt/entity"
	"go-jwt/package/errs"
	"go-jwt/package/helpers"
	"go-jwt/repository/user_repository"
	"net/http"
)

type UserService interface {
	CreateNewUser(payload dto.NewUserRequest) (*dto.NewUserResponse, errs.MessageErr)
	Login(payload dto.NewUserRequest) (*dto.LoginResponse, errs.MessageErr)
}

type userService struct {
	userRepo user_repository.UserRepository
}

func NewUserService(userRepo user_repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

// xwk kosong apa kaga
func (u *userService) Login(payload dto.NewUserRequest) (*dto.LoginResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(payload)

	if err != nil {
		return nil, err
	}

	user, err := u.userRepo.GetUserByEmail(payload.Email)

	if err != nil {
		if err.Status() == http.StatusNotFound {
			return nil, errs.NewUnauthenticatedError("invalid email/password")
		}
		return nil, err
	}

	isValidPassword := user.ComparePassword(payload.Password)

	if !isValidPassword {
		return nil, errs.NewUnauthenticatedError("invalid email/password")
	}

	response := dto.LoginResponse{
		Result:     "success",
		Message:    "logged in successfullyy, welcome to IF Products",
		StatusCode: http.StatusOK,
		Data: dto.TokenResponse{
			Token: user.GenerateToken(),
		},
	}

	return &response, nil
}

func (u *userService) CreateNewUser(payload dto.NewUserRequest) (*dto.NewUserResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(payload)

	if err != nil {
		return nil, err
	}
	// ini buat narok email pass
	userEntity := entity.User{
		Email:    payload.Email,
		Password: payload.Password,
	}

	//  buat encrypt pass nya

	err = userEntity.HashPassword()

	if err != nil {
		return nil, err
	}

	err = u.userRepo.CreateNewUser(userEntity)

	if err != nil {
		return nil, err
	}

	response := dto.NewUserResponse{
		Result:     "success",
		Message:    "user registered successfully",
		StatusCode: http.StatusCreated,
	}

	return &response, nil
}
