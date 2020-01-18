package search

import (
	"fmt"
	"omega/internal/param"
	"strings"
)

// Parse is used for handling the variety of search
func Parse(params param.Param, pattern string) (whereStr string) {
	// TODO: error should be returned in case the pattern was wrong
	var whereArr []string

	if params.PreCondition != "" {
		whereArr = append(whereArr, params.PreCondition)
	}

	if strings.Contains(params.Search, ">") {
		conditionsArr := strings.Split(params.Search, "~")
		for _, v := range conditionsArr {
			strArr := strings.Split(v, ">")
			whereArr = append(whereArr, fmt.Sprintf(" %v = '%v' ", strArr[0], strArr[1]))
		}

	} else {
		whereArr = append(whereArr, fmt.Sprintf(pattern, params.Search))
	}

	if len(whereArr) > 0 {
		whereStr = strings.Join(whereArr[:], " AND ")
	}

	return

}
