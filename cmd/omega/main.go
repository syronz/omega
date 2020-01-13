package main

import (
	// "os"

	// "github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	// "rest-gin-gorm/internal"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"rest-gin-gorm/internal/one"
	"rest-gin-gorm/pkg/invoice"
	"rest-gin-gorm/pkg/product"
	"rest-gin-gorm/pkg/user"
	"rest-gin-gorm/server"
)

func initDB() *gorm.DB {

	// internal.InternalPing()
	one.OnePing()

	db, err := gorm.Open("mysql", "root:Qaz1@345@tcp(127.0.0.1:3306)/alpha?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&product.Product{})
	db.AutoMigrate(&invoice.Invoice{})
	db.AutoMigrate(&user.User{})

	return db
}

func main() {
	// Log as JSON instead of the default ASCII formatter.
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logrus.SetOutput(os.Stdout)

	// Only logrus the warning severity or above.
	// logrus.SetLevel(logrus.WarnLevel)
	logrus.SetLevel(logrus.InfoLevel)
	logrus.Info("THIS IS LOGRUS TEST .................................")
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})

	fmt.Printf("\n################ %T \n", log)
	log.Info("THIS IS LOGRUS TEST .................................")

	db := initDB()
	defer db.Close()

	s := server.Setup(db, log)

	err := s.Run()
	if err != nil {
		panic(err)
	}
}
