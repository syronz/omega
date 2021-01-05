package table

import (
	"omega/domain/base/basrepo"
	"omega/domain/eaccounting/eacmodel"
	"omega/domain/eaccounting/eacrepo"
	"omega/domain/eaccounting/enum/transactiontype"
	"omega/domain/service"
	"omega/internal/consts"
	"omega/internal/core"
	"omega/internal/types"
	"omega/pkg/glog"
	"omega/pkg/helper"
	"time"
)

// InsertTransactions for add required accounts
func InsertTransactions(engine *core.Engine) {
	phoneServ := service.ProvideBasPhoneService(basrepo.ProvidePhoneRepo(engine))
	accountRepo := basrepo.ProvideAccountRepo(engine)
	accountService := service.ProvideBasAccountService(accountRepo, phoneServ)

	currencyRepo := eacrepo.ProvideCurrencyRepo(engine)
	currencyService := service.ProvideEacCurrencyService(currencyRepo)

	slotRepo := eacrepo.ProvideSlotRepo(engine)
	slotService := service.ProvideEacSlotService(slotRepo, currencyService, accountService)

	transactionRepo := eacrepo.ProvideTransactionRepo(engine)
	transactionService := service.ProvideEacTransactionService(transactionRepo, slotService)

	time1, _ := time.Parse(consts.TimeLayoutZone, "2020-10-19 15:10:00 +0300")

	transactions := []eacmodel.Transaction{
		{ // A -- 1000$ -- > B
			FixedNode: types.FixedNode{
				CompanyID: 1001,
				NodeID:    101,
			},
			Pioneer:     31,
			Follower:    32,
			CurrencyID:  1,
			Amount:      1000,
			PostDate:    time1,
			Description: helper.StrPointer("A -> B : 1000$"),
			Type:        transactiontype.Manual,
			CreatedBy:   11,
		},
		{ // A -- 800$ -- > B
			FixedNode: types.FixedNode{
				CompanyID: 1001,
				NodeID:    101,
			},
			Pioneer:     31,
			Follower:    32,
			CurrencyID:  1,
			Amount:      800,
			PostDate:    time1,
			Description: helper.StrPointer("A -> B : 800$"),
			Type:        transactiontype.Manual,
			CreatedBy:   11,
		},
		{ // C -- 200$ -- > B
			FixedNode: types.FixedNode{
				CompanyID: 1001,
				NodeID:    101,
			},
			Pioneer:     33,
			Follower:    32,
			CurrencyID:  1,
			Amount:      200,
			PostDate:    time1,
			Description: helper.StrPointer("C -> B : 200$"),
			Type:        transactiontype.Manual,
			CreatedBy:   11,
		},
		{ // D -- 300$ -- > A
			FixedNode: types.FixedNode{
				CompanyID: 1001,
				NodeID:    101,
			},
			Pioneer:     34,
			Follower:    31,
			CurrencyID:  1,
			Amount:      300,
			PostDate:    time1,
			Description: helper.StrPointer("D -> A : 300$"),
			Type:        transactiontype.Manual,
			CreatedBy:   11,
		},
	}

	for _, v := range transactions {
		if _, err := transactionService.Transfer(v); err != nil {
			glog.Fatal(err)
		}
	}

}
