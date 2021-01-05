package param

import "strings"

func (p *Param) parseWhere(cols []string) (whereStr string, err error) {
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

// ParseWhere combine preConditions and filter with each other
func (p *Param) ParseWhere(cols []string) (whereStr string, err error) {
	return p.parseWhere(cols)
}

// ParseWhereDelete is used when the table has deleted_at column
func (p *Param) ParseWhereDelete(cols []string) (whereStr string, err error) {
	if whereStr, err = p.parseWhere(cols); err != nil {
		return
	}

	if !p.ShowDeletedRows {
		whereStr = " deleted_at is NULL AND " + whereStr
	} else {
		whereStr = " deleted_at is not NULL AND " + whereStr
	}

	return
}
