package service

import (
	"errors"
	"omega/domain/base/basmodel"
	"omega/domain/base/basrepo"
	"omega/internal/core"
	"omega/internal/types"
	"omega/test/kernel"
	"testing"
)

func initPhoneTest() (engine *core.Engine, phoneService BasPhoneServ) {
	queryLog, debugLevel := initServiceTest()

	engine = kernel.StartMotor(queryLog, debugLevel)

	phoneService = ProvideBasPhoneService(basrepo.ProvidePhoneRepo(engine))

	return

}

func TestPhoneCreate(test *testing.T) {
	//we will call the initPhoneTest for starting the generating the engine special for TDD
	//then we fetch the phone service which included the phone repo
	//the engine is skipped

	_, phoneService := initPhoneTest()

	// we create a struct of phone model along with the error
	//then we treat each element of the struct as a test and pass it to the system for test.

	//First test element has no issue and should return NO ERRORS at all.
	//2nd test element has error because the input for phone is more than 8 digits
	//3rd test: ERROR: b/c input for phone is less than 5
	//4th test: ERROR: input for Notes is greater than 255 characters
	testCollector := []struct {
		phone basmodel.Phone
		err   error
	}{
		{
			phone: basmodel.Phone{

				Phone:     "077022222",
				Notes:     "This phone number has been created",
				AccountID: 1,
			},
			err: nil,
		},
		{
			phone: basmodel.Phone{

				Phone:     "07702232133123213213",
				Notes:     "This phone number has been created",
				AccountID: 1,
			},
			err: errors.New("this phone has lenght more than 8 digits"),
		},

		{
			phone: basmodel.Phone{
				Phone:     "077",
				Notes:     "this phone  number has been created",
				AccountID: 1,
			},
			err: errors.New("phone has less than 5 digits"),
		},

		{
			phone: basmodel.Phone{
				Phone: "321332131",
				Notes: "This phone has been created, This phone has been created, This phone has been created, This phone has been created,This phone has been created, This phone has been created, This phone has been created, This phone has been created, This phone has been created,",
			},
			err: errors.New("The length of notes is greater than 255"),
		},
	}

	for _, value := range testCollector {
		_, err := phoneService.Create(value.phone)
		if (value.err == nil && err != nil) || (value.err != nil && err == nil) {
			test.Errorf("\nERROR FOR :::%+v::: \nRETURNS :::%+v:::, \nIT SHOULD BE :::%+v:::", value.phone, err, value.err)
		}

	}
}

func TestPhoneUpdate(test *testing.T) {
	//the engine is skipped
	_, phoneService := initPhoneTest()

	type err error
	collector := []struct {
		phone basmodel.Phone
		err   error
	}{
		{
			phone: basmodel.Phone{
				FixedNode: types.FixedNode{
					ID: 1,
				},
				Phone: "23134142",
				Notes: "phone has been updated",
			},
			err: nil,
		},
		{
			phone: basmodel.Phone{
				FixedNode: types.FixedNode{
					ID: 1,
				},
				Phone: "3131233",
				Notes: "phone has been updated",
			},
			err: errors.New("Phone doesn't exist"),
		},
		{
			phone: basmodel.Phone{
				FixedNode: types.FixedNode{
					ID: 1,
				},
				Notes: "phone has been updated",
			},
			err: errors.New("Phone is required"),
		},
	}

	for _, value := range collector {

		_, err := phoneService.Save(value.phone)
		if (value.err == nil && err != nil) || (value.err != nil && err == nil) {
			test.Errorf("\nERROR FOR :::%+v::: \nRETURNS :::%+v:::, \nIT SHOULD BE :::%+v:::", value.phone, err, value.err)
		}
	}
}

func TestPhoneDelete(t *testing.T) {
	_, phoneService := initPhoneTest()

	testCollector := []struct {
		fix types.FixedNode
		err error
	}{
		{
			fix: types.FixedNode{
				ID:        1,
				CompanyID: 1001,
				NodeID:    101,
			},
			err: nil,
		},
		{
			fix: types.FixedNode{
				ID: 2,
			},
			err: errors.New("phone was not found to be deleted"),
		},
	}

	for _, value := range testCollector {
		_, err := phoneService.Delete(value.fix)
		if (value.err == nil && err != nil) || (value.err != nil && err == nil) {
			t.Errorf("ERROR FOR ::::%+v::: \nRETURNS :::%+v:::, \nIT SHOULD BE :::%+v:::", value.fix.ID, err, value.err)
		}
	}
}
