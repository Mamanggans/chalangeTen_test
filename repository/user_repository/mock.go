package user_repository

import (
	"go-jwt/entity"
	"go-jwt/package/errs"
)

var (
	CreateNewUser  func(user entity.User) errs.MessageErr
	GetUserById    func(userId int) (*entity.User, errs.MessageErr)
	GetUserByEmail func(userEmail string) (*entity.User, errs.MessageErr)
)

type userRepoMock struct{}

func NewUserRepoMock() UserRepository {
	return &userRepoMock{}
}

func (u *userRepoMock) CreateNewUser(user entity.User) errs.MessageErr {
	return CreateNewUser(user)
}
func (u *userRepoMock) GetUserById(userId int) (*entity.User, errs.MessageErr) {
	return GetUserById(userId)
}
func (u *userRepoMock) GetUserByEmail(userEmail string) (*entity.User, errs.MessageErr) {
	return GetUserByEmail(userEmail)
}
