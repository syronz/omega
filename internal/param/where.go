package param

import "strings"

// ParseWhere combine preConditions and filter with each other
func (p *Param) ParseWhere(cols []string) (whereStr string, err error) {
	var whereArr []string
	var resultFilter string
	//TODO: pack each part in seperate brackets

	if resultFilter, err = p.parseFilter(cols); err != nil {
		return
	}

	if resultFilter != "" {
		whereArr = append(whereArr, resultFilter)
	}

	if p.PreCondition != "" {
		whereArr = append(whereArr, p.PreCondition)
	}

	if len(whereArr) > 0 {
		whereStr = strings.Join(whereArr[:], " AND ")
	}

	return
}
