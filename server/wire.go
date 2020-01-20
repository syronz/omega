//+build wireinject

package server

import (
	"github.com/google/wire"
	"omega/engine"
	"omega/pkg/activity"
	"omega/pkg/auth"
	"omega/pkg/role"
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

func initActivityAPI(e engine.Engine) activity.API {
	wire.Build(activity.ProvideRepo, activity.ProvideService, activity.ProvideAPI)
	return activity.API{}
}

func initRoleAPI(e engine.Engine) role.API {
	wire.Build(role.ProvideRepo, role.ProvideService, role.ProvideAPI)
	return role.API{}
}
