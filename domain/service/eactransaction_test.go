package service

import (
	"omega/domain/base/basrepo"
	"omega/domain/eaccounting/eacmodel"
	"omega/domain/eaccounting/eacrepo"
	"omega/domain/eaccounting/enum/transactiontype"
	"omega/internal/consts"
	"omega/internal/core"
	"omega/internal/types"
	"omega/pkg/helper"
	"omega/test/kernel"
	"testing"
	"time"
)

func initTransactionTest() (engine *core.Engine, transactionServ EacTransactionServ) {
	logQuery, debugLevel := initServiceTest()
	engine = kernel.StartMotor(logQuery, debugLevel)

	phoneServ := ProvideBasPhoneService(basrepo.ProvidePhoneRepo(engine))
	accountServ := ProvideBasAccountService(basrepo.ProvideAccountRepo(engine), phoneServ)
	currencyServ := ProvideEacCurrencyService(eacrepo.ProvideCurrencyRepo(engine))
	slotServ := ProvideEacSlotService(eacrepo.ProvideSlotRepo(engine), currencyServ, accountServ)
	transactionServ = ProvideEacTransactionService(eacrepo.ProvideTransactionRepo(engine), slotServ)

	return
}

func TestTransactionTransfer(t *testing.T) {
	_, transactionServ := initTransactionTest()
	// time1, err := time.Parse(consts.TimeLayout, "2020-10-20 15:10:00")
	time1, err := time.Parse(consts.TimeLayoutZone, "2020-10-20 15:10:00 +0300")
	if err != nil {
		t.Errorf("error in parsing date %v in layout %v", consts.DefaultLimit, "2020-10-21 21:10:35")
	}

	samples := []struct {
		in  eacmodel.Transaction
		err error
	}{
		{in: eacmodel.Transaction{
			FixedNode: types.FixedNode{
				CompanyID: 1001,
				NodeID:    101,
			},
			Type:       transactiontype.Manual,
			CreatedBy:  11,
			Pioneer:    31,
			Follower:   32,
			CurrencyID: 1,
			Amount:     1000,
			PostDate:   time1,
		},
			err: nil,
		},
	}

	for _, v := range samples {
		_, err := transactionServ.Transfer(v.in)
		if (v.err == nil && err != nil) || (v.err != nil && err == nil) {
			t.Errorf("\nERROR FOR :::%+v::: \nRETURNS :::%+v:::, \nIT SHOULD BE :::%+v:::", v.in, err, v.err)
		}
	}

}

func TestTransactionEditManual(t *testing.T) {
	_, transactionServ := initTransactionTest()
	time1, err := time.Parse(consts.TimeLayoutZone, "2020-10-19 15:10:00 +0300")
	if err != nil {
		t.Errorf("error in parsing date %v in layout %v", consts.DefaultLimit, "2020-10-21 21:10:35")
	}

	samples := []struct {
		in  eacmodel.Transaction
		err error
	}{
		{
			in: eacmodel.Transaction{
				FixedNode: types.FixedNode{
					CompanyID: 1001,
					NodeID:    101,
					ID:        1,
				},
				Pioneer:     33,
				Follower:    32,
				CurrencyID:  1,
				Amount:      500,
				PostDate:    time1,
				Description: helper.StrPointer("changed!, A -> C & 1000$ -> 500$"),
				Type:        transactiontype.Manual,
				CreatedBy:   11,
			},
			err: nil,
		},
	}

	for _, v := range samples {
		_, err := transactionServ.EditTransfer(v.in)
		if (v.err == nil && err != nil) || (v.err != nil && err == nil) {
			t.Errorf("\nERROR FOR :::%+v::: \nRETURNS :::%+v:::, \nIT SHOULD BE :::%+v:::", v.in, err, v.err)
		}
	}

}
