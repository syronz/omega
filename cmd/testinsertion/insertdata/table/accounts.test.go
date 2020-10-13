package table

import (
	"omega/domain/base/basmodel"
	"omega/domain/base/basrepo"
	"omega/domain/base/enum/accounttype"
	"omega/domain/service"
	"omega/domain/sync/accountstatus"
	"omega/internal/core"
	"omega/internal/types"
	"omega/pkg/glog"
)

// InsertAccounts for add required accounts
func InsertAccounts(engine *core.Engine) {
	accountRepo := basrepo.ProvideAccountRepo(engine)
	accountService := service.ProvideBasAccountService(accountRepo)

	// reset the accounts table
	// reset in the roles.test.go

	accounts := []basmodel.Account{
		{
			GormCol: types.GormCol{
				ID: 1,
			},
			Name:   "Fee",
			Type:   accounttype.Fee,
			Status: accountstatus.Active,
		},
		{
			GormCol: types.GormCol{
				ID: 2,
			},
			Name:   "Trader",
			Type:   accounttype.Fixer,
			Status: accountstatus.Active,
		},
		{
			GormCol: types.GormCol{
				ID: 3,
			},
			Name:   "Gold Provider",
			Type:   accounttype.Provider,
			Status: accountstatus.Inactive,
		},
		{
			GormCol: types.GormCol{
				ID: 4,
			},
			Name:   "for foreign 1",
			Type:   accounttype.Trader,
			Status: accountstatus.Active,
		},
		{
			GormCol: types.GormCol{
				ID: 5,
			},
			Name:   "for update 1",
			Type:   accounttype.Trader,
			Status: accountstatus.Inactive,
		},
		{
			GormCol: types.GormCol{
				ID: 6,
			},
			Name:   "for update 2",
			Type:   accounttype.Trader,
			Status: accountstatus.Active,
		},
		{
			GormCol: types.GormCol{
				ID: 7,
			},
			Name:   "for delete 1",
			Type:   accounttype.Trader,
			Status: accountstatus.Active,
		},
		{
			GormCol: types.GormCol{
				ID: 8,
			},
			Name:   "for search 1, searchTerm1",
			Type:   accounttype.Trader,
			Status: accountstatus.Active,
		},
		{
			GormCol: types.GormCol{
				ID: 9,
			},
			Name:   "for search 2, searchTerm1",
			Type:   accounttype.Trader,
			Status: accountstatus.Active,
		},
		{
			GormCol: types.GormCol{
				ID: 10,
			},
			Name:   "for search 3, searchTerm1",
			Type:   accounttype.Trader,
			Status: accountstatus.Active,
		},
		{
			GormCol: types.GormCol{
				ID: 21,
			},
			Name:   "for delete 2",
			Type:   accounttype.Trader,
			Status: accountstatus.Active,
		},
		{
			GormCol: types.GormCol{
				ID: 30,
			},
			Name:   "active provider",
			Type:   accounttype.Trader,
			Status: accountstatus.Active,
		},
		{
			GormCol: types.GormCol{
				ID: 31,
			},
			Name:   "active follower",
			Type:   accounttype.Trader,
			Status: accountstatus.Active,
		},
		{
			GormCol: types.GormCol{
				ID: 32,
			},
			Name:   "inactive provider",
			Type:   accounttype.Trader,
			Status: accountstatus.Inactive,
		},
		{
			GormCol: types.GormCol{
				ID: 33,
			},
			Name:   "inactive follower",
			Type:   accounttype.Trader,
			Status: accountstatus.Inactive,
		},
	}

	for _, v := range accounts {
		if _, err := accountService.Save(v); err != nil {
			glog.Fatal(err)
		}
	}

}
