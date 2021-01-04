package table

import (
	"omega/domain/service"
	"omega/domain/sync/synmodel"
	"omega/domain/sync/synrepo"
	"omega/internal/core"
	"omega/internal/types"
	"omega/pkg/glog"
	"time"
)

// InsertCompanys for add required companys
func InsertCompanys(engine *core.Engine) {
	expiryDate, _ := time.Parse("2006-01-02", "2030-01-01")
	companyRepo := synrepo.ProvideCompanyRepo(engine)
	companyService := service.ProvideSynCompanyService(companyRepo)
	companys := []synmodel.Company{
		{
			GormCol: types.GormCol{
				ID: 1001,
			},
			Name:          "Base",
			LegalName:     "Base",
			ServerAddress: "127.0.0.1",
			Expiration:    &expiryDate,
			License:       "10 Years",
			Plan:          "Base",
			Detail:        "generated when program initiated",
			Phone:         "07505149171",
			Email:         "lab@iqonline.com",
			Website:       "lab.iqonline.com",
			Type:          "base",
			Code:          "1001",
		},
	}

	for _, v := range companys {
		if _, err := companyService.Save(v); err != nil {
			glog.Fatal("Disable in production: error in saving companys", err)
		}
	}

}
