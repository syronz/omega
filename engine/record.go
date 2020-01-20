package engine

import (
	"encoding/json"
	"fmt"
	"omega/internal/models"

	"github.com/gin-gonic/gin"
)

// RecordType is and int used as an enum
type RecordType int

const (
	read RecordType = iota
	writeBefore
	writeAfter
	writeBoth
)

// Record an activity
func (e *Engine) Record(c *gin.Context, event string, data ...interface{}) {
	var userID uint64
	var username string

	recordType := e.findRecordType(data...)
	before, after := e.fillBeforeAfter(recordType, data...)

	if e.isRecordSetInEnvironment(recordType) {
		return
	}
	if userIDtmp, ok := c.Get("USER_ID"); ok {
		userID = userIDtmp.(uint64)
	}
	if usernameTmp, ok := c.Get("USERNAME"); ok {
		username = usernameTmp.(string)
	}

	activity := models.Activity{
		Event:    event,
		UserID:   userID,
		Username: username,
		IP:       c.ClientIP(),
		Path:     c.Request.URL.Path,
		Before:   string(before),
		After:    string(after),
	}

	err := e.ActivityDB.Save(&activity).Error
	e.CheckError(err, fmt.Sprintf("Failed in saving activity for %+v", activity))
}

func (e *Engine) fillBeforeAfter(recordType RecordType, data ...interface{}) (before, after []byte) {
	var err error
	switch {
	case recordType == writeBefore || recordType == writeBoth:
		before, err = json.Marshal(data[0])
		e.CheckError(err, "error in encoding data to before-json")
		fallthrough
	case recordType == writeAfter || recordType == writeBoth:
		after, err = json.Marshal(data[1])
		e.CheckError(err, "error in encoding data to after-json")
	}

	return
}

func (e *Engine) findRecordType(data ...interface{}) RecordType {
	switch len(data) {
	case 0:
		return read
	case 2:
		return writeBoth
	default:
		if data[0] == nil {
			return writeAfter
		}
	}

	return writeBefore
}

func (e *Engine) isRecordSetInEnvironment(recordType RecordType) bool {
	switch recordType {
	case read:
		if !e.Environments.Setting.RecordRead {
			return true
		}
	default:
		if !e.Environments.Setting.RecordWrite {
			return true
		}
	}
	return false
}

func recordActivity(num int, c *gin.Context, recordType RecordType) {

}
