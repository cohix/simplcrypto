package simplcrypto

import (
	"crypto/hmac"
	"crypto/sha256"
)

// HMACWithSecretAndData HMAC's a string using a secret string
func HMACWithSecretAndData(secret, data string) []byte {
	hmaccer := hmac.New(sha256.New, []byte(secret))
	hmaccer.Write([]byte(data))

	return hmaccer.Sum(nil)
}
