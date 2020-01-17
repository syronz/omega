package response

// import "github.com/gin-gonic/gin"

// Result is a standard output for success and faild requests
type Result struct {
	Status  bool        `json:"status,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
	Count   int         `json:"count,omitempty"`
	Code    int         `json:"code,omitempty"`
}

// Success is for status 2xx
type Success struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Count   int         `json:"count,omitempty"`
}

// Failed is for http status 4xx and 5xx
type Failed struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Code    int         `json:"code,omitempty"`
}

/*
func (r *Response) Success(httpCode int, message string, data interface{}, count uint) {
	r.Context.JSON(httpCode, gin.H{
		"status":  true,
		"message": message,
		"data":    data,
		"count":   count,
	})
	return
}
func (r *Response) Failed(httpCode int, code int, message string, data interface{}) {
	r.Context.JSON(httpCode, gin.H{
		"status":  false,
		"code":    code,
		"message": message,
		"data":    data,
	})
	return
}
*/
