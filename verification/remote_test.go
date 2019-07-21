package verification_test

import (
	"testing"

	"github.com/nodo/certo/verification"
)

type RemoteTest struct {
	URL        string
	CACertPath string
	Expected   error
}

func TestRemoteValidate(t *testing.T) {
	tt := []RemoteTest{
		{"", "", nil},
		{"https://example.com", "", nil},
		{"", "b", nil},
	}
	for _, test := range tt {
		v := verification.NewRemote(test.URL, test.CACertPath)
		if ok := v.Validate(); ok {
			t.Errorf("Expected Validate() to return false, when URL: %s, CACert: %s\n", test.URL, test.CACertPath)
		}
	}
	if ok := verification.NewRemote("a", "b").Validate(); !ok {
		t.Errorf("Expected Validate() to return true when both URL and cacert are not empty")
	}
}
