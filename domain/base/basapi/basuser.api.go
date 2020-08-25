package basapi

import (
	"fmt"
	"net/http"
	"omega/domain/base/basevent"
	"omega/domain/base/basmodel"
	"omega/domain/service"
	"omega/internal/core"
	"omega/internal/param"
	"omega/internal/response"
	"omega/internal/term"
	"omega/internal/types"
	"omega/utils/excel"

	"github.com/gin-gonic/gin"
)

const thisBasUser = "user"
const thisBasUsers = "users"

// BasUserAPI for injecting user service
type BasUserAPI struct {
	Service service.BasUserServ
	Engine  *core.Engine
}

// ProvideBasUserAPI for user is used in wire
func ProvideBasUserAPI(c service.BasUserServ) BasUserAPI {
	return BasUserAPI{Service: c, Engine: c.Engine}
}

// FindByID is used for fetch a user by it's id
func (p *BasUserAPI) FindByID(c *gin.Context) {
	resp := response.New(p.Engine, c)
	var err error
	var user basmodel.BasUser

	if user.ID, err = types.StrToRowID(c.Param("userID")); err != nil {
		resp.Error(term.Invalid_ID).JSON()
		return
	}

	if user, err = p.Service.FindByID(user.ID); err != nil {
		resp.Status(http.StatusNotFound).Error(err).MessageT(term.Record_Not_Found).JSON()
		return
	}

	user.Password = ""

	resp.Record(basevent.BasUserView)
	resp.Status(http.StatusOK).
		MessageT(term.V_info, thisBasUser).
		JSON(user)
}

// FindByUsername is used when we try to find a user with username
func (p *BasUserAPI) FindByUsername(c *gin.Context) {
	resp := response.New(p.Engine, c)
	username := c.Param("username")

	user, err := p.Service.FindByUsername(username)
	if err != nil {
		resp.Status(http.StatusNotFound).Error(err).JSON()
		return
	}

	user.Password = ""

	resp.Status(http.StatusOK).JSON(user)
}

// List of users
func (p *BasUserAPI) List(c *gin.Context) {
	resp := response.New(p.Engine, c)

	params := param.Get(c, p.Engine, thisBasUsers)

	data, err := p.Service.List(params)
	if err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(basevent.BasUserList)
	resp.Status(http.StatusOK).
		MessageT(term.List_of_V, thisBasUsers).
		JSON(data)
}

// Create user
func (p *BasUserAPI) Create(c *gin.Context) {

	var user basmodel.BasUser
	resp := response.New(p.Engine, c)

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, err)
		return
	}

	params := param.Get(c, p.Engine, thisBasUsers)
	createdBasUser, err := p.Service.Create(user, params)
	if err != nil {
		resp.Error(err).JSON()
		return
	}

	user.Password = ""
	resp.Record(basevent.BasUserCreate, nil, user)

	resp.Status(http.StatusOK).
		Message(term.User_created_successfully).
		JSON(createdBasUser)
}

// Update user
func (p *BasUserAPI) Update(c *gin.Context) {
	resp := response.New(p.Engine, c)
	var err error

	var user, userBefore, userUpdated basmodel.BasUser

	if user.ID, err = types.StrToRowID(c.Param("userID")); err != nil {
		resp.Error(term.Invalid_ID).JSON()
		return
	}

	if err = c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, err)
		return
	}

	if userBefore, err = p.Service.FindByID(user.ID); err != nil {
		resp.Status(http.StatusNotFound).Error(term.Record_Not_Found).JSON()
		return
	}

	if userUpdated, err = p.Service.Save(user); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(basevent.BasUserUpdate, userBefore, userUpdated)

	resp.Status(http.StatusOK).
		MessageT(term.V_updated_successfully, thisBasUser).
		JSON(userUpdated)

}

// Delete user
func (p *BasUserAPI) Delete(c *gin.Context) {
	resp := response.New(p.Engine, c)
	var err error
	var user basmodel.BasUser

	if user.ID, err = types.StrToRowID(c.Param("userID")); err != nil {
		p.Engine.CheckError(err, err.Error())
		resp.Error(term.Invalid_ID).JSON()
		return
	}

	params := param.Get(c, p.Engine, thisBasUser)

	if user, err = p.Service.Delete(user.ID, params); err != nil {
		resp.Status(http.StatusInternalServerError).Error(err).JSON()
		return
	}

	resp.Record(basevent.BasUserDelete, user)
	resp.Status(http.StatusOK).
		MessageT(term.V_deleted_successfully, thisBasUser).
		JSON()
}

// Excel generate excel files based on search
func (p *BasUserAPI) Excel(c *gin.Context) {
	resp := response.New(p.Engine, c)

	params := param.Get(c, p.Engine, thisBasUsers)

	users, err := p.Service.Excel(params)
	if err != nil {
		resp.Status(http.StatusNotFound).Error(term.Record_Not_Found).JSON()
		return
	}

	ex := excel.New("node").
		AddSheet("Nodes").
		AddSheet("Summary").
		Active("Nodes").
		SetPageLayout("landscape", "A4").
		SetPageMargins(0.2).
		SetHeaderFooter().
		SetColWidth("A", "A", 20).
		SetColWidth("B", "C", 15.3).
		SetColWidth("F", "F", 20).
		SetColWidth("L", "M", 20).
		Active("Summary").
		Active("Nodes").
		WriteHeader("ID", "BasUsername", "Role", "Language", "Email")

	for i, v := range users {
		extra := v.Extra.(map[string]interface{})
		column := &[]interface{}{
			v.ID,
			v.Username,
			extra["role"],
			v.Language,
			v.Email,
		}
		err = ex.File.SetSheetRow(ex.ActiveSheet, fmt.Sprint("A", i+2), column)
		p.Engine.CheckError(err, "Error in writing to the excel in user")
	}

	ex.Sheets[ex.ActiveSheet].Row = len(users) + 1

	ex.AddTable()

	buffer, downloadName, err := ex.Generate()
	if err != nil {
		c.JSON(http.StatusInternalServerError, &response.Result{
			Message: "Error in generating Excel file",
		})
		return
	}

	resp.Record(basevent.BasUserExcel)

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename="+downloadName)
	c.Data(http.StatusOK, "application/octet-stream", buffer.Bytes())

}
