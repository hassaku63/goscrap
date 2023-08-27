//go:build wireinject
// +build wireinject

package myservice

import (
	"github.com/google/wire"
)

func InitializeUserService() *UserService {
	wire.Build(
		NewUserService,
		NewUserRepository,
		// UserService は struct ではなく、interface である UserRepository に依存している
		// このケースでは、wire.Bind() を利用する必要がある
		wire.Bind(new(UserRepository), new(*User)),
	)
	return &UserService{}
}
