package core

import (
	"omega/engine"
	"omega/pkg/user"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func swapper(db *gorm.DB, i ...interface{}) {
	db.AutoMigrate(i...)
}

func passer(db *gorm.DB, i ...interface{}) {
	swapper(db, i...)
}

func initDB(e engine.Engine, dbType string, dsn string) *gorm.DB {
	db, err := gorm.Open(dbType, dsn)
	if err != nil {
		e.ServerLog.Fatalln(err)
	}

	db.LogMode(false)

	if gin.IsDebugging() {
		db.LogMode(true)
	}

	passer(db, &user.User{})

	if e.Environments.Setting.AutoMigrate {
		// usObject := swapper(&user.User{})
		// db.AutoMigrate(usObject)
		// db.AutoMigrate(&user.User{})
	}

	return db
}
