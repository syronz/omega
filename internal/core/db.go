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
	}

	// db.Model(&user.User{}).Related(&role.Role{})

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
