package user_repository

import (
	"go-jwt/entity"
	"go-jwt/package/errs"
)

type UserRepository interface {
	CreateNewUser(user entity.User) errs.MessageErr
	// dapetin satu user
	GetUserById(userId int) (*entity.User, errs.MessageErr)
	GetUserByEmail(userEmail string) (*entity.User, errs.MessageErr)
}
