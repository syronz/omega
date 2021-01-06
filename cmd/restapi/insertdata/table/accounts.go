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
	engine.DB.Exec("UPDATE bas_accounts SET deleted_at = null WHERE id IN (1,2,3,4,5)")
	phoneServ := service.ProvideBasPhoneService(basrepo.ProvidePhoneRepo(engine))
	accountRepo := basrepo.ProvideAccountRepo(engine)
	accountService := service.ProvideBasAccountService(accountRepo, phoneServ)
	accounts := []basmodel.Account{
		{
			FixedNode: types.FixedNode{
				ID:        1,
				CompanyID: engine.Envs.ToUint64(sync.CompanyID),
				NodeID:    engine.Envs.ToUint64(sync.NodeID),
			},
			Name:   "Asset",
			Code:   "1",
			Type:   accounttype.Asset,
			Status: accountstatus.Active,
		},
		{
			FixedNode: types.FixedNode{
				ID:        2,
				CompanyID: engine.Envs.ToUint64(sync.CompanyID),
				NodeID:    engine.Envs.ToUint64(sync.NodeID),
			},
			Name:     "Cash",
			ParentID: types.RowIDPointer(1),
			Code:     "11",
			Type:     accounttype.Cash,
			Status:   accountstatus.Active,
		},
		{
			FixedNode: types.FixedNode{
				ID:        3,
				CompanyID: engine.Envs.ToUint64(sync.CompanyID),
				NodeID:    engine.Envs.ToUint64(sync.NodeID),
			},
			Name:     "Users",
			ParentID: types.RowIDPointer(1),
			Code:     "12",
			Type:     accounttype.User,
			Status:   accountstatus.Active,
		},
		{
			FixedNode: types.FixedNode{
				ID:        4,
				CompanyID: engine.Envs.ToUint64(sync.CompanyID),
				NodeID:    engine.Envs.ToUint64(sync.NodeID),
			},
			Name:   "Equity",
			Code:   "2",
			Type:   accounttype.Equity,
			Status: accountstatus.Active,
		},
		{
			FixedNode: types.FixedNode{
				ID:        5,
				CompanyID: engine.Envs.ToUint64(sync.CompanyID),
				NodeID:    engine.Envs.ToUint64(sync.NodeID),
			},
			Name:     "Capital",
			ParentID: types.RowIDPointer(4),
			Code:     "21",
			Type:     accounttype.Capital,
			Status:   accountstatus.Active,
		},
	}

	for _, v := range accounts {
		if _, err := accountService.Save(v); err != nil {
			glog.Fatal("error in saving accounts", err)
		}
	}

}
