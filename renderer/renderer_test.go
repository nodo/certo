package renderer_test

import (
	"testing"
	"time"

	"github.com/nodo/certo/renderer"
)

var prettyCert = renderer.PrettyCert{
	IssuerCommonName:  "issuer",
	SubjectCommonName: "subject",
	NotBefore:         time.Date(2014, 9, 12, 8, 48, 42, 0, time.UTC),
	NotAfter:          time.Date(2024, 9, 9, 8, 48, 42, 0, time.UTC),
	DNSNames:          []string{"dns"},
}

func TestTextRenderer(t *testing.T) {
	r := renderer.TextRenderer{prettyCert}
	out := r.Render()
	expected := `Issuer              issuer
Subject             subject
Not Before          2014-09-12 08:48:42 +0000 UTC
Not After           2024-09-09 08:48:42 +0000 UTC
`
	if out != expected {
		t.Errorf("\nexpected:\n%s\n\ngot:\n\n%s", expected, out)
	}
}

func TestJSONRenderer(t *testing.T) {
	r := renderer.JSONRenderer{prettyCert}
	out := r.Render()
	expected := `{"issuer_common_name":"issuer","subject_common_name":"subject","not_before":"2014-09-12T08:48:42Z","not_after":"2024-09-09T08:48:42Z","DNSNames":["dns"]}`

	if out != expected {
		t.Errorf("\nexpected:\n%s\n\ngot:\n\n%s", expected, out)
	}
}

func TestYAMLRenderer(t *testing.T) {
	r := renderer.YAMLRenderer{prettyCert}
	out := r.Render()
	expected := `issuer_common_name: issuer
subject_common_name: subject
not_before: 2014-09-12T08:48:42Z
not_after: 2024-09-09T08:48:42Z
dnsnames:
- dns
`

	if out != expected {
		t.Errorf("\nexpected:\n%s\n\ngot:\n\n%s", expected, out)
	}
}
