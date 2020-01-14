package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"omega/config"
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
			"clientIP":   clientIP,
			"method":     c.Request.Method,
			"uri":        uri,
			"req_body":   reqBody,
			"req_params": reqParam,
			"referer":    referer,
			"userAgent":  clientUserAgent,
		}).Info(requestIndex)

		c.Set("msgIndex", requestIndex)

		c.Next()

		stop := time.Since(start)
		latency := int(math.Ceil(float64(stop.Nanoseconds()) / 1000000.0))
		statusCode := c.Writer.Status()
		dataLength := c.Writer.Size()
		if dataLength < 0 {
			dataLength = 0
		}
		resBody := readBody(blw.body)

		// statusCode := c.Writer.Status()
		// if statusCode >= 400 {
		//ok this is an request with error, let's make a record for it
		// now print body (or log in your preferred way)
		msgIndex, _ := c.Get("msgIndex")
		fmt.Println("Response body: " + blw.body.String())
		cfg.Debug(blw.body, latency, statusCode, clientIP, clientUserAgent, referer, dataLength)
		cfg.Logapi.WithFields(logrus.Fields{
			"statusCode": statusCode,
			"latency":    latency, // time to process
			"clientIP":   clientIP,
			"method":     c.Request.Method,
			"uri":        uri,
			"req_body":   reqBody,
			"req_params": reqParam,
			"referer":    referer,
			"dataLength": dataLength,
			"userAgent":  clientUserAgent,
			// "response":   stripQutations(blw.body.String()),
			"response": resBody,
		}).Info(msgIndex)

		// }
	}
}

func readBody(reader io.Reader) interface{} {
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)

	// return stripQutations(buf.String())

	var obj interface{}

	// json.Unmarshal(buf, &obj)
	// if err := json.NewEncoder(buf).Encode(&obj); err != nil {
	if err := json.NewDecoder(buf).Decode(&obj); err != nil {
		// res.WriteHeader(res, "whoops", http.StatusInternalServerError)
		// return
		fmt.Println(err)
	}

	fmt.Printf("!!!!!!!!!!!++++++++++++++++++++++++++++++++++ %+v \n\n", obj)

	return obj

	// s := strings.Replace(buf.String(), "\"", "", -1)
	// return s
}

func stripQutations(str string) string {
	return strings.Replace(str, "\"", "", -1)
}
