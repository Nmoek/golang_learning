// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"wire_test/repository"
	"wire_test/repository/dao"
)

// Injectors from wire.go:

func InitUserRepository() *repository.UserRepository {
	db := InitDB()
	userDao := dao.NewUserDao(db)
	userRepository := repository.NewUserRepository(userDao)
	return userRepository
}
