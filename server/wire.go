//+build wireinject

package server

import (
	"github.com/google/wire"
	"omega/engine"
	"omega/pkg/auth"
	"omega/pkg/user"
)

func initUserAPI(e engine.Engine) user.API {
	wire.Build(user.ProvideRepo, user.ProvideService, user.ProvideAPI)

	return user.API{}
}

func initAuthAPI(e engine.Engine) auth.API {
	wire.Build(auth.ProvideRepo, auth.ProvideService, auth.ProvideAPI)

	return auth.API{}
}
