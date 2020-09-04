package limberr

import (
	"errors"
	"log"
)

// Translator is an outside function for localize the string
type Translator func(string, ...interface{}) string

// Parse convert chained error to the Final format for send in JSON format
func Parse(err error, translator Translator) error {
	var final Final

	for err != nil {
		switch e := err.(type) {
		case interface{ Unwrap() error }:
			err = errors.Unwrap(err)
		case *WithMessage:
			final.Message = e.Msg
			err = e.Err
		case *WithCode:
			final.Code = e.Code
			err = e.Err
		case *WithType:
			final.Type = e.Type
			final.Title = translator(e.Title)
			err = e.Err
		case *WithPath:
			final.Path += appendText(final.Path, e.Path)
			err = e.Err
		case *WithStatus:
			final.Status = e.Status
			err = e.Err
		case *WithDomain:
			final.Domain = e.Domain
			err = e.Err
		case error:
			final.OriginalError += e.Error()
			err = errors.Unwrap(err)
		default:
			log.Println("There shouldn't be a default error", err)
			return &final
		}
	}
	return &final
}

func appendText(str string, txt string) (result string) {
	if str == "" {
		result = txt
	} else {
		result = str + ", " + txt
	}
	return
}
