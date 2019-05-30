package validemail

import "strings"

func EMailChecker(email string) bool {
	s := strings.Split(email, "@")
	if len(s) == 2 {
		// println(s[0])
		// println(s[1])
		// println(string(s[1][len(s[1])-1]))

		if !(byte(s[1][len(s[1])-1]) >= byte('A') && byte(s[1][len(s[1])-1]) <= byte('z')) {
			println("false")
			return false
		}

		for _, s0 := range s[0] {
			if !(byte(s0) >= byte('A') && byte(s0) <= byte('z')) {
				println("false")
				return false
			}
		}

		for _, s1 := range s[1] {
			if !(byte(s1) >= byte('A') && byte(s1) <= byte('z') || (byte(s1) <= byte('.'))) {
				println("false")
				return false
			}
		}
	} else {
		println("false")
		return false
	}

	return true
}

// func main() {
// 	eMailChecker("as@ges.sh")
// }
