package engine

import (
	"encoding/json"
	"fmt"
	"omega/config"
	"omega/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

// Engine to keep all database connections and
// logs configuration and environments and etc
type Engine struct {
	DB           *gorm.DB
	ActivityDB   *gorm.DB
	ServerLog    *logrus.Logger
	ApiLog       *logrus.Logger
	Environments config.Environment
}

// Debug print struct with details with logrus ability
func (e *Engine) Debug(objs ...interface{}) {
	for _, v := range objs {
		e.ServerLog.Debug(fmt.Sprintf("%T :: %+[1]v", v))
	}
}

// DumpError print struct with details with logrus ability
func (e *Engine) DumpError(objs ...interface{}) {
	for _, v := range objs {
		e.ServerLog.Error(fmt.Sprintf("%T :: %+[1]v", v))
	}
}

// CheckError print all errors which happened inside the services, mainly they just have
// an error and a message
func (e *Engine) CheckError(err error, message string) {
	if err != nil {
		e.ServerLog.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Error(message)
	}
}

// CheckInfo print all errors which happened inside the services, mainly they just have
// an error and a message
func (e *Engine) CheckInfo(err error, message string) {
	if err != nil {
		e.ServerLog.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Info(message)
	}
}

// Record an activity
func (e *Engine) Record(c *gin.Context, event string, data ...interface{}) {
	var before, after []byte
	var err error
	var userID uint64
	var username string

	type RecordType int
	const (
		read RecordType = iota
		write
	)
	var recordType RecordType = read

	num := len(data)
	if num > 0 {
		if data[0] != nil {
			before, err = json.Marshal(data[0])
			e.CheckError(err, "error in encoding data to before-json")
		}
		recordType = write
	}

	switch recordType {
	case read:
		if !e.Environments.Setting.RecordRead {
			return
		}
	default:
		if !e.Environments.Setting.RecordRead {
			return
		}
	}

	if num > 1 {
		after, err = json.Marshal(data[1])
		e.CheckError(err, "error in encoding data to after-json")
	}

	userIDtmp, ok := c.Get("USER_ID")
	if ok {
		userID = userIDtmp.(uint64)
	}

	usernameTmp, ok := c.Get("USERNAME")
	if ok {
		username = usernameTmp.(string)
	}

	activity := models.Activity{
		Event:    event,
		UserID:   userID,
		IP:       c.ClientIP(),
		Path:     c.Request.URL.Path,
		Before:   string(before),
		After:    string(after),
		Username: username,
	}

	err = e.ActivityDB.Save(&activity).Error
	e.CheckInfo(err, fmt.Sprintf("Failed in saving activity for %+v", activity))

	return
}
