package verification

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/url"

	"github.com/pkg/errors"
)

type Remote struct {
	URL        string
	CACertPath string
}

func NewRemote(url, caCertPath string) ValidateVerifier {
	return Remote{
		URL:        url,
		CACertPath: caCertPath,
	}
}

func (r Remote) Validate() bool {
	return r.URL != "" && r.CACertPath != ""
}

func parseURL(rawurl string) (string, error) {
	u, err := url.Parse(rawurl)
	if err != nil {
		return "", err
	}
	if u.Scheme == "" {
		return "", errors.New("Unknown protocol")
	}

	port := "443"
	if u.Port() != "" {
		port = u.Port()
	}

	return fmt.Sprintf("%s:%s", u.Hostname(), port), nil
}

func (r Remote) Verify() (bool, error) {
	rootPEM, err := ioutil.ReadFile(r.CACertPath)
	if err != nil {
		return false, errors.Wrap(err, "failed to read the CA certificate")
	}

	parsedURL, err := parseURL(r.URL)
	if err != nil {
		return false, errors.Wrap(err, "invalid URL")
	}

	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM(rootPEM)
	if !ok {
		return false, errors.New("failed to parse root certificate")
	}

	conn, err := tls.Dial("tcp", parsedURL, &tls.Config{
		RootCAs: roots,
	})
	if err != nil {
		return false, errors.Wrap(err, "failed to connect")
	}
	conn.Close()

	return true, nil
}
