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
	"omega/internal/param"
	"omega/internal/response"
	"omega/internal/types"
	"omega/pkg/dict"
	"omega/pkg/excel"
	"omega/pkg/glog"

	"github.com/gin-gonic/gin"
)

// UserAPI for injecting user service
type UserAPI struct {
	Service service.BasUserServ
	Engine  *core.Engine
}

// ProvideUserAPI for user is used in wire
func ProvideUserAPI(c service.BasUserServ) UserAPI {
	return UserAPI{Service: c, Engine: c.Engine}
}

// FindByID is used for fetch a user by it's id
func (p *UserAPI) FindByID(c *gin.Context) {
	resp := response.New(p.Engine, c, base.Domain)
	var err error
	var user basmodel.User
	var fix types.FixedCol

	if fix, err = resp.GetFixedCol(c.Param("userID"), "E1090173", basterm.User); err != nil {
		return
	}

	if !resp.CheckRange(fix.CompanyID, fix.NodeID) {
		return
	}

	if user, err = p.Service.FindByID(fix); err != nil {
		resp.Error(err).JSON()
		return
	}

	user.Password = ""

	resp.Record(base.ViewUser)
	resp.Status(http.StatusOK).
		MessageT(corterm.VInfo, basterm.User).
		JSON(user)
}

// FindByUsername is used when we try to find a user with username
func (p *UserAPI) FindByUsername(c *gin.Context) {
	resp := response.New(p.Engine, c, base.Domain)
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
func (p *UserAPI) List(c *gin.Context) {
	resp, params := response.NewParam(p.Engine, c, basmodel.UserTable, base.Domain)

	if username := c.Query("username"); username != "" {
		params.Filter = fmt.Sprintf("username[eq]'%v'", username)
	}

	data := make(map[string]interface{})
	var err error

	if params.CompanyID, err = resp.GetCompanyID("E1019954"); err != nil {
		return
	}

	if !resp.CheckRange(params.CompanyID, 0) {
		return
	}

	if data["list"], data["count"], err = p.Service.List(params); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(base.ListUser)
	resp.Status(http.StatusOK).
		MessageT(corterm.ListOfV, basterm.Users).
		JSON(data)
}

// Create user
func (p *UserAPI) Create(c *gin.Context) {
	resp := response.New(p.Engine, c, base.Domain)
	var user, createdUser basmodel.User
	var err error

	if user.CompanyID, user.NodeID, err = resp.GetCompanyNode("E1061527", base.Domain); err != nil {
		resp.Error(err).JSON()
		return
	}

	if user.CompanyID, err = resp.GetCompanyID("E1031989"); err != nil {
		return
	}

	if !resp.CheckRange(user.CompanyID, user.NodeID) {
		return
	}

	if err = resp.Bind(&user, "E1082301", base.Domain, basterm.User); err != nil {
		return
	}

	if createdUser, err = p.Service.Create(user); err != nil {
		resp.Error(err).JSON()
		return
	}

	user.Password = ""

	resp.RecordCreate(base.CreateUser, user)
	resp.Status(http.StatusOK).
		MessageT(corterm.VCreatedSuccessfully, dict.R(basterm.User)).
		JSON(createdUser)
}

// Update user
func (p *UserAPI) Update(c *gin.Context) {
	resp := response.New(p.Engine, c, base.Domain)
	var err error

	var user, userBefore, userUpdated basmodel.User
	var fix types.FixedCol

	if fix, err = resp.GetFixedCol(c.Param("userID"), "E1097541", basterm.User); err != nil {
		return
	}

	if !resp.CheckRange(fix.CompanyID, fix.NodeID) {
		return
	}

	if err = resp.Bind(&user, "E1065844", base.Domain, basterm.User); err != nil {
		return
	}

	if userBefore, err = p.Service.FindByID(fix); err != nil {
		resp.Error(err).JSON()
		return
	}

	// if user.ID, err = types.StrToRowID(c.Param("userID")); err != nil {
	// 	resp.Error(corerr.InvalidID).JSON()
	// 	return
	// }

	// if err = c.ShouldBindJSON(&user); err != nil {
	// 	c.AbortWithStatusJSON(http.StatusNotAcceptable, err)
	// 	return
	// }

	// if userBefore, err = p.Service.FindByID(user.ID); err != nil {
	// 	resp.Error(err).JSON()
	// 	return
	// }

	user.ID = fix.ID
	user.CompanyID = fix.CompanyID
	user.NodeID = fix.NodeID
	if userUpdated, err = p.Service.Save(user); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(base.UpdateUser, userBefore, userUpdated)

	resp.Status(http.StatusOK).
		MessageT(corterm.VUpdatedSuccessfully, basterm.User).
		JSON(userUpdated)

}

// Delete user
func (p *UserAPI) Delete(c *gin.Context) {
	resp := response.New(p.Engine, c, base.Domain)
	var err error
	var user basmodel.User
	var fix types.FixedCol

	// if user.ID, err = types.StrToRowID(c.Param("userID")); err != nil {
	// 	resp.Error(corerr.InvalidID).JSON()
	// 	return
	// }

	if fix, err = resp.GetFixedCol(c.Param("userID"), "E1046157", basterm.User); err != nil {
		return
	}

	if user, err = p.Service.Delete(fix); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(base.DeleteUser, user)
	resp.Status(http.StatusOK).
		MessageT(corterm.VDeletedSuccessfully, basterm.User).
		JSON()
}

// Excel generate excel files based on search
func (p *UserAPI) Excel(c *gin.Context) {
	resp := response.New(p.Engine, c, base.Domain)

	params := param.Get(c, p.Engine, basterm.Users)

	users, err := p.Service.Excel(params)
	if err != nil {
		resp.Error(err).JSON()
		return
	}

	ex := excel.New("node").
		AddSheet("Nodes").
		AddSheet("Summary").
		Active("Nodes").
		SetPageLayout("landscape", "A4").
		SetPageMargins(0.2).
		SetHeaderFooter().
		SetColWidth("A", "C", 17).
		SetColWidth("D", "F", 15.3).
		SetColWidth("H", "H", 20).
		SetColWidth("N", "O", 20).
		Active("Summary").
		Active("Nodes").
		WriteHeader("ID", "Company ID", "Node ID", "Username", "Role", "Lang", "Email")

	for i, v := range users {
		column := &[]interface{}{
			v.ID,
			v.CompanyID,
			v.NodeID,
			v.Username,
			v.Role,
			v.Lang,
			v.Email,
		}
		err = ex.File.SetSheetRow(ex.ActiveSheet, fmt.Sprint("A", i+2), column)
		glog.CheckError(err, "Error in writing to the excel in user")
	}

	ex.Sheets[ex.ActiveSheet].Row = len(users) + 1

	ex.AddTable()

	buffer, downloadName, err := ex.Generate()
	if err != nil {
		// resp.Error(err).JSON()
		c.JSON(http.StatusInternalServerError, &response.Result{
			Message: "Error in generating Excel file",
		})
		return
	}

	resp.Record(base.ExcelUser)

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename="+downloadName)
	c.Data(http.StatusOK, "application/octet-stream", buffer.Bytes())

}
