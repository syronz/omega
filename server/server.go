package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Setup(db *gorm.DB) *gin.Engine {

	r := gin.Default()

	router(r, db)
	return r
}
