package nsign

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
)

type RSAMethod struct {
	h          crypto.Hash
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

func NewRSAMethod(h crypto.Hash, privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) *RSAMethod {
	var nRSA = &RSAMethod{}
	nRSA.h = h
	nRSA.privateKey = privateKey
	nRSA.publicKey = publicKey
	return nRSA
}

func (method *RSAMethod) Sign(data []byte) ([]byte, error) {
	var h = method.h.New()
	if _, err := h.Write(data); err != nil {
		return nil, err
	}
	var hashed = h.Sum(nil)
	return rsa.SignPKCS1v15(rand.Reader, method.privateKey, method.h, hashed)
}

func (method *RSAMethod) Verify(data []byte, signature []byte) error {
	var h = method.h.New()
	if _, err := h.Write(data); err != nil {
		return err
	}
	var hashed = h.Sum(nil)
	return rsa.VerifyPKCS1v15(method.publicKey, method.h, hashed, signature)
}
