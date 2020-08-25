package response

import (
	"omega/domain/base/basrepo"
	"omega/domain/service"
	"omega/internal/types"
)

// Record is used for saving activity
func (r *Response) Record(ev types.Event, data ...interface{}) {
	activityServ := service.ProvideBasActivityService(basrepo.ProvideActivityRepo(r.Engine))
	activityServ.Record(r.Context, ev, data...)
}
