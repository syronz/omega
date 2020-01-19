package core

import (
	"omega/engine"

	"github.com/jinzhu/gorm"
)

func initDB(e engine.Engine, dbType string, dsn string, models ...interface{}) *gorm.DB {
	db, err := gorm.Open(dbType, dsn)
	if err != nil {
		e.ServerLog.Fatalln(err)
	}

	db.LogMode(true)
	db.AutoMigrate(models...)

	return db
}
