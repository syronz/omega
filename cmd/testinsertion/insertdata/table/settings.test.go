package table

import (
	"omega/cmd/restapi/enum/settingfields"
	"omega/domain/base/basmodel"
	"omega/domain/base/basrepo"
	"omega/domain/service"
	"omega/internal/core"
	"omega/internal/types"
	"omega/pkg/glog"
)

// InsertSettings for add required settings
func InsertSettings(engine *core.Engine) {
	settingRepo := basrepo.ProvideSettingRepo(engine)
	settingService := service.ProvideBasSettingService(settingRepo)

	// reset the table by deleting everything
	settingRepo.Engine.DB.Table(basmodel.SettingTable).Unscoped().Delete(basmodel.Setting{})

	settings := []basmodel.Setting{
		{
			GormCol: types.GormCol{
				ID: 1,
			},
			Property:    settingfields.CompanyName,
			Value:       "item",
			Type:        "string",
			Description: "company's name in the header of invoices",
		},
		{
			GormCol: types.GormCol{
				ID: 2,
			},
			Property:    settingfields.DefaultLang,
			Value:       "ku",
			Type:        "string",
			Description: "in case of user JWT not specified this value has been used",
		},
		{
			GormCol: types.GormCol{
				ID: 3,
			},
			Property:    settingfields.CompanyLogo,
			Value:       "invoice",
			Type:        "string",
			Description: "searchTerm1, logo for showed on the application and not invoices",
		},
		{
			GormCol: types.GormCol{
				ID: 4,
			},
			Property:    settingfields.InvoiceLogo,
			Value:       "public/logo.png",
			Type:        "string",
			Description: "path of logo, if branch logo won’t defined use this logo for invoices, searchTerm1",
		},
		{
			GormCol: types.GormCol{
				ID: 5,
			},
			Property:    settingfields.InvoiceNumberPattern,
			Value:       "location_year_series",
			Type:        "string",
			Description: "location_year_series, location_series, series, year_series, fullyear_series, location_fullyear_series, searchTerm1",
		},
	}

	for _, v := range settings {
		if _, err := settingService.Save(v); err != nil {
			glog.Fatal("error in inserting settings", err)
		}
	}

}
