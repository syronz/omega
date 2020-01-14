package middleware

import (
	// "encoding/json"
	"omega/config"

	"github.com/gin-gonic/gin"
	// "io/ioutil"
	// "math"
	// "os"
	// "time"
)

//var timeFormat = "02/Jan/2006:15:04:05 -0700"
//var timeFormat = "2000/12/30:15:04:05 -0700"

func Logger(cfg config.CFG) gin.HandlerFunc {
	// hostname, err := os.Hostname()
	// if err != nil {
	// 	hostname = "unknown"

	// }
	return func(c *gin.Context) {

		cfg.Log.Warn("this is logger middleware ########################")
		// tmp := c.Request.Body
		// _, err := ioutil.ReadAll(tmp)
		// if err != nil {
		// 	cfg.Debug(err)
		// }

		// var objmap interface{}
		// if err := c.Bind(&objmap); err != nil {
		// 	c.JSON(400, gin.H{"error": err.Error()})
		// 	return
		// }
		// // var objmap map[string]interface{}
		// err = json.Unmarshal(reqBody, &objmap)
		// if err != nil {
		// 	cfg.Debug(err)
		// }

		// cfg.Debug(objmap)
		c.Next()

		// cfg.Debug(c.Response)

		/*
			path := c.Request.URL.Path
			req := c.Request.RequestURI
			reqBody, _ := ioutil.ReadAll(c.Request.Body)
			reqParam := c.Request.URL.Query()
			start := time.Now()
			c.Next()
			stop := time.Since(start)
			latency := int(math.Ceil(float64(stop.Nanoseconds()) / 1000000.0))
			statusCode := c.Writer.Status()
			clientIP := c.ClientIP()
			clientUserAgent := c.Request.UserAgent()
			referer := c.Request.Referer()
			dataLength := c.Writer.Size()
			if dataLength < 0 {
				dataLength = 0
			}

			var objmap map[string]interface{}
			json.Unmarshal(reqBody, &objmap)

			entry := logger.WithFields(logrus.Fields{
				"hostname":   hostname,
				"statusCode": statusCode,
				"latency":    latency, // time to process
				"clientIP":   clientIP,
				"method":     c.Request.Method,
				"req":        req,
				"req_body":   objmap,
				"req_params": reqParam,
				"referer":    referer,
				"dataLength": dataLength,
				"userAgent":  clientUserAgent,
			})

			if len(c.Errors) > 0 {
				entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
			} else {
				msg := fmt.Sprintf("%s - %s [%s] \"%s %s\" %d %d \"%s\" \"%s\" (%dms)", clientIP, hostname, time.Now().Format(time.RFC3339), c.Request.Method, path, statusCode, dataLength, referer, clientUserAgent, latency)
				if statusCode > 499 {
					entry.Error(msg)
				} else if statusCode > 399 {
					entry.Warn(msg)
				} else {
					entry.Info(msg)
				}
			}
		*/
	}

}
