package user

import (
	"net/http"
	"omega/internal/param"
	"omega/internal/response"
	"omega/resources"
	"omega/shared"
	"strconv"

	"github.com/gin-gonic/gin"
)

// BuildFindAll users
func (p *API) BuildFindAll(c *gin.Context) {
	if p.Engine.CheckAccess(c, "users:names") {
		response.NoPermission(c)
		return
	}
	users, err := p.Service.FindAll()

	if err != nil {
		response.RecordNotFound(c, err, "users")
		return
	}

	response.Success(c, users)
}

// BuildList of users
func (p *API) BuildList(c *gin.Context) {
	if p.Engine.CheckAccess(c, "users:read") {
		response.NoPermissionRecord(c, p.Engine, "user-list-forbidden")
		return
	}

	params := param.Get(c)

	p.Engine.Debug(params)
	data, err := p.Service.List(params)
	if err != nil {
		response.RecordNotFound(c, err, "users")
		return
	}

	p.Engine.Record(c, "user-list")
	response.Success(c, data)
}

// BuildFindByID is used for fetch a user by his id
func (p *API) BuildFindByID(c *gin.Context) {
	if p.Engine.CheckAccess(c, "users:read") {
		response.NoPermissionRecord(c, p.Engine, "user-view-forbidden")
		return
	}
	id, err := strconv.ParseUint(c.Param("id"), 10, 16)
	if err != nil {
		response.InvalidID(c, err)
		return
	}

	user, err := p.Service.FindByID(id)

	if err != nil {
		response.RecordNotFound(c, err, "user")
		return
	}

	p.Engine.Record(c, "user-view")
	response.Success(c, user)
}

// BuildCreate user
func (p *API) BuildCreate(c *gin.Context) {

	builder := shared.New(c, p.Engine, p.Service, "user-create").
		Bind(&User{}).
		CheckAccess(resources.UserWrite).
		Save(p.Service.BuildSave).
		Record()
	p.Engine.Debug("++++++++++++++++++", builder.NewModel)

	createdUser := (builder.SavedModel.(User))

	if builder.Error() != nil {
		response.ErrorInBinding(c, builder.Error(), "KKKKKKKKKKKKKKKK")
		return
	}

	response.SuccessSave(c, createdUser, "user/create")
}

// BuildUpdate user
func (p *API) BuildUpdate(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 16)
	if err != nil {
		response.InvalidID(c, err)
		return
	}

	var user User

	if err = c.BindJSON(&user); err != nil {
		response.ErrorInBinding(c, err, "update user")
		return
	}
	user.ID = id
	if p.Engine.CheckAccess(c, "users:write") {
		response.NoPermissionRecord(c, p.Engine, "user-update-forbidden", nil, user)
		return
	}

	userBefore, err := p.Service.FindByID(id)
	if err != nil {
		response.RecordNotFound(c, err, "update user")
		return
	}

	updatedUser, err := p.Service.Save(user)
	if err != nil {
		response.ErrorOnSave(c, err, "update user")
		return
	}

	userBefore.Password = ""
	updatedUser.Password = ""

	p.Engine.Record(c, "user-update", userBefore, updatedUser)
	response.SuccessSave(c, updatedUser, "user updated")
}

// BuildDelete user
func (p *API) BuildDelete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 16)
	if err != nil {
		response.InvalidID(c, err)
		return
	}
	if p.Engine.CheckAccess(c, "users:write") {
		response.NoPermissionRecord(c, p.Engine, "user-delete-forbidden", nil, id)
		return
	}

	var user User

	user, err = p.Service.FindByID(id)
	if err != nil {
		response.RecordNotFound(c, err, "delete user")
		return
	}

	err = p.Service.Delete(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &response.Result{
			Message: "Something went wrong, cannot delete this user",
			Code:    1500,
		})
		return
	}

	user.Password = ""
	p.Engine.Record(c, "user-delete", user)
	response.SuccessSave(c, user, "user successfully deleted")
}
