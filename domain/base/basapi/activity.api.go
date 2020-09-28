package basapi

import (
	"net/http"
	"omega/domain/base"
	"omega/domain/base/basmodel"
	"omega/domain/service"
	"omega/internal/core"
	"omega/internal/core/corterm"
	"omega/internal/param"
	"omega/internal/response"

	"github.com/gin-gonic/gin"
)

const thisActivity = "activity"
const thisActivities = "bas_activities"

// ActivityAPI for injecting activity service
type ActivityAPI struct {
	Service service.BasActivityServ
	Engine  *core.Engine
}

// ProvideActivityAPI for activity is used in wire
func ProvideActivityAPI(c service.BasActivityServ) ActivityAPI {
	return ActivityAPI{Service: c, Engine: c.Engine}
}

// Create activity
func (p *ActivityAPI) Create(c *gin.Context) {
	var activity basmodel.Activity
	resp := response.New(p.Engine, c, base.Domain)

	if err := c.ShouldBindJSON(&activity); err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, err)
		return
	}

	createdActivity, err := p.Service.Save(activity)
	if err != nil {
		c.JSON(403, gin.H{"error": err.Error()})
		return
	}

	resp.Status(203).
		Message("activity created successfully").
		JSON(createdActivity)
}

// List of activities
func (p *ActivityAPI) List(c *gin.Context) {
	resp := response.New(p.Engine, c, base.Domain)

	params := param.Get(c, p.Engine, thisActivities)

	data, err := p.Service.List(params)
	if err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(base.AllActivity)
	resp.Status(http.StatusOK).
		MessageT(corterm.ListOfV, thisActivities).
		JSON(data)
}
