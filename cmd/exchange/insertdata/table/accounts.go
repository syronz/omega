package table

import (
	"omega/domain/base/basmodel"
	"omega/domain/base/basrepo"
	"omega/domain/base/enum/accountstatus"
	"omega/domain/base/enum/accounttype"
	"omega/domain/service"
	"omega/domain/sync"
	"omega/internal/core"
	"omega/internal/types"
	"omega/pkg/glog"
)

// InsertAccounts for add required accounts
func InsertAccounts(engine *core.Engine) {
	engine.DB.Exec("UPDATE bas_accounts SET deleted_at = null WHERE id IN (1,2,3)")
	accountRepo := basrepo.ProvideAccountRepo(engine)
	accountService := service.ProvideBasAccountService(accountRepo)
	accounts := []basmodel.Account{
		{
			FixedCol: types.FixedCol{
				ID:        1,
				CompanyID: engine.Envs.ToUint64(sync.CompanyID),
				NodeID:    engine.Envs.ToUint64(sync.NodeID),
			},
			Name:   "Asset",
			Type:   accounttype.Asset,
			Status: accountstatus.Active,
		},
		{
			FixedCol: types.FixedCol{
				ID:        2,
				CompanyID: engine.Envs.ToUint64(sync.CompanyID),
				NodeID:    engine.Envs.ToUint64(sync.NodeID),
			},
			Name:   "Capital",
			Type:   accounttype.Capital,
			Status: accountstatus.Active,
		},
		{
			FixedCol: types.FixedCol{
				ID:        3,
				CompanyID: engine.Envs.ToUint64(sync.CompanyID),
				NodeID:    engine.Envs.ToUint64(sync.NodeID),
			},
			Name:   "Cash",
			Type:   accounttype.Cash,
			Status: accountstatus.Active,
		},
		{
			FixedCol: types.FixedCol{
				ID:        4,
				CompanyID: engine.Envs.ToUint64(sync.CompanyID),
				NodeID:    engine.Envs.ToUint64(sync.NodeID),
			},
			Name:   "Equity",
			Type:   accounttype.Equity,
			Status: accountstatus.Active,
		},
	}

	for _, v := range accounts {
		if _, err := accountService.Save(v); err != nil {
			glog.Fatal("error in saving accounts", err)
		}
	}

}
