package simplcrypto

import (
	"crypto/ecdsa"
	"errors"
	"math/big"

	"crypto/elliptic"

	jose "gopkg.in/square/go-jose.v2"
	jcrypto "gopkg.in/square/go-jose.v2/crypto"
)

// ECDSA describes a signing key
type ECDSA struct {
	KID string `json:"kid"`
	Ext bool   `json:"ext"`
	Kty string `json:"kty"`
	D   string `json:"d"`
	Pub *ECDSAPub
}

// ECDSAPub describes a verifying key
type ECDSAPub struct {
	X string `json:"x"`
	Y string `json:"y"`
}

// Sign creates a signature from data
func (e *ECDSA) Sign(data []byte) (*Signature, error) {
	if e.Pub == nil {
		return nil, errors.New("missing ecdsa pubkey")
	}

	rawX, _ := Base64URLDecode(e.Pub.X)
	bigX := &big.Int{}
	bigX.SetBytes(rawX)

	rawY, _ := Base64URLDecode(e.Pub.Y)
	bigY := &big.Int{}
	bigX.SetBytes(rawY)

	rawD, _ := Base64URLDecode(e.D)
	bigD := &big.Int{}
	bigX.SetBytes(rawD)

	key := ecdsa.PrivateKey{
		PublicKey: ecdsa.PublicKey{
			X:     bigX,
			Y:     bigY,
			Curve: elliptic.P256(),
		},
		D: bigD,
	}

	sigKey := jose.SigningKey{
		Algorithm: "ES256",
		Key:       &key,
	}

	signer, err := jose.NewSigner(sigKey, nil)
	if err != nil {
		return nil, err
	}

	sig, err := signer.Sign(data)
	if err != nil {
		return nil, err
	}

	sigBytes := sig.Signatures[0].Signature

	outSig := &Signature{
		KID:       e.KID,
		Signature: sigBytes,
	}

	return outSig, nil
}

func (e *ECDSAPub) Verify(data []byte, sig *Signature) {
	method := jcrypto.SigningMethodES256
}
