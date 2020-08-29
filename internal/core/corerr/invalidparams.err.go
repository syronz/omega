package corerr

import (
	"net/http"
	"omega/pkg/dict"
)

// FieldError is a type of error for demonstrate binding problem
type FieldError struct {
	Err    string  `json:"error,omitempty"`
	Fields []Field `json:"fields,omitempty"`
}

// NewInvalidParams is used for initiate the new node of builder
// func NewInvalidParams(err string) *FieldError {
// 	fieldError := FieldError{
// 		Err: err,
// 	}
// 	return &fieldError
// }

// Add is used for add new element to the array of fields error
func (p *CustomError) Add(err string, params interface{}, fieldName string) *CustomError {
	var field Field
	field.Term = err
	field.Params = params
	field.Field = fieldName
	p.InvalidParams = append(p.InvalidParams, field)
	return p
}

// NewInvalidParams is used when findbyid returns nill
func (p CustomError) NewInvalidParams(part string, path string) *CustomError {
	part = dict.T(part, p.Lang)
	return &CustomError{
		Code:          p.Code,
		Domain:        p.Domain,
		Type:          p.ErrPanel + string(p.Lang) + ".html#VALIDATION_FAILED",
		Title:         dict.T(Validation_failed, p.Lang),
		Message:       dict.T(Record__NotFoundIn_, p.Lang, part),
		Path:          path,
		Status:        http
		.StatusUnprocessableEntity,
		OriginalError: p.OriginalError,
	}
}

func (p *CustomError) Final() error {
	if len(p.InvalidParams) > 0 {
		return p
	}
	return nil
}
