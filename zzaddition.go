package validator

import (
	"regexp"
)

const (
	phoneRegexString = "^1[0-9]{10}$"
)

var (
	phoneRegex = regexp.MustCompile(phoneRegexString)
)

func init() {
	bakedInValidators["phone"] = isPhone
}

func isPhone(f1 FieldLevel) bool {
	return phoneRegex.MatchString(f1.Field().String())
}
