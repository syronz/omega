package response

import (
	"net/http"
	"omega/internal/glog"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Result is a standard output for success and faild requests
type Result struct {
	// TODO: by my opinion (DSH) status is't necessery
	Status  bool        `json:"status,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
	Code    uint        `json:"code,omitempty"`
}

// InvalidID returns in case of the id in uri not number
func InvalidID(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, &Result{
		Status:  false,
		Message: "Invalid ID",
		Code:    1400,
		Errors:  err.Error(),
	})
}

// RecordNotFound has two different state, the record not exist or internal error
func RecordNotFound(c *gin.Context, err error, part string) {
	glog.Debug(err)
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, &Result{
			Status:  false,
			Message: part + " Not Exist",
			Code:    1404,
			Errors:  err.Error(),
		})

	} else {
		c.JSON(http.StatusInternalServerError, &Result{
			Status:  false,
			Message: "Failed to Fetch " + part,
			Code:    1500,
			Errors:  err.Error(),
		})
	}
}

// SuccessOne used for returning a record of anything, used inside FindByID, FindByUsername
// and etc
func SuccessOne(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &Result{
		Status: true,
		Data:   data,
	})
}

// SuccessAll return a list of rows
func SuccessAll(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &Result{
		Status: true,
		Data:   data,
	})
}

// ErrorInBinding is used for any kind of error during binding
func ErrorInBinding(c *gin.Context, err error, part string) {
	c.JSON(http.StatusBadRequest, &Result{
		Status:  false,
		Message: "Error in binding " + part,
		Code:    1400,
		Errors:  err.Error(),
	})
}

// ErrorInCreating check if duplication happened return the proper message related to
// error if not show the error
func ErrorInCreating(c *gin.Context, err error, part string) {
	errMessage := err.Error()
	if strings.Contains(strings.ToUpper(errMessage), "DUPLICATE") {
		c.JSON(http.StatusConflict, &Result{
			Status:  false,
			Message: "This " + part + " is exist",
			Code:    1409,
			Errors:  errMessage,
		})
	} else {
		c.JSON(http.StatusInternalServerError, &Result{
			Status:  false,
			Message: "Error in creating " + part,
			Code:    1500,
			Errors:  errMessage,
		})

	}

}
