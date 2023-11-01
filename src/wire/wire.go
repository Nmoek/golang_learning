//go:build wireinject

// Package wire
// @Description: wire工具示例
package wire

import (
	"github.com/google/wire"
	"wire_test/repository"
	"wire_test/repository/dao"
)

func InitUserRepository() *repository.UserRepository {
	wire.Build(repository.NewUserRepository, dao.NewUserDao, InitDB)

	return &repository.UserRepository{}
}
