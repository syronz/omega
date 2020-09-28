package filter

import (
	"fmt"
	"omega/pkg/helper"
	"omega/pkg/limberr"
	"regexp"
	"strings"
)

//Parser will break the filter into the sub-query
func Parser(str string, cols []string) (string, error) {
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

	ops := []string{"eq", "ne", "gt", "lt", "gte", "lte", "like"}

	regCol := regexp.MustCompile(`\w+[\.\w+]*`)
	arr := regCol.FindAllString(str, -1)

	regAfterDot := regexp.MustCompile(`\w+$`)
	var reducedCols []string
	for _, v := range cols {
		if strings.Contains(v, ".") {
			reducedCols = append(reducedCols, regAfterDot.FindString(v))
		}
	}

	cols = append(cols, reducedCols...)

	if len(arr) == 0 {
		return "", fmt.Errorf("filter is not valid")
	}

	pre := arr[0]
	for _, v := range arr {
		if ok, _ := helper.Includes(ops, v); ok {
			if ok, err := helper.Includes(cols, pre); !ok || err != nil {
				if err != nil {
					return "", err
				}
				err := fmt.Errorf("col '%s' not exist", pre)
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
