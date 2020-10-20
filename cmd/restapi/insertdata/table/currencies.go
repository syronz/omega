package table

import (
	"omega/domain/eaccounting/eacmodel"
	"omega/domain/eaccounting/eacrepo"
	"omega/domain/service"
	"omega/domain/sync"
	"omega/internal/core"
	"omega/internal/types"
	"omega/pkg/glog"
)

// InsertCurrencies for add required currencies
func InsertCurrencies(engine *core.Engine) {
	engine.DB.Exec("UPDATE eac_currencies SET deleted_at = null WHERE id IN (1,2)")
	currencyRepo := eacrepo.ProvideCurrencyRepo(engine)
	currencyService := service.ProvideEacCurrencyService(currencyRepo)
	currencies := []eacmodel.Currency{
		{
			FixedCol: types.FixedCol{
				ID:        1,
				CompanyID: engine.Envs.ToUint64(sync.CompanyID),
				NodeID:    engine.Envs.ToUint64(sync.NodeID),
			},
			Name:   "Dollar",
			Symbol: "$",
			Code:   "USD",
		},
		{
			FixedCol: types.FixedCol{
				ID:        2,
				CompanyID: engine.Envs.ToUint64(sync.CompanyID),
				NodeID:    engine.Envs.ToUint64(sync.NodeID),
			},
			Name:   "Dinar",
			Symbol: "IQD",
			Code:   "IQD",
		},
	}

	for _, v := range currencies {
		if _, err := currencyService.Save(v); err != nil {
			glog.Fatal("error in saving currencies", err)
		}
	}

}
