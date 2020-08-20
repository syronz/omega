package insertdata

import (
	"omega/cmd/restapi/insertdata/table"
	"omega/internal/core"
)

// Insert is used for add static rows to database
func Insert(engine *core.Engine) {

	if engine.Envs.ToBool(core.AutoMigrate) {
		table.InsertBasRoles(engine)
		table.InsertBasUsers(engine)
		table.InsertBasSettings(engine)
	}

}
