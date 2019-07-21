package actions

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"

	"github.com/pkg/errors"
)

func CheckURL(url, caCertPath string) (bool, error) {
	rootPEM, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		return false, errors.Wrap(err, "failed to read the CA certificate")
	}

	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM(rootPEM)
	if !ok {
		return false, errors.New("failed to parse root certificate")
	}

	conn, err := tls.Dial("tcp", url, &tls.Config{
		RootCAs: roots,
	})
	if err != nil {
		return false, errors.Wrap(err, "failed to connect")
	}
	conn.Close()

	return true, nil
}
