package table

import (
	"omega/domain/base/basmodel"
	"omega/domain/base/basrepo"
	"omega/domain/base/enum/accounttype"
	"omega/domain/service"
	"omega/domain/sync/accountstatus"
	"omega/internal/consts"
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
			GormCol: types.GormCol{
				ID: consts.AccFeeID,
			},
			Name:   "Fee",
			Type:   accounttype.Fee,
			Status: accountstatus.Active,
		},
		{
			GormCol: types.GormCol{
				ID: consts.ACCTraderID,
			},
			Name:   "Lock",
			Type:   accounttype.Fixer,
			Status: accountstatus.Active,
		},
		{
			GormCol: types.GormCol{
				ID: 3,
			},
			Name:   "Gold to Provider",
			Type:   accounttype.Fixer,
			Status: accountstatus.Active,
		},
		{
			GormCol: types.GormCol{
				ID: 4,
			},
			Name:   "Cash",
			Type:   accounttype.Asset,
			Status: accountstatus.Inactive,
		},
	}

	for _, v := range accounts {
		if _, err := accountService.Save(v); err != nil {
			glog.Fatal("error in saving accounts", err)
		}
	}

}
