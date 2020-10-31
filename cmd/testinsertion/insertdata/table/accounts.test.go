package table

import (
	"omega/domain/base/basmodel"
	"omega/domain/base/basrepo"
	"omega/domain/base/enum/accountstatus"
	"omega/domain/base/enum/accounttype"
	"omega/domain/service"
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
			FixedNode: types.FixedNode{
				ID:        1,
				CompanyID: 1001,
				NodeID:    101,
			},
			Name:   "Asset",
			Type:   accounttype.Asset,
			Status: accountstatus.Active,
		},
		{
			FixedNode: types.FixedNode{
				ID:        2,
				CompanyID: 1001,
				NodeID:    101,
			},
			Name:   "Capital",
			Type:   accounttype.Capital,
			Status: accountstatus.Active,
		},
		{
			FixedNode: types.FixedNode{
				ID:        3,
				CompanyID: 1001,
				NodeID:    101,
			},
			Name:   "Cash",
			Type:   accounttype.Cash,
			Status: accountstatus.Inactive,
		},
		{
			FixedNode: types.FixedNode{
				ID:        4,
				CompanyID: 1001,
				NodeID:    101,
			},
			Name:   "for foreign 1",
			Type:   accounttype.Equity,
			Status: accountstatus.Active,
		},
		{
			FixedNode: types.FixedNode{
				ID:        5,
				CompanyID: 1001,
				NodeID:    101,
			},
			Name:   "for update 1",
			Type:   accounttype.Partner,
			Status: accountstatus.Inactive,
		},
		{
			FixedNode: types.FixedNode{
				ID:        6,
				CompanyID: 1001,
				NodeID:    101,
			},
			Name:   "for update 2",
			Type:   accounttype.Partner,
			Status: accountstatus.Active,
		},
		{
			FixedNode: types.FixedNode{
				ID:        7,
				CompanyID: 1001,
				NodeID:    101,
			},
			Name:   "for delete 1",
			Type:   accounttype.Partner,
			Status: accountstatus.Active,
		},
		{
			FixedNode: types.FixedNode{
				ID:        8,
				CompanyID: 1001,
				NodeID:    101,
			},
			Name:   "for search 1, searchTerm1",
			Type:   accounttype.Partner,
			Status: accountstatus.Active,
		},
		{
			FixedNode: types.FixedNode{
				ID:        9,
				CompanyID: 1001,
				NodeID:    101,
			},
			Name:   "for search 2, searchTerm1",
			Type:   accounttype.Partner,
			Status: accountstatus.Active,
		},
		{
			FixedNode: types.FixedNode{
				ID:        10,
				CompanyID: 1001,
				NodeID:    101,
			},
			Name:   "for search 3, searchTerm1",
			Type:   accounttype.Partner,
			Status: accountstatus.Active,
		},
		{
			FixedNode: types.FixedNode{
				ID:        21,
				CompanyID: 1001,
				NodeID:    101,
			},
			Name:   "for delete 2",
			Type:   accounttype.Partner,
			Status: accountstatus.Active,
		},
		{
			FixedNode: types.FixedNode{
				ID:        30,
				CompanyID: 1001,
				NodeID:    101,
			},
			Name:   "active provider",
			Type:   accounttype.Partner,
			Status: accountstatus.Active,
		},
		{
			FixedNode: types.FixedNode{
				ID:        31,
				CompanyID: 1001,
				NodeID:    101,
			},
			Name:   "A",
			Type:   accounttype.Partner,
			Status: accountstatus.Active,
		},
		{
			FixedNode: types.FixedNode{
				ID:        32,
				CompanyID: 1001,
				NodeID:    101,
			},
			Name:   "B",
			Type:   accounttype.Partner,
			Status: accountstatus.Active,
		},
		{
			FixedNode: types.FixedNode{
				ID:        33,
				CompanyID: 1001,
				NodeID:    101,
			},
			Name:   "C",
			Type:   accounttype.Partner,
			Status: accountstatus.Active,
		},
		{
			FixedNode: types.FixedNode{
				ID:        34,
				CompanyID: 1001,
				NodeID:    101,
			},
			Name:   "D",
			Type:   accounttype.Partner,
			Status: accountstatus.Active,
		},
	}

	for _, v := range accounts {
		if _, err := accountService.Save(v); err != nil {
			glog.Fatal(err)
		}
	}

}
