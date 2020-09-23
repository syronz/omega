package search

import (
	"fmt"
	"omega/internal/param"
	"strings"
)

// Parse is used for handling the variety of search
func ParseDeprecated(params param.Param, pattern string) (whereStr string) {
	// TODO: error should be returned in case the pattern was wrong
	var whereArr []string

	if params.PreCondition != "" {
		whereArr = append(whereArr, params.PreCondition)
	}

	// if params.Filter != "" {
	// 	whereArr = append(whereArr, params.Filter)
	// }

	if params.Search != "" {
		whereArr = cleanSearch(params.Search, whereArr, pattern)
	}

	if len(whereArr) > 0 {
		whereStr = strings.Join(whereArr[:], " AND ")
	}

	return

}

func cleanSearch(search string, whereArr []string, pattern string) []string {
	if strings.Contains(search, ">") {
		conditionsArr := strings.Split(search, "~")
		for _, v := range conditionsArr {
			strArr := strings.Split(v, ">")
			whereArr = append(whereArr, fmt.Sprintf(" %v = '%v' ", strArr[0], strArr[1]))
		}

	} else {
		whereArr = append(whereArr, fmt.Sprintf(pattern, search))
	}

	return whereArr
}
