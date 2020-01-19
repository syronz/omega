package user

import (
	"log"
	"net/http"
	"net/http/httptest"
	"omega/engine"

	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func TestAPICreate(t *testing.T) {
	resp := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	c, r := gin.CreateTestContext(resp)

	r.Use(func(c *gin.Context) {
		c.Set("profile", "myfakeprofile")
	})

	engine := engine.Engine{}
	db, err := gorm.Open("mysql", "root:Qaz1@345@tcp(127.0.0.1:3306)/omega_test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalln(err)
	}
	db.LogMode(true)
	db.AutoMigrate(&User{})
	engine.DB = db
	repo := ProvideRepo(engine)
	service := ProvideService(repo)
	api := ProvideAPI(service)

	// userFeed := User{
	// 	Name:     "Ako",
	// 	Username: "ako1222",
	// 	Password: "0750xxxx",
	// }

	// result, err := service.Save(userFeed)
	// log.Println(">>>>>>>>", result, err)

	r.GET("/test", func(c *gin.Context) {
		// users, err := service.FindAll()
		api.FindAll(c)

		// _, found := c.Get("profile")
		// t.Error(found, users, err)
		c.Status(200)
	})
	c.Request, _ = http.NewRequest(http.MethodGet, "/test", nil)
	r.ServeHTTP(resp, c.Request)
	t.Log("it is working fine", resp)
}

/*
func TestAPIFindByID(t *testing.T) {
	req, err := http.NewRequest("GET", "http://127.0.0.1:8083/api/omega/v1/user/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("id", "1")
	r := httptest.NewRecorder()

	// TODO , need to be fixed or removed
		handler := http.HandlerFunc(nil)

		handler.ServeHTTP(r, req)
		if status := r.Code; status != http.StatusOK {
			t.Errorf("FindByID handler returned wrong status: %v, expected: %v\n", status, http.StatusOK)
		}

		expected := `
					{
					  "count": 0,
					  "data": {
						"ID": 1,
						"CreatedAt": "2020-01-17T04:44:16.734096Z",
						"UpdatedAt": "2020-01-17T05:27:30.45251Z",
						"DeletedAt": null,
						"name": "john",
						"username": "uncle_john",
						"password": "$2a$10$fSBu9h9paoh4ip9huJtn9.t8mxZ8L6/ZGCVuoCHMlRePp0ykbmUB6",
						"extra": {
						  "LastVisit": "2019",
						  "Mark": -15
						}
					  },
					  "message": "",
					  "status": true
					}`

		if r.Body.String() != expected {
			t.Errorf("FindByID returned unexpected body: %v\n, want: %v\n", r.Body.String(), expected)
		}

}
*/
