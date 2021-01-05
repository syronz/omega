package service

import (
	"encoding/json"
	"fmt"
	"omega/domain/base/basmodel"
	"omega/domain/base/enum/accounttype"
	"omega/internal/types"
	"omega/pkg/helper"
	"testing"
)

func TestChartOfAccounts(t *testing.T) {
	accounts := []basmodel.Account{
		{
			FixedNode: types.FixedNode{
				ID: 1,
			},
			Code: helper.StrPointer("1"),
			Name: "Asset",
			Type: accounttype.Asset,
		},
		{
			FixedNode: types.FixedNode{
				ID: 2,
			},
			ParentID: types.RowIDPointer(1),
			Code:     helper.StrPointer("11"),
			Name:     "Cash USD",
			Type:     accounttype.Cash,
		},
		{
			FixedNode: types.FixedNode{
				ID: 3,
			},
			ParentID: types.RowIDPointer(1),
			Code:     helper.StrPointer("12"),
			Name:     "Cash IQD",
			Type:     accounttype.Cash,
		},
	}

	for _, v := range accounts {
		fmt.Println(v.Name, v.ID)
	}

	b, err := json.MarshalIndent(accounts, "", " ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

	t.Log("here is we testing", accounts)
}
