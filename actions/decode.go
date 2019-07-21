package actions

import (
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"

	"github.com/nodo/certo/renderer"
	"github.com/pkg/errors"
)

func Decode(path, format string) (string, error) {
	certBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return "", errors.Wrap(err, "could not read the certificate file")
	}
	block, _ := pem.Decode(certBytes)
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return "", errors.Wrap(err, "could not parse the certificate")
	}
	return renderer.New(cert, format).Render(), nil
}
