package ncrypto

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
)

func DecodeCertificate(data []byte) (*x509.Certificate, error) {
	der, err := decodeCertificate(data)
	if err != nil {
		return nil, err
	}
	cert, err := x509.ParseCertificate(der)
	if err != nil {
		return nil, err
	}
	return cert, nil
}

func decodeCertificate(data []byte) ([]byte, error) {
	if len(data) == 0 {
		return nil, errors.New("invalid certificate")
	}

	if data[0] == '-' {
		block, _ := pem.Decode(data)
		if block == nil {
			return nil, errors.New("invalid certificate")
		}
		return block.Bytes, nil
	}
	return base64decode(data)
}
