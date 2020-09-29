package limberr

import (
	"errors"
	"log"
	"net/http"
)

// Translator is an outside function for localize the string
type Translator func(string, ...interface{}) string

// Parse convert chained error to the Final format for send in JSON format
func Parse(err error, translator Translator) (int, error) {
	var final Final
	var status int

	for err != nil {
		switch e := err.(type) {
		case interface{ Unwrap() error }:
			err = errors.Unwrap(err)
		case *WithMessage:
			if final.Message == "" {
				final.Message = translator(e.Msg, e.Params...)
			}
			err = e.Err
		case *WithCode:
			if final.Code == "" {
				final.Code = e.Code
			}
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
			status = e.Status
			err = e.Err
		case *WithDomain:
			final.Domain = e.Domain
			err = e.Err
		case *WithInvalidParam:
			field := Field{
				Field:        e.Field,
				Reason:       translator(e.Reason, e.Params...),
				ReasonParams: e.Params,
			}
			final.InvalidParams = append(final.InvalidParams, field)
			err = e.Err
		case *WithCustom:
			err = e.Err
		case error:
			final.OriginalError += e.Error()
			err = errors.Unwrap(err)
		default:
			log.Println("There shouldn't be a default error", err)
			return http.StatusInternalServerError, &final
		}
	}
	return status, &final
}

func GetCustom(err error) (customError CustomError) {
	for err != nil {
		switch e := err.(type) {
		case interface{ Unwrap() error }:
			err = errors.Unwrap(err)
		case *WithCustom:
			return e.Custom
		case error:
			if errCast, ok := e.(*WithMessage); ok {
				err = errCast.Err
				continue
			}
			if errCast, ok := e.(*WithCode); ok {
				err = errCast.Err
				continue
			}
			if errCast, ok := e.(*WithType); ok {
				err = errCast.Err
				continue
			}
			if errCast, ok := e.(*WithPath); ok {
				err = errCast.Err
				continue
			}
			if errCast, ok := e.(*WithStatus); ok {
				err = errCast.Err
				continue
			}
			if errCast, ok := e.(*WithDomain); ok {
				err = errCast.Err
				continue
			}
			if errCast, ok := e.(*WithInvalidParam); ok {
				err = errCast.Err
				continue
			}
			return
		default:
			log.Println("There shouldn't be a default for getting custom", err)
			return
		}
	}

	return
}

func ApplyCustom(err error, theme ErrorTheme, errPage string) error {
	err = AddType(err, errPage+theme.Type, theme.Title)
	err = AddDomain(err, theme.Domain)
	err = AddStatus(err, theme.Status)
	return err

}

func appendText(str string, txt string) (result string) {
	if str == "" {
		result = txt
	} else {
		result = str + ", " + txt
	}
	return
}
