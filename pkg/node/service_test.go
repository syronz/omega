package node

import (
	"errors"
	"omega/test/core"
	"testing"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//  go test ./pkg/node/ -run TestServiceCreate -v -count=1 -cover
func checkNodeResponse(node, resp Node) (err error) {
	if node.Name != resp.Name {
		err = errors.New("Name isn't equal")
		return
	}
	if node.Phone != resp.Phone {
		err = errors.New("Phone isn't equal")
		return
	}
	return
}

func TestServiceCreate(t *testing.T) {
	engine := core.StartEngine(&Node{})
	engine.DB.Exec("TRUNCATE TABLE nodes;")

	repo := ProvideRepo(engine)
	service := ProvideService(repo)

	samples := []struct {
		in  Node
		out Node
		err error
	}{
		{
			in: Node{
				Name:  "Diako",
				Phone: "07505149171",
			},
			out: Node{
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
		if err = checkNodeResponse(v.out, result); err != nil {
			t.Error(err)
		}
	}

}
