package limberr

import (
	"fmt"
)

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
}

// Field is used as an array inside the FieldError
type Field struct {
	Field        string        `json:"field,omitempty"`
	Reason       string        `json:"reason,omitempty"`
	ReasonParams []interface{} `json:"reason_params,omitempty"`
}

func (p *Final) Error() string {
	errStr := fmt.Sprintf("#%v, err:%v", p.Code, p.OriginalError)
	// if len(p.InvalidParams) > 0 {
	// 	errStr += fmt.Sprintf(", invalid_params:%+v", p.InvalidParams)
	// }
	return errStr
}
