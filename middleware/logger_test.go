package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"omega/internal/core"

	// "io"
	"io/ioutil"
	"net/http"
	"net/url"
	"omega/pkg/auth"
	"strings"
	"testing"

	"omega/internal/glog"

	"github.com/gin-gonic/gin"
)

func TestLogRequest(t *testing.T) {
	_ = core.StartEngine()

	bbo := ioutil.NopCloser(strings.NewReader("{\"password\":\"k2i\",\"username\":\"diako\"}"))
	buf := new(bytes.Buffer)
	buf.ReadFrom(bbo)
	bbo.Close()

	u := url.URL{
		Path: "this is url",
	}

	authObj := auth.Auth{
		Username: "user1",
		Password: "pass1",
	}

	payloadBytes, err := json.Marshal(authObj)
	if err != nil {
		glog.Debug(err)
	}
	body := bytes.NewReader(payloadBytes)
	req, err := http.NewRequest("GET", "127.0.0.1/login", body)
	if err != nil {
		glog.Debug(err)
	}

	req.Header.Add("X-Forwarded-For", "127.0.0.1")
	// buf, _ := ioutil.ReadAll(c.Request.Body)
	// rdr := ioutil.NopCloser(bytes.NewBuffer(buf))

	r := http.Request{

		Method: "post",
		// Body:       ioutil.NopCloser(strings.NewReader("Hello Diako")),
		RequestURI: "/url/good",
		URL:        &u,
	}

	c := gin.Context{
		Request: req,
	}

	// var c *gin.Context
	requestIndex := uint(25)

	// c.Request.Method = "post"

	// logRequest(&c, 25, bbo)
	bodyResult := getBody(bbo)
	glog.Debug("@@@@@@@@@@@@@@@@@ THIS IS DEBUG @@@@@@@", bodyResult, req.Body)

	t.Log("This is logRequest test", fmt.Sprintf("%T :: %+[1]v", r), c, requestIndex, bbo)

}
