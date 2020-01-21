package migrate

import (
	"omega/engine"
)

// InsertData is used for inserting data to the tables
func InsertData(e engine.Engine) {

	if e.Environments.Setting.AutoMigrate {
		migrateRoles(e)
		migrateUsers(e)
	}
}
