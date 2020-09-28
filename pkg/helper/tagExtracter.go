package helper

import (
	"fmt"
	"reflect"
	"regexp"
)

type extractor struct {
	arr   []string
	re    *regexp.Regexp
	table string
}

// TagExtracter extract the name of table and field from json and table tag
func TagExtracter(t reflect.Type, table string) []string {
	ext := extractor{
		arr:   []string{},
		re:    regexp.MustCompile(`\w+`),
		table: table,
	}

	ext.getTag(t)

	return ext.arr
}

func (p *extractor) getTag(t reflect.Type) {
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		if field.Type.Kind() == reflect.Struct {
			p.getTag(field.Type)
			continue
		}

		column := field.Tag.Get("json")
		column = p.re.FindString(column)
		externalTable := field.Tag.Get("table")

		switch {
		case externalTable == "-":
			continue
		case externalTable != "":
			column = externalTable
		default:
			column = fmt.Sprintf("%v.%v", p.table, column)
		}

		p.arr = append(p.arr, column)
	}

}
