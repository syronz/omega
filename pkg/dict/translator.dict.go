package dict

import (
	"fmt"
)

// SafeTranslate doesn't add !!! around word in case of not exist for translate
func SafeTranslate(str string, language Language, params ...interface{}) (string, bool) {

	term, ok := thisTerms[str]
	if ok {
		var pattern string

		switch language {
		case En:
			pattern = term.En
		case Ku:
			pattern = term.Ku
		case Ar:
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
func T(str string, language Language, params ...interface{}) string {

	pattern, ok := SafeTranslate(str, language, params...)
	if ok {
		return pattern
	}

	return "!!! " + str + " !!!"
}

/*
// TranslateArr get an array and translate all of them and return back an array
func (d *Dict) TranslateArr(strs []string, language Language) []string {
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
