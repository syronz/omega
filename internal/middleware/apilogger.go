package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"math"
	"omega/internal/core"
	"omega/pkg/glog"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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
func APILogger(engine *core.Engine) gin.HandlerFunc {
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

		logRequest(engine, c, reqID, reqDataReader)

		c.Next()

		latency := int(math.Ceil(float64(time.Since(start).Nanoseconds()) / 1000000.0))

		logResponse(engine, c, latency, blw)

	}
}

// Logging Response
func logRequest(engine *core.Engine, c *gin.Context, reqID uint, reqDataReader io.Reader) {
	request := getBody(engine, reqDataReader)
	// prevent to save the passwords
	if strings.Contains(c.Request.URL.Path, "login") {
		request = nil
	}
	engine.APILog.WithFields(logrus.Fields{
		"reqID": reqID,
		// "ip":  c.ClientIP(),
		"method":     c.Request.Method,
		"uri":        c.Request.RequestURI,
		"path":       c.Request.URL.Path,
		"request":    request,
		"params":     c.Request.URL.Query(),
		"referer":    c.Request.Referer(),
		"user_agent": c.Request.UserAgent(),
	}).Info("request")
	c.Set("resID", reqID)
}

// Logging Response
func logResponse(engine *core.Engine, c *gin.Context, latency int, blw *bodyLogWriter) {
	resID, ok := c.Get("resID")
	if !ok {
		glog.Debug("there is no resIndex for element", getBody(engine, blw.body))
	}
	engine.APILog.WithFields(logrus.Fields{
		"resID":       resID,
		"status":      c.Writer.Status(),
		"latency":     latency, // time to process
		"data_length": c.Writer.Size(),
		"response":    getBody(engine, blw.body),
	}).Info("response")
}

// Read body
func getBody(engine *core.Engine, reader io.Reader) interface{} {

	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(reader); err != nil {
		glog.Debug(err)
	}

	var obj interface{}

	if err := json.NewDecoder(buf).Decode(&obj); err != nil {
		if err.Error() != "EOF" {
			glog.Debug(err, obj, err.Error())
		}
	}

	return obj
}
