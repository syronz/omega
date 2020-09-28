package param

import (
	"omega/internal/core"
	"omega/internal/types"
	"omega/pkg/glog"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// Get is a function for filling param.Model
func Get(c *gin.Context, engine *core.Engine, part string) (param Param) {
	var err error

	generateOrder(c, &param, part)
	generateSelectedColumns(c, &param)
	generateLimit(c, &param)
	generateOffset(c, &param)

	// param.Search = strings.TrimSpace(c.Query("search"))
	param.Filter = strings.TrimSpace(c.Query("filter"))

	userID, ok := c.Get("USER_ID")
	if ok {
		glog.CheckInfo(err, "User ID is not exist")
		param.UserID = userID.(types.RowID)
	}

	param.Lang = core.GetLang(c, engine)

	param.ErrPanel = engine.Envs[core.ErrPanel]

	return param
}

func generateOrder(c *gin.Context, param *Param, part string) {
	orderBy := part + ".id"
	direction := "desc"

	if c.Query("order_by") != "" {
		orderBy = c.Query("order_by")
	}

	if c.Query("direction") != "" {
		direction = c.Query("direction")
	}

	param.Order = orderBy + " " + direction
}

func generateSelectedColumns(c *gin.Context, param *Param) {
	param.Select = "*"
	if c.Query("select") != "" {
		param.Select = c.Query("select")
	}
}

func generateLimit(c *gin.Context, param *Param) {
	var err error
	param.Limit = 10
	if c.Query("page_size") != "" {
		param.Limit, err = strconv.ParseUint(c.Query("page_size"), 10, 16)
		if err != nil {
			// TODO: get path from gin.Context
			glog.CheckError(err, "Limit is not a number")
			param.Limit = 10
		}
	}
}

func generateOffset(c *gin.Context, param *Param) {
	var page uint64
	page = 0
	var err error
	if c.Query("page") != "" {
		page, err = strconv.ParseUint(c.Query("page"), 10, 16)
		if err != nil {
			// TODO: get path from gin.Context
			glog.CheckError(err, "Offset is not a positive number")
			page = 0
		}
	}

	param.Offset = param.Limit * (page)
}
