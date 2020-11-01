package insertdata

import (
	"omega/cmd/testinsertion/insertdata/table"
	"omega/internal/core"
)

// Insert is used for add static rows to database
func Insert(engine *core.Engine) {

	if engine.Envs.ToBool(core.AutoMigrate) {
		table.InsertSettings(engine)
		table.InsertRoles(engine)
		table.InsertUsers(engine)
		table.InsertAccounts(engine)
		table.InsertCurrencies(engine)
		table.InsertTransactions(engine)
	}

}
