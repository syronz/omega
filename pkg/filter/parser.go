package filter

import (
	"fmt"
	"omega/pkg/helper"
	"regexp"
	"strings"
)

//Parser will break the filter into the sub-query
func Parser(str string) string {
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

	re := regexp.MustCompile(`\w+`)
	arr := re.FindAllString(str, -1)

	pre := arr[0]
	for _, v := range arr {
		if ok, _ := helper.Includes(ops, v); ok {
			isColExist(pre)
		}
		pre = v
	}
	fmt.Println(">>>++ ", arr)

	for k, v := range swap {
		str = strings.Replace(str, k, v, -1)
	}

	return str
}

func isColExist(col string) {
	fmt.Println(">>>> ##", col)
}
