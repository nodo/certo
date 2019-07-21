package actions_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/nodo/certo/actions"
)

type DecodeTest struct {
	Path     string
	Format   string
	Expected error
}

func TestDecode(t *testing.T) {
	tt := []DecodeTest{
		{"fixtures/correct.pem", "json", nil},
		{"fixtures/correct.pem", "yaml", nil},
		{"fixtures/correct.pem", "random", nil},
		{"fixtures/non-existent.pem", "random", errors.New("could not read the certificate file")},
		{"fixtures/invalid.pem", "random", errors.New("could not parse the certificate")},
	}
	for _, test := range tt {
		_, err := actions.Decode(test.Path, test.Format)
		if test.Expected == nil && err != nil {
			t.Errorf(`expected no error, but got "%v"\n`, err)
			continue
		}
		if test.Expected != nil && err == nil {
			t.Errorf(`expected "%v", got no error\n`, test.Expected)
			continue
		}
		if test.Expected == nil && err == nil {
			continue
		}
		if !strings.Contains(err.Error(), test.Expected.Error()) {
			t.Errorf(`expected "%v", got "%v"\n`, test.Expected, err)
		}
	}
}
