package initiate

import (
	"omega/internal/core"
	"omega/pkg/user"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func initDB(e core.Engine, dbType string, url string) *gorm.DB {

	db, err := gorm.Open(dbType, url)
	if err != nil {
		e.Log.Fatalln(err)
	}

	db.LogMode(false)

	if gin.IsDebugging() {
		db.LogMode(true)
	}

	db.AutoMigrate(&user.User{})

	return db
}
