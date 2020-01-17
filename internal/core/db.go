package core

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"omega/engine"
	"omega/pkg/user"
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

	if e.Environments.Setting.AutoMigrate == true {
		db.AutoMigrate(&user.User{})
	}

	return db
}
