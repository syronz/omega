package server

import (
	"net/http"
	"omega/domain/base"
	"omega/domain/base/basmid"
	"omega/domain/eaccounting"
	"omega/domain/material"
	"omega/internal/core"

	"github.com/gin-gonic/gin"
)

// Route trigger router and api methods
func Route(rg gin.RouterGroup, engine *core.Engine) {
	// Base Domain
	basAuthAPI := initAuthAPI(engine)
	basUserAPI := initUserAPI(engine)
	basRoleAPI := initRoleAPI(engine)
	basSettingAPI := initSettingAPI(engine)
	basActivityAPI := initActivityAPI(engine)
	basAccountAPI := initAccountAPI(engine)

	// EAccountig Domain
	eacCurrencyAPI := initCurrencyAPI(engine)
	eacSlotAPI := initSlotAPI(engine, eacCurrencyAPI.Service, basAccountAPI.Service)
	eacTransactionAPI := initTransactionAPI(engine, eacSlotAPI.Service)

	// Material Domain
	matCompanyAPI := initMatCompanyAPI(engine)

	// Html Domain
	rg.StaticFS("/public", http.Dir("public"))

	rg.POST("/login", basAuthAPI.Login)
	rg.POST("/register", basAuthAPI.Register)

	rg.Use(basmid.AuthGuard(engine))

	access := basmid.NewAccessMid(engine)

	rg.POST("/logout", basAuthAPI.Logout)

	// Base Domain
	rg.GET("/temporary/token", basAuthAPI.TemporaryToken)

	rg.GET("/companies/:companyID/settings",
		access.Check(base.SettingRead), basSettingAPI.List)
	rg.GET("/companies/:companyID/nodes/:nodeID/settings/:settingID",
		access.Check(base.SettingRead), basSettingAPI.FindByID)
	rg.PUT("/companies/:companyID/nodes/:nodeID/settings/:settingID",
		access.Check(base.SettingWrite), basSettingAPI.Update)
	rg.GET("/excel/companies/:companyID/settings",
		access.Check(base.SettingExcel), basSettingAPI.Excel)

	rg.GET("/companies/:companyID/roles",
		access.Check(base.RoleRead), basRoleAPI.List)
	rg.GET("/companies/:companyID/nodes/:nodeID/roles/:roleID",
		access.Check(base.RoleRead), basRoleAPI.FindByID)
	rg.POST("/companies/:companyID/roles",
		access.Check(base.RoleWrite), basRoleAPI.Create)
	rg.PUT("/companies/:companyID/nodes/:nodeID/roles/:roleID",
		access.Check(base.RoleWrite), basRoleAPI.Update)
	rg.DELETE("companies/:companyID/nodes/:nodeID/roles/:roleID",
		access.Check(base.RoleWrite), basRoleAPI.Delete)
	rg.GET("/excel/companies/:companyID/roles",
		access.Check(base.RoleExcel), basRoleAPI.Excel)

	rg.GET("/companies/:companyID/accounts",
		access.Check(base.AccountRead), basAccountAPI.List)
	rg.GET("/companies/:companyID/nodes/:nodeID/accounts/:accountID",
		access.Check(base.AccountRead), basAccountAPI.FindByID)
	rg.POST("/companies/:companyID/accounts",
		access.Check(base.AccountWrite), basAccountAPI.Create)
	rg.PUT("/companies/:companyID/nodes/:nodeID/accounts/:accountID",
		access.Check(base.AccountWrite), basAccountAPI.Update)
	rg.DELETE("/companies/:companyID/nodes/:nodeID/accounts/:accountID",
		access.Check(base.AccountWrite), basAccountAPI.Delete)
	rg.GET("/excel/companies/:companyID/accounts",
		access.Check(base.AccountExcel), basAccountAPI.Excel)

	rg.GET("/username/:username",
		access.Check(base.UserRead), basUserAPI.FindByUsername)
	rg.GET("/companies/:companyID/users",
		access.Check(base.UserRead), basUserAPI.List)
	rg.GET("/companies/:companyID/nodes/:nodeID/users/:userID",
		access.Check(base.UserRead), basUserAPI.FindByID)
	rg.POST("/companies/:companyID/users",
		access.Check(base.UserWrite), basUserAPI.Create)
	rg.PUT("/companies/:companyID/nodes/:nodeID/users/:userID",
		access.Check(base.UserWrite), basUserAPI.Update)
	rg.DELETE("/companies/:companyID/nodes/:nodeID/users/:userID",
		access.Check(base.UserWrite), basUserAPI.Delete)
	rg.GET("/excel/companies/:companyID/users",
		access.Check(base.UserExcel), basUserAPI.Excel)

	rg.GET("/activities",
		access.Check(base.ActivityAll), basActivityAPI.List)

	// EAccountig Domain
	rg.GET("/companies/:companyID/currencies",
		access.Check(eaccounting.CurrencyRead), eacCurrencyAPI.List)
	rg.GET("/companies/:companyID/nodes/:nodeID/currencies/:currencyID",
		access.Check(eaccounting.CurrencyRead), eacCurrencyAPI.FindByID)
	rg.POST("/companies/:companyID/currencies",
		access.Check(eaccounting.CurrencyWrite), eacCurrencyAPI.Create)
	rg.PUT("/companies/:companyID/nodes/:nodeID/currencies/:currencyID",
		access.Check(eaccounting.CurrencyWrite), eacCurrencyAPI.Update)
	rg.DELETE("/companies/:companyID/nodes/:nodeID/currencies/:currencyID",
		access.Check(eaccounting.CurrencyWrite), eacCurrencyAPI.Delete)
	rg.GET("/excel/companies/:companyID/currencies",
		access.Check(eaccounting.CurrencyExcel), eacCurrencyAPI.Excel)

	rg.GET("/companies/:companyID/transactions",
		access.Check(eaccounting.TransactionRead), eacTransactionAPI.List)
	rg.GET("/companies/:companyID/nodes/:nodeID/transactions/:transactionID",
		access.Check(eaccounting.TransactionRead), eacTransactionAPI.FindByID)
	rg.POST("/companies/:companyID/transactions",
		access.Check(eaccounting.TransactionManual), eacTransactionAPI.ManualTransfer)
	rg.PUT("/companies/:companyID/nodes/:nodeID/transactions/:transactionID",
		access.Check(eaccounting.TransactionUpdate), eacTransactionAPI.Update)
	rg.DELETE("/companies/:companyID/nodes/:nodeID/transactions/:transactionID",
		access.Check(eaccounting.TransactionDelete), eacTransactionAPI.Delete)
	rg.GET("/excel/companies/:companyID/transactions",
		access.Check(eaccounting.TransactionExcel), eacTransactionAPI.Excel)

	// Material Domain
	rg.GET("/companies/:companyID/companies",
		access.Check(material.CompanyRead), matCompanyAPI.List)
	rg.GET("/companies/:companyID/nodes/:nodeID/companies/:companyID",
		access.Check(material.CompanyRead), matCompanyAPI.FindByID)
	rg.POST("/companies/:companyID/companies",
		access.Check(material.CompanyWrite), matCompanyAPI.Create)
	rg.PUT("/companies/:companyID/nodes/:nodeID/companies/:companyID",
		access.Check(material.CompanyWrite), matCompanyAPI.Update)
	rg.DELETE("/companies/:companyID/nodes/:nodeID/companies/:companyID",
		access.Check(material.CompanyWrite), matCompanyAPI.Delete)
	rg.GET("/excel/companies/:companyID/companies",
		access.Check(material.CompanyExcel), matCompanyAPI.Excel)

}
