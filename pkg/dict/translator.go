package dict

import (
	"fmt"
	"omega/internal/core/lang"
)

// Term is list of languages
type Term struct {
	En string
	Ku string
	Ar string
}

// thisTerms used for holding language identifier as a string and Term Struct as value
var thisTerms map[string]Term

// SafeTranslate doesn't add !!! around word in case of not exist for translate
func SafeTranslate(str string, language lang.Language, params ...interface{}) (string, bool) {

	term, ok := thisTerms[str]
	if ok {
		var pattern string

		switch language {
		case lang.En:
			pattern = term.En
		case lang.Ku:
			pattern = term.Ku
		case lang.Ar:
			pattern = term.Ar
		default:
			pattern = str
		}

		if params != nil {
			if !(params[0] == nil || params[0] == "") {
				pattern = fmt.Sprintf(pattern, params...)
			}
		}

		return pattern, true

	}

	return "", false

}

// T the requested term
func T(str string, language lang.Language, params ...interface{}) string {

	pattern, ok := SafeTranslate(str, language, params...)
	if ok {
		return pattern
	}

	return "!!! " + str + " !!!"
}

/*
// TranslateArr get an array and translate all of them and return back an array
func (d *Dict) TranslateArr(strs []string, language lang.Language) []string {
	result := make([]string, len(strs))

	for i, v := range strs {
		result[i] = d.Translate(v, language)
	}

	return result

}

// TODO: should be developed for translate words and params
// func (d *Dict) safeTranslate(str interface{}, language string) string {
// 	term, ok := d.Terms[str]
// 	if ok {

// 		switch language {
// 		case "en":
// 			str = term.En
// 		case "ku":
// 			str = term.Ku
// 		case "ar":
// 			str = term.Ar
// 		}

// 	}

// 	return str

// }
*/
