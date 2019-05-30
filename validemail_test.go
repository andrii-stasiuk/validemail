package validemail

import "testing"

type emailAddresses struct {
	email  string
	result bool
}

var testValues = []emailAddresses{
	{"bjornk@mac.com", true},
	{"stern@gmail.com.", false},
	{"tlinden32@att.net", true},
	{"bowmanbs@gmail.com", true},
	{"khris@mac.book.com", true},
	{"mgreen@yahoo..ca", false},
	{"joglo@com32.cast.net", true},
	{"gar@land@aol.com", false},
	{"rnewman@aol.com.pl.net", true},
	{"camp_bell@att.net", true},
}

func TestEMailValidator(t *testing.T) {
	for _, tt := range testValues {
		if res := EMailValidator(tt.email); res != tt.result {
			t.Error("For", tt.email, "expected", tt.result, "got", res)
		}
	}
}
