package core

import (
	"fmt"
	"omega/pkg/dict"
)

// FieldError is a type of error for demonstrate binding problem
type FieldError struct {
	Err    string  `json:"error,omitempty"`
	Fields []Field `json:"fields,omitempty"`
}

// Field is used as an array inside the FieldError
type Field struct {
	Field   string      `json:"field,omitempty"`
	Message string      `json:"message,omitempty"`
	Term    string      `json:"term,omitempty"`
	Params  interface{} `json:"params,omitempty"`
}

func (b *FieldError) Error() string {
	// return fmt.Sprintf("%v", b.Err)
	return fmt.Sprintf("%v, %v", b.Err, b.Fields)
}

func (b *FieldError) Summary() string {
	return fmt.Sprintf("%v", b.Err)
}

// NewFieldError is used for initiate the new node of builder
func NewFieldError(err string) *FieldError {
	fieldError := FieldError{Err: err}
	return &fieldError
}

// Add is used for add new element to the array of fields error
func (b *FieldError) Add(err string, params interface{}, fieldName string) *FieldError {
	var field Field
	field.Term = err
	field.Params = params
	field.Field = fieldName
	b.Fields = append(b.Fields, field)
	return b
}

// Set used for assign Message and Field
func (b *FieldError) Set(msg, fieldName string) *FieldError {
	var field Field
	field.Message = msg
	field.Field = fieldName
	b.Fields = append(b.Fields, field)
	return b
}

// Translate create the final output for fields error
func (b *FieldError) Translate(engine *Engine, lang dict.Lang) {
	for i, v := range b.Fields {
		switch v.Params.(type) {
		case []interface{}:
			params, _ := v.Params.([]interface{})
			b.Fields[i].Message = dict.T(v.Term, lang, params...)
		case []string:
			params, _ := v.Params.([]string)
			convertedParams := make([]interface{}, len(params))
			for i, v := range params {
				convertedParams[i] = v
			}
			b.Fields[i].Message = dict.T(v.Term, lang, convertedParams...)

		default:
			b.Fields[i].Message = dict.T(v.Term, lang, v.Params)
		}
	}
}

// HasError check the length of the Fields
func (b *FieldError) HasError2() bool {
	return len(b.Fields) > 0
}

// BindingError is a type of error for demonstrate binding problem
type BindingError struct {
	Err   string
	Extra interface{}
}

func (b *BindingError) Error() string {
	return fmt.Sprintf("%v", b.Err)
}
