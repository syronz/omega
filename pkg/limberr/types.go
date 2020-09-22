package limberr

import (
	"fmt"
)

// Final is hold different parts of the error based on rfc 7807
type Final struct {
	Code          string        `json:"code,omitempty"`
	Type          string        `json:"type,omitempty"`
	Title         string        `json:"title,omitempty"`
	Domain        string        `json:"domain,omitempty"`
	Message       string        `json:"message,omitempty"`
	MessageParams []interface{} `json:"message_params,omitempty"`
	Path          string        `json:"path,omitempty"`
	InvalidParams []Field       `json:"invalid_params,omitempty"`
	Status        int           `json:"-"`
	OriginalError string        `json:"original_error,omitempty"`
	Limb          error         `json:"limb,omitempty"`
}

// Field is used as an array inside the FieldError
type Field struct {
	Field        string        `json:"field,omitempty"`
	Reason       string        `json:"reason,omitempty"`
	ReasonParams []interface{} `json:"reason_params,omitempty"`
}

func (p *Final) Error() string {
	errStr := fmt.Sprintf("#%v, err:%v", p.Code, p.OriginalError)
	return errStr
}

// CustomError is used for defining errors related to this application, this is a bridge between the
// limberr and local configuration of the error. It holds error theme
type CustomError int

//ErrorTheme hold the error's type and title
type ErrorTheme struct {
	Type   string
	Title  string
	Domain string
	Status int
}

// CustomErrorMap is used for defining the error for each domain
type CustomErrorMap map[CustomError]ErrorTheme
