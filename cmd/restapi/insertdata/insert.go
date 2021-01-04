package insertdata

import (
	"omega/internal/core"
)

// Insert is used for add static rows to database
func Insert(engine *core.Engine) {

	if engine.Envs.ToBool(core.AutoMigrate) {
		// table.InsertCompanys(engine)
		// table.InsertRoles(engine)
		// table.InsertAccounts(engine)
		// table.InsertUsers(engine)
		// table.InsertSettings(engine)

		// table.InsertCurrencies(engine)
	}

}
