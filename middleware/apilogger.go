package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"omega/config"
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
func APILogger(cfg config.CFG) gin.HandlerFunc {
	var requestIndex uint

	return func(c *gin.Context) {
		start := time.Now()
		buf, _ := ioutil.ReadAll(c.Request.Body)
		rdr := ioutil.NopCloser(bytes.NewBuffer(buf))
		//We have to create a new Buffer, because rdr will be read.
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(buf))
		requestIndex++

		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		logRequest(cfg, c, requestIndex, rdr)

		c.Next()

		stop := time.Since(start)
		latency := int(math.Ceil(float64(stop.Nanoseconds()) / 1000000.0))

		logResponse(cfg, c, latency, blw)

	}
}

func logRequest(cfg config.CFG, c *gin.Context, requestIndex uint, rdr io.Reader) {
	cfg.Logapi.WithFields(logrus.Fields{
		"ip":         c.ClientIP(),
		"method":     c.Request.Method,
		"uri":        c.Request.RequestURI,
		"path":       c.Request.URL.Path,
		"request":    readBody(rdr),
		"params":     c.Request.URL.Query(),
		"referer":    c.Request.Referer(),
		"user_agent": c.Request.UserAgent(),
	}).Info(requestIndex)
	c.Set("msgIndex", requestIndex)
}

func logResponse(cfg config.CFG, c *gin.Context, latency int, blw *bodyLogWriter) {
	msgIndex, _ := c.Get("msgIndex")
	cfg.Logapi.WithFields(logrus.Fields{
		"status":      c.Writer.Status(),
		"latency":     latency, // time to process
		"data_length": c.Writer.Size(),
		"response":    readBody(blw.body),
	}).Info(msgIndex)
}

func readBody(reader io.Reader) interface{} {
	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(reader); err != nil {
		fmt.Println(err)
	}

	var obj interface{}

	if err := json.NewDecoder(buf).Decode(&obj); err != nil {
		fmt.Println(err)
	}

	return obj
}
