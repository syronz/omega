package filter

import (
	"fmt"
	"omega/pkg/helper"
	"omega/pkg/limberr"
	"regexp"
	"strings"
)

//Parser will break the filter into the sub-query
func Parser(str, cols string) (string, error) {
	swap := make(map[string]string)
	swap["[eq]"] = " = "
	swap["[ne]"] = " != "
	swap["[gt]"] = " > "
	swap["[lt]"] = " < "
	swap["[gte]"] = " >= "
	swap["[lte]"] = " <= "
	swap["[like]"] = " LIKE "
	swap["[and]"] = " AND "
	swap["[or]"] = " OR "

	ops := []string{"eq", "ne", "gt", "lt",
		"gte", "lte", "like"}

	re := regexp.MustCompile(`\w+[\.\w+]*`)
	arr := re.FindAllString(str, -1)

	pre := arr[0]
	for _, v := range arr {
		if ok, _ := helper.Includes(ops, v); ok {
			if !strings.Contains(cols, pre) {
				err := fmt.Errorf("col %v not exist", pre)
				err = limberr.AddInvalidParam(err, pre,
					"column %v not not exist", pre)
				return "", err
			}
		}
		pre = v
	}

	for k, v := range swap {
		str = strings.Replace(str, k, v, -1)
	}

	return str, nil
}
