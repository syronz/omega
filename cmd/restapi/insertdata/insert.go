package insertdata

import (
	"omega/cmd/restapi/insertdata/table"
	"omega/internal/core"
)

// Insert is used for add static rows to database
func Insert(engine *core.Engine) {

	if engine.Envs.ToBool(core.AutoMigrate) {
		table.InsertRoles(engine)
		table.InsertUsers(engine)
		table.InsertSettings(engine)
	}

}
