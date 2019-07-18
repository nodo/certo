package renderer

import (
	"crypto/x509"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	yaml "gopkg.in/yaml.v2"
)

type PrettyCert struct {
	IssuerCommonName  string    `json:"issuer_common_name" yaml:"issuer_common_name"`
	SubjectCommonName string    `json:"subject_common_name" yaml:"subject_common_name"`
	NotBefore         time.Time `json:"not_before" yaml:"not_before"`
	NotAfter          time.Time `json:"not_after" yaml:"not_after"`
	DNSNames          []string  `json:",omitempty" yaml:",omitempty"`
}

type Renderer interface {
	Render() string
}

func New(Cert *x509.Certificate, format string) (r Renderer) {
	pc := PrettyCert{
		IssuerCommonName:  Cert.Issuer.CommonName,
		SubjectCommonName: Cert.Subject.CommonName,
		NotBefore:         Cert.NotBefore,
		NotAfter:          Cert.NotAfter,
		DNSNames:          Cert.DNSNames,
	}
	switch format {
	case "json":
		return JSONRenderer{Cert: pc}
	case "yaml":
		return YAMLRenderer{Cert: pc}
	}
	return TextRenderer{Cert: pc}
}

type TextRenderer struct {
	Cert PrettyCert
}

func (r TextRenderer) Render() string {
	var str strings.Builder
	str.WriteString(fmt.Sprintf("%-20s%s\n", "Issuer", r.Cert.IssuerCommonName))
	str.WriteString(fmt.Sprintf("%-20s%s\n", "Subject", r.Cert.SubjectCommonName))
	str.WriteString(fmt.Sprintf("%-20s%s\n", "Not Before", r.Cert.NotBefore))
	str.WriteString(fmt.Sprintf("%-20s%s\n", "Not After", r.Cert.NotAfter))
	return str.String()
}

type JSONRenderer struct {
	Cert PrettyCert
}

func (r JSONRenderer) Render() string {
	CertBytes, err := json.Marshal(r.Cert)
	if err != nil {
		panic(err)
	}
	return string(CertBytes)
}

type YAMLRenderer struct {
	Cert PrettyCert
}

func (r YAMLRenderer) Render() string {
	CertBytes, err := yaml.Marshal(r.Cert)
	if err != nil {
		panic(err)
	}
	return string(CertBytes)
}
