package response

import (
	"errors"
	"fmt"
	"omega/internal/core"
	"omega/internal/core/corerr"
	"omega/internal/param"
	"omega/pkg/dict"
	"omega/pkg/limberr"

	"github.com/gin-gonic/gin"
)

// Result is a standard output for success and faild requests
type Result struct {
	Message string                 `json:"message,omitempty"`
	Data    interface{}            `json:"data,omitempty"`
	Error   error                  `json:"error,omitempty"`
	Extra   map[string]interface{} `json:"extra,omitempty"`
	// CustomError corerr.CustomError     `json:"custom_error,omitempty"`
}

// Response holding method related to response
type Response struct {
	Result  Result
	status  int
	Engine  *core.Engine
	Context *gin.Context
	abort   bool
	params  param.Param
	Domain  string
}

// New initiate the Response object
func New(engine *core.Engine, context *gin.Context, domain string) *Response {
	return &Response{
		Engine:  engine,
		Context: context,
		Domain:  domain,
	}
}

// NewParam initiate the Response object and params
func NewParam(engine *core.Engine, context *gin.Context,
	part string, domain string) (*Response, param.Param) {
	params := param.Get(context, engine, part)
	return &Response{
		Engine:  engine,
		Context: context,
		params:  params,
	}, params
}

// Params will return the JWT and uri parameters
func (r *Response) Params(part string) param.Param {
	r.params = param.Get(r.Context, r.Engine, part)
	return r.params
}

// Error is used for add error to the result
func (r *Response) Error(err interface{}, data ...interface{}) *Response {
	if errCast, ok := err.(string); ok {
		r.Result.Error = errors.New(errCast)
	}
	if errCast, ok := err.(error); ok {
		r.Result.Error = errCast
	}
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

func translator(lang dict.Lang) limberr.Translator {
	return func(str string, params ...interface{}) string {
		return dict.T(str, lang, params...)
	}
}

// JSON write ouptut as json
func (r *Response) JSON(data ...interface{}) {
	var parsedError error
	if r.Result.Error != nil {
		r.Result.Error = limberr.AddPath(r.Result.Error, r.Context.Request.RequestURI)

		customError := limberr.GetCustom(r.Result.Error)
		lang := core.GetLang(r.Context, r.Engine)
		errorDocPath := fmt.Sprintf("%v%v.html", r.Engine.Envs[core.ErrPanel], lang)
		r.Result.Error = limberr.ApplyCustom(r.Result.Error,
			corerr.UniqErrorMap[customError], errorDocPath)

		tra := translator(lang)
		r.status, parsedError = limberr.Parse(r.Result.Error, tra)
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
			Error:   parsedError,
			Data:    finalData,
		})
	} else {
		r.Context.JSON(r.status, &Result{
			Message: r.Result.Message,
			Error:   parsedError,
			Data:    finalData,
			// CustomError: r.Result.Error,
		})
	}
}
