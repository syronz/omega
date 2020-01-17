package user

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// User entity unit tests

func TestAPI_FindByID(t *testing.T) {
	req, err := http.NewRequest("GET", "http://127.0.0.1:8083/api/omega/v1/user/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("id", "1")
	r := httptest.NewRecorder()

	// TODO , need to be fixed
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