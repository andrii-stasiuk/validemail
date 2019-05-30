package validemail

import "testing"

type emails struct {
	email string
	res   bool
}

var testValues = []emails{
	{"dsfsdfsdf", false},
	{"dsfsdfsdf", false},
	{"dsfsdfsdf", false},
	{"as@ges.sh", true},
	{"dsfsdfsdf", false},
	{"dsfsdfsdf", false},
	{"dsfsdfsdf", false},
}

func TestEMailChecker(t *testing.T) {
	for _, tt := range testValues {
		res := EMailChecker(tt.email)
		if res != tt.res {
			t.Error("For", tt.email, "expected", tt.res, "got", res)
		}
	}
}
