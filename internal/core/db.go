package core

import (
	"omega/engine"
	"omega/pkg/user"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func initDB(e engine.Engine, dbType string, dsn string) *gorm.DB {
	db, err := gorm.Open(dbType, dsn)
	if err != nil {
		e.ServerLog.Fatalln(err)
	}

	db.LogMode(false)

	if gin.IsDebugging() {
		db.LogMode(true)
	}

	if e.Environments.Setting.AutoMigrate {
		db.AutoMigrate(&user.User{})
	}

	return db
}
