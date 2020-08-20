package basapi

import (
	"net/http"
	"omega/domain/base/basevent"
	"omega/domain/base/basmodel"
	"omega/domain/base/basresource"
	"omega/domain/service"
	"omega/internal/core"
	"omega/internal/param"
	"omega/internal/response"
	"omega/internal/term"

	"github.com/gin-gonic/gin"
)

const thisBasActivity = "activity"
const thisActivities = "bas_activities"

// BasActivityAPI for injecting activity service
type BasActivityAPI struct {
	Service service.BasActivityServ
	Engine  *core.Engine
}

// ProvideBasActivityAPI for activity is used in wire
func ProvideBasActivityAPI(c service.BasActivityServ) BasActivityAPI {
	return BasActivityAPI{Service: c, Engine: c.Engine}
}

// Create activity
func (p *BasActivityAPI) Create(c *gin.Context) {
	var activity basmodel.BasActivity
	resp := response.New(p.Engine, c)

	if err := c.ShouldBindJSON(&activity); err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, err)
		return
	}

	createdBasActivity, err := p.Service.Save(activity)
	if err != nil {
		c.JSON(403, gin.H{"error": err.Error()})
		return
	}

	resp.Status(203).
		Message("activity created successfully").
		JSON(createdBasActivity)
}

// List of activities
func (p *BasActivityAPI) List(c *gin.Context) {
	resp := response.New(p.Engine, c)

	if resp.CheckAccess(basresource.BasActivityAll) {
		resp.Status(http.StatusForbidden).Error(term.You_dont_have_permission).JSON()
		return
	}

	params := param.Get(c, p.Engine, thisActivities)

	data, err := p.Service.List(params)
	if err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(basevent.BasActivityAll)
	resp.Status(http.StatusOK).
		MessageT(term.List_of_V, thisActivities).
		JSON(data)
}
