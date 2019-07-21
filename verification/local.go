package verification

import (
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"

	"github.com/pkg/errors"
)

type ValidateVerifier interface {
	Validate() bool
	Verify() (bool, error)
}

type Local struct {
	CertPath   string
	CACertPath string
}

func NewLocal(certPath, caCertPath string) ValidateVerifier {
	return Local{
		CertPath:   certPath,
		CACertPath: caCertPath,
	}
}

func (l Local) Validate() bool {
	return l.CertPath != "" && l.CACertPath != ""
}

func (l Local) Verify() (bool, error) {
	certPEM, err := ioutil.ReadFile(l.CertPath)
	if err != nil {
		return false, errors.Wrap(err, "failed to read the certificate")
	}
	rootPEM, err := ioutil.ReadFile(l.CACertPath)
	if err != nil {
		return false, errors.Wrap(err, "failed to read the CA certificate")
	}

	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM(rootPEM)
	if !ok {
		return false, errors.New("failed to parse root certificate")
	}

	block, _ := pem.Decode([]byte(certPEM))
	if block == nil {
		return false, errors.New("failed to parse certificate PEM")
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return false, errors.Wrap(err, "failed to parse certificate")
	}

	opts := x509.VerifyOptions{
		Roots: roots,
	}

	if _, err := cert.Verify(opts); err != nil {
		return false, errors.Wrap(err, "failed to verify certificate")
	}
	return true, nil
}
