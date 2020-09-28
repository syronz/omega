package service

import (
	"encoding/json"
	"fmt"
	"omega/domain/base"
	"omega/domain/base/basmodel"
	"omega/domain/base/basrepo"
	"omega/internal/core"
	"omega/internal/param"
	"omega/internal/types"
	"omega/pkg/glog"

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

// BasActivityServ for injecting auth basrepo
type BasActivityServ struct {
	Repo   basrepo.ActivityRepo
	Engine *core.Engine
}

// ProvideBasActivityService for activity is used in wire
func ProvideBasActivityService(p basrepo.ActivityRepo) BasActivityServ {
	return BasActivityServ{Repo: p, Engine: p.Engine}
}

// Save activity
func (p *BasActivityServ) Save(activity basmodel.Activity) (createdActivity basmodel.Activity, err error) {
	createdActivity, err = p.Repo.Create(activity)

	// p.Engine.CheckInfo(err, fmt.Sprintf("Failed in saving activity for %+v", activity))

	return
}

// Record will save the activity
func (p *BasActivityServ) Record(c *gin.Context, ev types.Event, data ...interface{}) {
	var userID types.RowID
	var username string

	recordType := p.findRecordType(data...)
	before, after := p.fillBeforeAfter(recordType, data...)

	if len(data) > 0 && !p.Engine.Envs.ToBool(base.RecordWrite) {
		return
	}

	if len(data) == 0 && !p.Engine.Envs.ToBool(base.RecordRead) {
		return
	}

	if p.isRecordSetInEnvironment(recordType) {
		return
	}
	if userIDtmp, ok := c.Get("USER_ID"); ok {
		userID = userIDtmp.(types.RowID)
	}
	if usernameTmp, ok := c.Get("USERNAME"); ok {
		username = usernameTmp.(string)
	}

	activity := basmodel.Activity{
		Event:    ev.String(),
		UserID:   userID,
		Username: username,
		IP:       c.ClientIP(),
		URI:      c.Request.RequestURI,
		Before:   string(before),
		After:    string(after),
	}

	_, err := p.Repo.Create(activity)
	glog.CheckError(err, fmt.Sprintf("Failed in saving activity for %+v", activity))
}

func (p *BasActivityServ) fillBeforeAfter(recordType RecordType, data ...interface{}) (before, after []byte) {
	var err error
	if recordType == writeBefore || recordType == writeBoth {
		before, err = json.Marshal(data[0])
		glog.CheckError(err, "error in encoding data to before-json")
	}
	if recordType == writeAfter || recordType == writeBoth {
		after, err = json.Marshal(data[1])
		glog.CheckError(err, "error in encoding data to after-json")
	}

	return
}

func (p *BasActivityServ) findRecordType(data ...interface{}) RecordType {
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

func (p *BasActivityServ) isRecordSetInEnvironment(recordType RecordType) bool {
	switch recordType {
	case read:
		if !p.Engine.Envs.ToBool(base.RecordRead) {
			return true
		}
	default:
		if !p.Engine.Envs.ToBool(base.RecordWrite) {
			return true
		}
	}
	return false
}

// List of activities, it support pagination and search and return back count
func (p *BasActivityServ) List(params param.Param) (data map[string]interface{}, err error) {

	data = make(map[string]interface{})

	data["list"], err = p.Repo.List(params)
	glog.CheckError(err, "activities list")
	if err != nil {
		return
	}

	data["count"], err = p.Repo.Count(params)
	glog.CheckError(err, "activities count")

	return
}
