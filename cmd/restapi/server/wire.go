// +build wireinject

package server

import (
	"omega/domain/base/basapi"
	"omega/domain/base/basrepo"
	"omega/domain/eaccounting/eacapi"
	"omega/domain/eaccounting/eacrepo"
	"omega/domain/material/matapi"
	"omega/domain/material/matrepo"
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

// EAccountig Domain
func initCurrencyAPI(e *core.Engine) eacapi.CurrencyAPI {
	wire.Build(eacrepo.ProvideCurrencyRepo, service.ProvideEacCurrencyService,
		eacapi.ProvideCurrencyAPI)
	return eacapi.CurrencyAPI{}
}

func initTransactionAPI(e *core.Engine, slotServ service.EacSlotServ) eacapi.TransactionAPI {
	wire.Build(eacrepo.ProvideTransactionRepo, service.ProvideEacTransactionService,
		eacapi.ProvideTransactionAPI)
	return eacapi.TransactionAPI{}
}

func initSlotAPI(e *core.Engine, currencyServ service.EacCurrencyServ,
	accountServ service.BasAccountServ) eacapi.SlotAPI {
	wire.Build(eacrepo.ProvideSlotRepo, service.ProvideEacSlotService,
		eacapi.ProvideSlotAPI)
	return eacapi.SlotAPI{}
}

// Material Domain
func initMatCompanyAPI(e *core.Engine) matapi.CompanyAPI {
	wire.Build(matrepo.ProvideCompanyRepo, service.ProvideMatCompanyService,
		matapi.ProvideCompanyAPI)
	return matapi.CompanyAPI{}
}
