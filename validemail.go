package validemail

import "strings"

// EMailChecker function checks if email is correct
func EMailChecker(email string) bool {
	s := strings.Split(email, "@")
	if len(s) == 2 {
		if !(byte(s[1][len(s[1])-1]) >= byte('A') && byte(s[1][len(s[1])-1]) <= byte('z')) {
			return false
		}
		for _, s0 := range s[0] {
			if !(byte(s0) >= byte('A') && byte(s0) <= byte('z')) {
				return false
			}
		}
		for _, s1 := range s[1] {
			if !(byte(s1) >= byte('A') && byte(s1) <= byte('z') || (byte(s1) <= byte('.'))) {
				return false
			}
		}
	} else {
		return false
	}
	return true
}
