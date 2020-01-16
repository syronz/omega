//+build wireinject

package server

import (
	"github.com/google/wire"
	"omega/internal/core"
	"omega/pkg/auth"
	"omega/pkg/user"
)

func initUserAPI(e core.Engine) user.API {
	wire.Build(user.ProvideRepo, user.ProvideService, user.ProvideAPI)

	return user.API{}
}

func initAuthAPI(e core.Engine) auth.API {
	wire.Build(auth.ProvideRepo, auth.ProvideService, auth.ProvideAPI)

	return auth.API{}
}
