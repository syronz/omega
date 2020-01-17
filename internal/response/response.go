package response

import "github.com/gin-gonic/gin"

type Response struct {
	Context *gin.Context
}

func (r *Response) Success(httpCode int, message string, data interface{}, count uint) {
	r.Context.JSON(httpCode, gin.H{
		"status": true,
		"message": message,
		"data":    data,
		"count":   count,
	})
	return
}
func (r *Response) Failed(httpCode int, code int, message string, data interface{}) {
	r.Context.JSON(httpCode, gin.H{
		"status": false,
		"code": code,
		"message": message,
		"data": data,
	})
	return
}
