package main

import (
	// "os"

	// "github.com/gin-gonic/gin"
	// "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// "rest-gin-gorm/internal"
	// "fmt"
	// "github.com/sirupsen/logrus"
	// "os"
	"rest-gin-gorm/initiate"
	// "rest-gin-gorm/internal/one"
	// "rest-gin-gorm/pkg/invoice"
	// "rest-gin-gorm/pkg/product"
	// "rest-gin-gorm/pkg/user"
	"rest-gin-gorm/server"
)

func main() {
	// Log as JSON instead of the default ASCII formatter.
	// logrus.SetFormatter(&logrus.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	// logrus.SetOutput(os.Stdout)

	// Only logrus the warning severity or above.
	// logrus.SetLevel(logrus.WarnLevel)
	// logrus.SetLevel(logrus.InfoLevel)

	// fmt.Printf("\n################ %T \n", log)

	cfg := initiate.Setup()
	defer cfg.DB.Close()

	s := server.Setup(cfg.DB, cfg.Log)

	err := s.Run()
	if err != nil {
		panic(err)
	}
}
