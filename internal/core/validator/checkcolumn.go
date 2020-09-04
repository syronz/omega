package validator

import (
	"omega/internal/param"
)

// CheckColumns will check columns for security
func CheckColumns(cols []string, variate string, params param.Param) (string, error) {
	// fieldError := core.NewFieldError(term.Error_in_url)
	// fieldError := corerr.NewSilent("E1059215", params, base.Domain, nil)

	// if variate == "*" {
	// 	return strings.Join(cols, ","), nil
	// }

	// variates := strings.Split(variate, ",")
	// for _, v := range variates {
	// 	if ok, _ := helper.Includes(cols, v); !ok {
	// 		fieldError.Add(corerr.V_is_not_valid, v, strings.Join(cols, ", "))
	// 	}
	// }

	return variate, nil

}
