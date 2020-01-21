package response

import (
	"net/http"
	"omega/engine"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Result is a standard output for success and faild requests
type Result struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
	Code    uint        `json:"code,omitempty"`
}

// InvalidID returns in case of the id in uri not number
func InvalidID(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, &Result{
		Message: "Invalid ID",
		Code:    1400,
		Error:   err.Error(),
	})
}

// RecordNotFound has two different state, the record not exist or internal error
func RecordNotFound(c *gin.Context, err error, part string) {
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, &Result{
			Message: part + " not exist",
			Code:    1404,
			Error:   err.Error(),
		})

	} else {
		c.JSON(http.StatusInternalServerError, &Result{
			Message: "Failed to fetch " + part,
			Code:    1500,
			Error:   err.Error(),
		})
	}
}

// Success used for returning a record of anything, used inside FindByID, FindByUsername
// and etc
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &Result{
		Data: data,
	})
}

// SuccessSave attach a message to the response
func SuccessSave(c *gin.Context, data interface{}, msg string) {
	c.JSON(http.StatusOK, &Result{
		Data:    data,
		Message: msg,
	})
}

// SuccessMessage attach a message to the response
func SuccessMessage(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, &Result{
		Message: msg,
	})
}

// ErrorInBinding is used for any kind of error during binding
func ErrorInBinding(c *gin.Context, err error, part string) {
	c.JSON(http.StatusBadRequest, &Result{
		Message: "Error in binding " + part,
		Code:    1400,
		Error:   err.Error(),
	})
}

// ErrorOnSave check if duplication happened return the proper message related to
// error if not show the error
func ErrorOnSave(c *gin.Context, err error, part string) {
	errMessage := err.Error()
	if strings.Contains(strings.ToUpper(errMessage), "DUPLICATE") {
		c.JSON(http.StatusConflict, &Result{
			Message: "Duplication happened for " + part,
			Code:    1409,
			Error:   errMessage,
		})
	} else {
		c.JSON(http.StatusInternalServerError, &Result{
			Message: "Error in creating " + part,
			Code:    1500,
			Error:   errMessage,
		})
	}
}

// NoPermission is simpoe func for showing this error
func NoPermission(c *gin.Context) {
	c.JSON(http.StatusForbidden, &Result{
		Message: "You don't have permission",
	})
}

// NoPermissionRecord is simpoe func for showing this error
func NoPermissionRecord(c *gin.Context, e engine.Engine, msg string, data ...interface{}) {
	e.Record(c, msg, data...)
	c.JSON(http.StatusForbidden, &Result{
		Message: "You don't have permission",
	})
}
