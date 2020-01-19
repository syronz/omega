package core

import (
	"fmt"
	"omega/engine"
	"time"

	"github.com/jinzhu/gorm"
)

func initDB(e engine.Engine, dbType string, dsn string, models ...interface{}) *gorm.DB {
	db, err := gorm.Open(dbType, dsn)
	if err != nil {
		e.ServerLog.Fatalln(err)
	}

	fmt.Printf("\n ---------->>>>>>>>>>>>>>>>>>>>>>>>> %+v  \n", dbType)
	db.LogMode(true)
	db.AutoMigrate(models...)
	// for _, v := range models {
	time.Sleep(2 * time.Second)
	fmt.Printf("\n ---------->>>>>>>>>>>>>>>>>>>>>>>>> %+v || %v \n", dbType, dsn)
	// 	db.AutoMigrate(v)
	// }

	// db.AutoMigrate(&user.User{})
	// for _, v := range models {
	// 	db.AutoMigrate(v)
	// }

	return db
}
