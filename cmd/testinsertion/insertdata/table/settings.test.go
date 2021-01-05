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
	engine.DB.Exec("TRUNCATE TABLE bas_settings;")

	settings := []basmodel.Setting{
		{
			FixedCol: types.FixedCol{
				ID:        1,
				CompanyID: 1001,
				NodeID:    101,
			},
			Property:    settingfields.CompanyName,
			Value:       "item",
			Type:        "string",
			Description: "company's name in the header of invoices",
		},
		{
			FixedCol: types.FixedCol{
				ID:        2,
				CompanyID: 1001,
				NodeID:    101,
			},
			Property:    settingfields.DefaultLang,
			Value:       "ku",
			Type:        "string",
			Description: "in case of user JWT not specified this value has been used",
		},
		{
			FixedCol: types.FixedCol{
				ID:        3,
				CompanyID: 1001,
				NodeID:    101,
			},
			Property:    settingfields.CompanyLogo,
			Value:       "invoice",
			Type:        "string",
			Description: "searchTerm1, logo for showed on the application and not invoices",
		},
		{
			FixedCol: types.FixedCol{
				ID:        4,
				CompanyID: 1001,
				NodeID:    101,
			},
			Property:    settingfields.InvoiceLogo,
			Value:       "public/logo.png",
			Type:        "string",
			Description: "path of logo, if branch logo won’t defined use this logo for invoices, searchTerm1",
		},
		{
			FixedCol: types.FixedCol{
				ID:        5,
				CompanyID: 1001,
				NodeID:    101,
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
