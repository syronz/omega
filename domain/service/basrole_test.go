package service

import (
	"errors"
	"omega/domain/base"
	"omega/domain/base/basmodel"
	"omega/domain/base/basrepo"
	"omega/internal/core"
	"omega/internal/param"
	"omega/internal/types"
	"omega/test/kernel"
	"testing"
	"time"
)

func initRoleTest() (engine *core.Engine, roleServ BasRoleServ) {
	logQuery, debugLevel := initServiceTest()
	engine = kernel.StartMotor(logQuery, debugLevel)
	roleServ = ProvideBasRoleService(basrepo.ProvideRoleRepo(engine))

	return
}

func TestRoleCreate(t *testing.T) {
	_, roleServ := initRoleTest()

	samples := []struct {
		in  basmodel.Role
		err error
	}{
		{
			in: basmodel.Role{
				FixedCol: types.FixedCol{
					CompanyID: 1001,
					NodeID:    101,
				},
				Name:        "created 1",
				Resources:   string(base.SuperAccess),
				Description: "created 1",
			},
			err: nil,
		},
		{
			in: basmodel.Role{
				FixedCol: types.FixedCol{
					CompanyID: 1001,
					NodeID:    101,
				},
				Name:        "created 1",
				Resources:   string(base.SuperAccess),
				Description: "created 2",
			},
			err: errors.New("duplicate"),
		},
		{
			in: basmodel.Role{
				FixedCol: types.FixedCol{
					CompanyID: 1001,
					NodeID:    101,
				},
				Name:      "minimum fields",
				Resources: string(base.SuperAccess),
			},
			err: nil,
		},
		{
			in: basmodel.Role{
				Name:        "long name: big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name",
				Resources:   string(base.SuperAccess),
				Description: "created 2",
			},
			err: errors.New("data too long for name"),
		},
		{
			in: basmodel.Role{
				Resources:   string(base.SuperAccess),
				Description: "created 3",
			},
			err: errors.New("name is required"),
		},
	}

	for _, v := range samples {
		_, err := roleServ.Create(v.in)
		if (v.err == nil && err != nil) || (v.err != nil && err == nil) {
			t.Errorf("\nERROR FOR :::%+v::: \nRETURNS :::%+v:::, \nIT SHOULD BE :::%+v:::", v.in, err, v.err)
		}
	}

}

func TestRoleUpdate(t *testing.T) {
	_, roleServ := initRoleTest()

	samples := []struct {
		in  basmodel.Role
		err error
	}{
		{
			in: basmodel.Role{
				FixedCol: types.FixedCol{
					ID:        5,
					CompanyID: 1001,
					NodeID:    101,
				},
				Name:        "num 1 update",
				Resources:   string(base.SuperAccess),
				Description: "num 1 update",
			},
			err: nil,
		},
		{
			in: basmodel.Role{
				FixedCol: types.FixedCol{
					ID:        6,
					CompanyID: 1001,
					NodeID:    101,
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

	samples := []struct {
		fix types.FixedCol
		err error
	}{
		{
			fix: types.FixedCol{
				ID:        7,
				CompanyID: 1001,
				NodeID:    101,
			},
			err: nil,
		},
		{
			fix: types.FixedCol{
				ID:        99999999,
				CompanyID: 1001,
				NodeID:    101,
			},
			err: errors.New("record not found"),
		},
	}

	for _, v := range samples {
		_, err := roleServ.Delete(v.fix)
		if (v.err == nil && err != nil) || (v.err != nil && err == nil) {
			t.Errorf("ERROR FOR ::::%+v::: \nRETURNS :::%+v:::, \nIT SHOULD BE :::%+v:::", v.fix.ID, err, v.err)
		}
	}
}

func TestRoleList(t *testing.T) {
	_, roleServ := initRoleTest()
	regularParam := getRegularParam("bas_roles.id asc")
	regularParam.Filter = "description[like]'searchTerm1'"

	samples := []struct {
		params param.Param
		count  int64
		err    error
	}{
		{
			params: param.Param{},
			err:    nil,
			count:  13,
		},
		{
			params: regularParam,
			err:    nil,
			count:  3,
		},
	}

	for _, v := range samples {
		_, count, err := roleServ.List(v.params)

		if (v.err == nil && err != nil) || (v.err != nil && err == nil) || count != v.count {
			t.Errorf("FOR :::%+v::: \nRETURNS :::%+v:::, \nIT SHOULD BE :::%+v:::", v.params, count, v.count)
		}
		time.Sleep(1 * time.Second)
	}
}

func TestRoleExcel(t *testing.T) {
	_, roleServ := initRoleTest()
	regularParam := getRegularParam("bas_roles.id asc")

	samples := []struct {
		params param.Param
		count  int64
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
		if (v.err == nil && err != nil) || (v.err != nil && err == nil) || int64(len(data)) < v.count {
			t.Errorf("FOR ::::%+v::: \nRETURNS :::%+v:::, \nIT SHOULD BE :::%+v::: \nErr :::%+v:::",
				v.params, int64(len(data)), v.count, err)
		}
	}
}
