package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"math"
	"omega/internal/glog"
	"time"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// APILogger is used to save requests and response by using logapi
func APILogger() gin.HandlerFunc {
	var reqID uint

	return func(c *gin.Context) {
		start := time.Now()
		buf, _ := ioutil.ReadAll(c.Request.Body)
		reqDataReader := ioutil.NopCloser(bytes.NewBuffer(buf))
		//We have to create a new Buffer, because reqDataReader will be read.
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(buf))
		reqID++

		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		logRequest(c, reqID, reqDataReader)

		c.Next()

		stop := time.Since(start)
		latency := int(math.Ceil(float64(stop.Nanoseconds()) / 1000000.0))

		logResponse(c, latency, blw)

	}
}

// Logging Response
func logRequest(c *gin.Context, reqId uint, reqDataReader io.Reader) {
	glog.Glog.ApiLog.WithFields(logrus.Fields{
		"reqID": reqId,
		// "ip":  c.ClientIP(),
		"method":     c.Request.Method,
		"uri":        c.Request.RequestURI,
		"path":       c.Request.URL.Path,
		"request":    getBody(reqDataReader),
		"params":     c.Request.URL.Query(),
		"referer":    c.Request.Referer(),
		"user_agent": c.Request.UserAgent(),
	}).Info("request")
	c.Set("resID", reqId)
}

// Logging Response
func logResponse(c *gin.Context, latency int, blw *bodyLogWriter) {
	resID, ok := c.Get("resID")
	if !ok {
		glog.Debug("there is no resIndex for element", getBody(blw.body))
	}
	glog.Glog.ApiLog.WithFields(logrus.Fields{
		"resID":       resID,
		"status":      c.Writer.Status(),
		"latency":     latency, // time to process
		"data_length": c.Writer.Size(),
		"response":    getBody(blw.body),
	}).Info("response")
}

// Read body
func getBody(reader io.Reader) interface{} {
	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(reader); err != nil {
		glog.Debug(err)
	}

	var obj interface{}

	if err := json.NewDecoder(buf).Decode(&obj); err != nil {
		glog.Debug(err)
	}

	return obj
}
