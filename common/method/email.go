package method

import "regexp"

func IsEmail(email string) bool {
	emailRegexp := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	if emailRegexp.MatchString(email) {
		return true
	} else {
		return false
	}
}
