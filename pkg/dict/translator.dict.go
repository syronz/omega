package dict

import (
	"fmt"
)

// SafeTranslate doesn't add !!! around word in case of not exist for translate
func SafeTranslate(str string, lang Lang, params ...interface{}) (string, bool) {
	if !translateInBackend {
		return str, true
	}

	term, ok := thisTerms[str]
	if ok {
		var pattern string

		switch lang {
		case En:
			pattern = term.En
		case Ku:
			pattern = term.Ku
		case Ar:
			pattern = term.Ar
		default:
			pattern = str
		}

		// if type of param is dict.R then translate it
		for i, v := range params {
			switch v.(type) {
			case R:
				term := v.(R)
				params[i] = T(string(term), lang)
			}
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
func T(str string, lang Lang, params ...interface{}) string {
	if !translateInBackend {
		return str
	}

	pattern, ok := SafeTranslate(str, lang, params...)
	if ok {
		return pattern
	}

	return "!!! " + str + " !!!"
}

/*
// TranslateArr get an array and translate all of them and return back an array
func (d *Dict) TranslateArr(strs []string, lang Lang) []string {
	result := make([]string, len(strs))

	for i, v := range strs {
		result[i] = d.Translate(v, lang)
	}

	return result

}

// TODO: should be developed for translate words and params
// func (d *Dict) safeTranslate(str interface{}, lang string) string {
// 	term, ok := d.Terms[str]
// 	if ok {

// 		switch lang {
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
