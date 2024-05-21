//go:build go1.20.0

package ncrypto

import (
	"crypto/ecdh"
	"errors"
)

func (p PKIXPublicKey) ECDHPublicKey() (*ecdh.PublicKey, error) {
	if p.err != nil {
		return nil, p.err
	}
	publicKey, ok := p.key.(*ecdh.PublicKey)
	if !ok {
		return nil, errors.New("key is not a valid *ecdh.PublicKey")
	}
	return publicKey, nil
}

func (p PKCS8PrivateKey) ECDHPrivateKey() (*ecdh.PrivateKey, error) {
	if p.err != nil {
		return nil, p.err
	}
	privateKey, ok := p.key.(*ecdh.PrivateKey)
	if !ok {
		return nil, errors.New("key is not a valid *ecdh.PrivateKey")
	}
	return privateKey, nil
}
