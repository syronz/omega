package service

import (
	"errors"
	"omega/domain/base/basmodel"
	"omega/domain/base/basrepo"
	"omega/domain/base/basresource"
	"omega/internal/core"
	"omega/internal/param"
	"omega/internal/types"
	"omega/test/kernel"
	"testing"
)

func initRoleTest() (engine *core.Engine, roleServ BasRoleServ) {
	logQuery, debugLevel := initServiceTest()
	engine = kernel.StartMotor(logQuery, debugLevel)
	roleServ = ProvideBasRoleService(basrepo.ProvideBasRoleRepo(engine))

	return
}

func TestRoleCreate(t *testing.T) {
	_, roleServ := initRoleTest()
	regularParam := getRegularParam("bas_roles.id asc")

	samples := []struct {
		in     basmodel.BasRole
		params param.Param
		err    error
	}{
		{
			in: basmodel.BasRole{
				Name:        "created 1",
				Resources:   basresource.SupperAccess,
				Description: "created 1",
			},
			params: regularParam,
			err:    nil,
		},
		{
			in: basmodel.BasRole{
				Name:        "created 1",
				Resources:   basresource.SupperAccess,
				Description: "created 1",
			},
			params: regularParam,
			err:    errors.New("duplicate"),
		},
		{
			in: basmodel.BasRole{
				Name:      "minimum fields",
				Resources: basresource.SupperAccess,
			},
			params: regularParam,
			err:    nil,
		},
		{
			in: basmodel.BasRole{
				Name:        "long name: big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name",
				Resources:   basresource.SupperAccess,
				Description: "created 2",
			},
			params: regularParam,
			err:    errors.New("data too long for name"),
		},
		{
			in: basmodel.BasRole{
				Resources:   basresource.SupperAccess,
				Description: "created 3",
			},
			params: regularParam,
			err:    errors.New("name is required"),
		},
	}

	for _, v := range samples {
		_, err := roleServ.Create(v.in, v.params)
		if (v.err == nil && err != nil) || (v.err != nil && err == nil) {
			t.Errorf("\nERROR FOR :::%+v::: \nRETURNS :::%+v:::, \nIT SHOULD BE :::%+v:::", v.in, err, v.err)
		}
	}

}

func TestRoleUpdate(t *testing.T) {
	_, roleServ := initRoleTest()

	samples := []struct {
		in  basmodel.BasRole
		err error
	}{
		{
			in: basmodel.BasRole{
				GormCol: types.GormCol{
					ID: 1001101000000005,
				},
				Name:        "num 1 update",
				Resources:   basresource.SupperAccess,
				Description: "num 1 update",
			},
			err: nil,
		},
		{
			in: basmodel.BasRole{
				GormCol: types.GormCol{
					ID: 1001101000000006,
				},
				Name:        "num 2 update",
				Description: "num 2 update",
			},
			err: errors.New("resources are required"),
		},
	}

	for _, v := range samples {
		_, err := roleServ.Save(v.in)
		if (v.err == nil && err != nil) || (v.err != nil && err == nil) {
			t.Errorf("ERROR FOR ::::%+v::: \nRETURNS :::%+v:::, \nIT SHOULD BE :::%+v:::", v.in, err, v.err)
		}
	}

}

func TestRoleDelete(t *testing.T) {
	_, roleServ := initRoleTest()
	regularParam := getRegularParam("bas_roles.id asc")

	samples := []struct {
		id  types.RowID
		err error
	}{
		{
			id:  7,
			err: nil,
		},
		{
			id:  99999999,
			err: errors.New("record not found"),
		},
	}

	for _, v := range samples {
		_, err := roleServ.Delete(v.id, regularParam)
		if (v.err == nil && err != nil) || (v.err != nil && err == nil) {
			t.Errorf("ERROR FOR ::::%+v::: \nRETURNS :::%+v:::, \nIT SHOULD BE :::%+v:::", v.id, err, v.err)
		}
	}
}

func TestRoleList(t *testing.T) {
	_, roleServ := initRoleTest()
	regularParam := getRegularParam("bas_roles.id asc")
	regularParam.Search = "searchTerm1"

	samples := []struct {
		params param.Param
		count  uint64
		err    error
	}{
		{
			params: param.Param{},
			err:    errors.New("error in url"),
			count:  0,
		},
		{
			params: regularParam,
			err:    nil,
			count:  3,
		},
	}

	for _, v := range samples {
		data, err := roleServ.List(v.params)
		var count uint64
		var ok bool
		if count, ok = data["count"].(uint64); !ok {
			count = 0
		}
		if (v.err == nil && err != nil) || (v.err != nil && err == nil) || count != v.count {
			t.Errorf("FOR ::::%+v::: \nRETURNS :::%+v:::, \nIT SHOULD BE :::%+v:::", v.params, data["count"], v.count)
		}
	}
}

func TestRoleExcel(t *testing.T) {
	_, roleServ := initRoleTest()
	regularParam := getRegularParam("bas_roles.id asc")

	samples := []struct {
		params param.Param
		count  uint64
		err    error
	}{
		{
			params: regularParam,
			err:    nil,
			count:  6,
		},
	}

	for _, v := range samples {
		data, err := roleServ.Excel(v.params)
		if (v.err == nil && err != nil) || (v.err != nil && err == nil) || uint64(len(data)) < v.count {
			t.Errorf("FOR ::::%+v::: \nRETURNS :::%+v:::, \nIT SHOULD BE :::%+v::: \nErr :::%+v:::",
				v.params, uint64(len(data)), v.count, err)
		}
	}
}
