package ssl

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"

	"github.com/juju/errors"
)

//NewConfig create a tls config
func NewConfig(caPEM string, certPEM string, keyPEM string) (*tls.Config, error) {
	// ClientCAs defines the set of root certificate authorities
	// that servers use if required to verify a client certificate
	// by the policy in ClientAuth.
	pool := x509.NewCertPool()
	crt, err := ioutil.ReadFile(caPEM)
	if err != nil {
		return nil, errors.Trace(err)
	}
	pool.AppendCertsFromPEM(crt)

	//use key.pem and cert.pem to create a certificate
	cert, err := tls.LoadX509KeyPair(certPEM, keyPEM)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return &tls.Config{
		RootCAs:            pool,
		Certificates:       []tls.Certificate{cert},
		//TLS accepts any certificate
        //presented by the server and any host name in that certificate
		InsecureSkipVerify: true,
	}, nil

}
