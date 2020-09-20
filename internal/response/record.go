package response

import (
	"omega/domain/base/basrepo"
	"omega/domain/service"
	"omega/internal/types"
)

// RecordCreate make it simpler for calling the record
func (r *Response) RecordCreate(ev types.Event, newData interface{}) {
	r.Record(ev, nil, newData)
}

// Record is used for saving activity
func (r *Response) Record(ev types.Event, data ...interface{}) {
	activityServ := service.ProvideBasActivityService(basrepo.ProvideActivityRepo(r.Engine))
	activityServ.Record(r.Context, ev, data...)
}
