// +build wireinject

package server

import (
	"omega/domain/base/basapi"
	"omega/domain/base/basrepo"
	"omega/domain/html/htmapi"
	"omega/domain/service"

	"omega/internal/core"

	"github.com/google/wire"
)

// Base Domain
func initSettingAPI(e *core.Engine) basapi.SettingAPI {
	wire.Build(basrepo.ProvideSettingRepo, service.ProvideBasSettingService,
		basapi.ProvideSettingAPI)
	return basapi.SettingAPI{}
}

func initRoleAPI(e *core.Engine) basapi.RoleAPI {
	wire.Build(basrepo.ProvideRoleRepo, service.ProvideBasRoleService,
		basapi.ProvideRoleAPI)
	return basapi.RoleAPI{}
}

func initUserAPI(engine *core.Engine) basapi.UserAPI {
	wire.Build(basrepo.ProvideUserRepo, service.ProvideBasUserService, basapi.ProvideUserAPI)
	return basapi.UserAPI{}
}

func initAuthAPI(e *core.Engine) basapi.AuthAPI {
	wire.Build(service.ProvideBasAuthService, basapi.ProvideAuthAPI)
	return basapi.AuthAPI{}
}

func initActivityAPI(engine *core.Engine) basapi.ActivityAPI {
	wire.Build(basrepo.ProvideActivityRepo, service.ProvideBasActivityService, basapi.ProvideActivityAPI)
	return basapi.ActivityAPI{}
}

func initAccountAPI(e *core.Engine) basapi.AccountAPI {
	wire.Build(basrepo.ProvideAccountRepo, service.ProvideBasAccountService,
		basapi.ProvideAccountAPI)
	return basapi.AccountAPI{}
}

// Html Domain
func initErrDescAPI(e *core.Engine) htmapi.ErrDescAPI {
	wire.Build(htmapi.GenErrDescAPI)
	return htmapi.ErrDescAPI{}
}
