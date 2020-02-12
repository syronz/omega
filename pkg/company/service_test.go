package company

import (
	"errors"
	"omega/test/core"
	"testing"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//  go test ./pkg/company/ -run TestServiceCreate -v -count=1 -cover
func checkCompanyResponse(company, resp Company) (err error) {
	if company.Name != resp.Name {
		err = errors.New("Name isn't equal")
		return
	}
	if company.Phone != resp.Phone {
		err = errors.New("Phone isn't equal")
		return
	}
	return
}

func TestServiceCreate(t *testing.T) {
	engine := core.StartEngine(&Company{})
	engine.DB.Exec("TRUNCATE TABLE companies;")

	repo := ProvideRepo(engine)
	service := ProvideService(repo)

	samples := []struct {
		in  Company
		out Company
		err error
	}{
		{
			in: Company{
				Name:  "Diako",
				Phone: "07505149171",
			},
			out: Company{
				Name:  "Diako",
				Phone: "07505149171",
			},
			err: nil,
		},
	}

	for _, v := range samples {
		result, err := service.Save(v.in)
		if err != nil {
			t.Error(err)
		}
		if err = checkCompanyResponse(v.out, result); err != nil {
			t.Error(err)
		}
	}

}
