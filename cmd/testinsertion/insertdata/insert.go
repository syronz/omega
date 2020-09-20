package insertdata

import (
	// "omega/cmd/testinsertion/insertdata/table"
	"omega/cmd/testinsertion/insertdata/table"
	"omega/internal/core"
)

// Insert is used for add static rows to database
func Insert(engine *core.Engine) {

	if engine.Envs.ToBool(core.AutoMigrate) {
		table.InsertSettings(engine)
		table.InsertRoles(engine)

	}

}
