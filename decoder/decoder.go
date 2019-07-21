package decoder

import (
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"

	"github.com/nodo/certo/renderer"
	"github.com/pkg/errors"
)

type Decoder struct {
	Path   string
	Format string
}

func New(path, format string) Decoder {
	return Decoder{path, format}
}

func (d Decoder) Validate() bool {
	return d.Path != ""
}

func (d Decoder) Decode() (string, error) {
	certBytes, err := ioutil.ReadFile(d.Path)
	if err != nil {
		return "", errors.Wrap(err, "could not read the certificate file")
	}
	block, _ := pem.Decode(certBytes)
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return "", errors.Wrap(err, "could not parse the certificate")
	}
	return renderer.New(cert, d.Format).Render(), nil
}
