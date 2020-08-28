package response

import (
	"net/http"
	"omega/internal/core"
	"omega/internal/term"
	"omega/pkg/dict"
	"strings"

	"github.com/gin-gonic/gin"
)

// Result is a standard output for success and faild requests
type Result struct {
	Message string                 `json:"message,omitempty"`
	Data    interface{}            `json:"data,omitempty"`
	Error   interface{}            `json:"error,omitempty"`
	Extra   map[string]interface{} `json:"extra,omitempty"`
}

// Response holding method related to response
type Response struct {
	Result  Result
	status  int
	Engine  *core.Engine
	Context *gin.Context
	abort   bool
}

// New initiate the Response object
func New(engine *core.Engine, context *gin.Context) *Response {
	return &Response{
		Engine:  engine,
		Context: context,
	}
}

// Error is used for add error to the result
func (r *Response) Error(error interface{}, data ...interface{}) *Response {
	r.Result.Error = error
	r.Result.Data = data
	return r
}

// Status is used for add error to the result
func (r *Response) Status(status int) *Response {
	r.status = status
	return r
}

// Message is used for add error to the result
func (r *Response) Message(message string) *Response {
	r.Result.Message = message
	return r
}

// MessageT get a message and params then translate it
func (r *Response) MessageT(message string, params ...interface{}) *Response {
	r.Result.Message = dict.T(message,
		core.GetLang(r.Context, r.Engine), params...)
	return r
}

// Abort prepare response to abort instead of return in last step (JSON)
func (r *Response) Abort() *Response {
	r.abort = true
	return r
}

// JSON write ouptut as json
func (r *Response) JSON(data ...interface{}) {
	var errText interface{}

	switch r.Result.Error.(type) {
	case nil:
		errText = nil

	case *core.FieldError:
		errorCast := r.Result.Error.(*core.FieldError)
		errorCast.Translate(r.Engine, core.GetLang(r.Context, r.Engine))
		errText = r.Result.Error
		r.status = http.StatusNotAcceptable

	case *core.ErrorWithStatus:
		errorCast := r.Result.Error.(*core.ErrorWithStatus)
		errText = errorCast.Error()
		if errorCast.Message == "" {
			r.Result.Message = dict.T(errText.(string),
				core.GetLang(r.Context, r.Engine))
		} else {
			r.Result.Message = dict.T(errorCast.Message,
				core.GetLang(r.Context, r.Engine))
		}
		r.status = errorCast.Status

	case error:
		if r.status == 0 {
			r.status = http.StatusInternalServerError
		}
		errorCast := r.Result.Error.(error)
		errText = errorCast.Error()
		r.Result.Message, _ = dict.SafeTranslate(errText.(string),
			core.GetLang(r.Context, r.Engine))
		if strings.Contains(strings.ToUpper(errText.(string)), "DUPLICATE") {
			r.status = http.StatusConflict
			r.Result.Message = dict.T(term.Duplication_happened,
				core.GetLang(r.Context, r.Engine))
		}

		if strings.Contains(strings.ToUpper(errText.(string)), "NOT FOUND") {
			r.status = http.StatusNotFound
		}

		if r.status == http.StatusInternalServerError && r.Result.Message == "" {
			r.Result.Message, _ = dict.SafeTranslate(term.Internal_Error,
				core.GetLang(r.Context, r.Engine))
		}

	case string:
		errorCast := r.Result.Error.(string)
		errText = errorCast
		r.Result.Message = dict.T(errText.(string), core.GetLang(r.Context, r.Engine))

	default:
		errText = "unknown ERROR!!!"
	}

	// if data is one element don't put it in array
	var finalData interface{}
	if data != nil {
		finalData = data
		if len(data) == 1 {
			finalData = data[0]
		}
	}

	if r.abort {
		r.Context.AbortWithStatusJSON(r.status, &Result{
			Message: r.Result.Message,
			Error:   errText,
			Data:    finalData,
		})
	} else {
		r.Context.JSON(r.status, &Result{
			Message: r.Result.Message,
			Error:   errText,
			Data:    finalData,
		})
	}
}
