package node

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"omega/test/core"
	"testing"
)

func TestAPICreate(t *testing.T) {
	resp := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	c, r := gin.CreateTestContext(resp)

	r.Use(func(c *gin.Context) {
		c.Set("profile", "myfakeprofile")
	})

	engine := core.StartEngine(&Node{})

	repo := ProvideRepo(engine)
	service := ProvideService(repo)
	api := ProvideAPI(service)

	r.GET("/test", func(c *gin.Context) {
		api.FindAll(c)

		c.Status(200)
	})
	c.Request, _ = http.NewRequest(http.MethodGet, "/test", nil)
	r.ServeHTTP(resp, c.Request)
	t.Log("it is working fine", resp)
}
