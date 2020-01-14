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

func Wrapper(cfg config.CFG) gin.HandlerFunc {
	var requestIndex uint

	// func GinBodyLogMiddleware(c *gin.Context) {
	return func(c *gin.Context) {
		buf, _ := ioutil.ReadAll(c.Request.Body)
		rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
		rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf)) //We have to create a new Buffer, because rdr1 will be read.
		path := c.Request.URL.Path
		uri := c.Request.RequestURI
		reqParam := c.Request.URL.Query()
		start := time.Now()
		reqBody := readBody(rdr1)

		fmt.Println(reqBody, path, uri, reqParam, start) // Print request body

		c.Request.Body = rdr2
		requestIndex++

		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		clientIP := c.ClientIP()
		referer := c.Request.Referer()
		clientUserAgent := c.Request.UserAgent()
		cfg.Logapi.WithFields(logrus.Fields{
			"ip":        clientIP,
			"method":    c.Request.Method,
			"uri":       uri,
			"request":   reqBody,
			"params":    reqParam,
			"referer":   referer,
			"userAgent": clientUserAgent,
		}).Info(requestIndex)
		c.Set("msgIndex", requestIndex)

		c.Next()

		stop := time.Since(start)
		latency := int(math.Ceil(float64(stop.Nanoseconds()) / 1000000.0))
		statusCode := c.Writer.Status()
		dataLength := c.Writer.Size()
		// if dataLength < 0 {
		// 	dataLength = 0
		// }
		resBody := readBody(blw.body)

		msgIndex, ok := c.Get("msgIndex")
		if !ok {
			msgIndex = -1
		}
		cfg.Logapi.WithFields(logrus.Fields{
			"status":      statusCode,
			"latency":     latency, // time to process
			"data_length": dataLength,
			"response":    resBody,
		}).Info(msgIndex)

	}
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
