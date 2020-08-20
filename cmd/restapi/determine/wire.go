// +build wireinject

package determine

import (
	"omega/domain/base/basapi"
	"omega/domain/base/basrepo"
	"omega/domain/service"

	"omega/internal/core"

	"github.com/google/wire"
)

// Base Domain
func initBasSettingAPI(e *core.Engine) basapi.BasSettingAPI {
	wire.Build(basrepo.ProvideBasSettingRepo, service.ProvideBasSettingService,
		basapi.ProvideBasSettingAPI)
	return basapi.BasSettingAPI{}
}

func initBasRoleAPI(e *core.Engine) basapi.BasRoleAPI {
	wire.Build(basrepo.ProvideBasRoleRepo, service.ProvideBasRoleService,
		basapi.ProvideBasRoleAPI)
	return basapi.BasRoleAPI{}
}

func initBasUserAPI(engine *core.Engine) basapi.BasUserAPI {
	wire.Build(basrepo.ProvideBasUserRepo, service.ProvideBasUserService, basapi.ProvideBasUserAPI)
	return basapi.BasUserAPI{}
}

func initBasAuthAPI(e *core.Engine) basapi.BasAuthAPI {
	wire.Build(service.ProvideBasAuthService, basapi.ProvideBasAuthAPI)
	return basapi.BasAuthAPI{}
}

func initBasActivityAPI(engine *core.Engine) basapi.BasActivityAPI {
	wire.Build(basrepo.ProvideBasActivityRepo, service.ProvideBasActivityService, basapi.ProvideBasActivityAPI)
	return basapi.BasActivityAPI{}
}
