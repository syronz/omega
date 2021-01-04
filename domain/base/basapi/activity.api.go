package basapi

import (
	"fmt"
	"net/http"
	"omega/domain/base"
	"omega/domain/base/basmodel"
	"omega/domain/base/message/basterm"
	"omega/domain/service"
	"omega/internal/core"
	"omega/internal/core/corterm"
	"omega/internal/response"

	"github.com/gin-gonic/gin"
)

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

// ListAll of all activities among all companies
func (p *ActivityAPI) ListAll(c *gin.Context) {
	resp, params := response.NewParam(p.Engine, c, basmodel.ActivityTable, base.Domain)

	data, err := p.Service.List(params)
	if err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(base.AllActivity)
	resp.Status(http.StatusOK).
		MessageT(corterm.ListOfV, basterm.Activities).
		JSON(data)
}

// ListCompany of all activities among all companies
func (p *ActivityAPI) ListCompany(c *gin.Context) {
	resp, params := response.NewParam(p.Engine, c, basmodel.ActivityTable, base.Domain)
	var err error

	if params.CompanyID, err = resp.GetCompanyID("E1029386"); err != nil {
		return
	}

	if !resp.CheckRange(params.CompanyID) {
		return
	}

	params.PreCondition = fmt.Sprintf("(bas_activities.company_id = %v) ", params.CompanyID)

	data, err := p.Service.List(params)
	if err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(base.AllActivity)
	resp.Status(http.StatusOK).
		MessageT(corterm.ListOfV, basterm.Activities).
		JSON(data)
}

// ListSelf of all activities among all companies
func (p *ActivityAPI) ListSelf(c *gin.Context) {
	resp, params := response.NewParam(p.Engine, c, basmodel.ActivityTable, base.Domain)
	var err error

	// if params.CompanyID, err = resp.GetCompanyID("E1027519"); err != nil {
	// 	return
	// }

	params.PreCondition = fmt.Sprintf("bas_activities.user_id = %v", params.UserID)

	data, err := p.Service.List(params)
	if err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(base.AllActivity)
	resp.Status(http.StatusOK).
		MessageT(corterm.ListOfV, basterm.Activities).
		JSON(data)
}
