package verification_test

import (
	"testing"

	"github.com/nodo/certo/verification"
)

type LocalTest struct {
	CertPath   string
	CACertPath string
	Expected   error
}

func TestLocalValidate(t *testing.T) {
	tt := []LocalTest{
		{"", "", nil},
		{"a", "", nil},
		{"", "b", nil},
	}
	for _, test := range tt {
		v := verification.NewLocal(test.CertPath, test.CACertPath)
		if ok := v.Validate(); ok {
			t.Errorf("Expected Validate() to return false, when Cert: %s, CACert: %s\n", test.CertPath, test.CACertPath)
		}
	}
	if ok := verification.NewLocal("a", "b").Validate(); !ok {
		t.Errorf("Expected Validate() to return true when both cert and cacert are not empty")
	}
}
