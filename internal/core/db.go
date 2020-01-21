package core

import (
	"omega/engine"
	"omega/internal/models"
	"omega/pkg/role"
	"omega/pkg/user"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func initDataDB(e engine.Engine, dbType string, dsn string) *gorm.DB {
	db, err := gorm.Open(dbType, dsn)
	if err != nil {
		e.ServerLog.Fatalln(err)
	}

	db.LogMode(false)

	if gin.IsDebugging() {
		db.LogMode(true)
	}

	if e.Environments.Setting.AutoMigrate {
		db.AutoMigrate(&role.Role{})
		db.AutoMigrate(&user.User{}).AddForeignKey("role_id", "roles(id)", "RESTRICT", "RESTRICT")

		roleRepo := role.ProvideRepo(e)
		roleService := role.ProvideService(roleRepo)
		roleObj := role.Role{
			Name:        "Admin",
			Resources:   "users:read users:write users:report activities:self activities:all roles:read roles:write",
			Description: "admin has all privileges",
		}

		_, _ = roleService, roleObj
		_, _ = roleService.Save(roleObj)

	}

	return db
}

func initActivityDB(e engine.Engine, dbType string, dsn string) *gorm.DB {
	db, err := gorm.Open(dbType, dsn)
	if err != nil {
		e.ServerLog.Fatalln(err)
	}

	db.LogMode(false)

	if gin.IsDebugging() {
		db.LogMode(true)
	}

	if e.Environments.Setting.AutoMigrate {
		db.AutoMigrate(&models.Activity{})
	}

	return db
}
