## SimplCrypto

This library aims to add a friendlier abstraction to the Go `crypto` package.

Currently, it supports AES-256 GCM symmetric key encryption, RSA-OEAP asymmetric encryption, and RSA PKCS1v15 Digital Signatures. It also includes some helper functions for base64 operations.

`KeyPair`s can contain a public key portion, or the public and private portions of the RSA key.

Plaintext `[]byte`s are encrypted by a `SymKey` or `KeyPair` into a `Message` struct, which includes the ciphertext and some metadata. Message structs can be serialized to JSON or Protobuf.

`[]byte`s can also be signed by a `KeyPair` that has a private key portion, creating a `Signature` struct. `Siguature`s can be verified by a `KeyPair` so long as they contain the public key portion. Signatures can be serialized to JSON or Protobuf.

`KeyPair` public keys can be converted to `SerializablePubKey` in order to be serialized into JSON or Protobuf. `SymKeys` can be serialized to JSON as-is.

### Install

To install, use `go get`, `gvt` or `dep`:
```
go get github.com/cohix/simplcrypto
```
```
gvt fetch github.com/cohix/simplcrypto
```
```
dep ensure --add github.com/cohix/simplcrypto
```

### Examples

Asymmetric encryption example:
```Go
keypair, err := simplcrypto.GenerateNewKeyPair()
if err != nil {
	log.Fatal(err)
}

plain := []byte("very secret, do not share")

msg, err := keypair.Encrypt(plain)
if err != nil {
	log.Fatal(err)
}

newPlain, err := keypair.Decrypt(msg)
if err != nil {
	log.Fatal(err)
}

fmt.Println(string(newPlain))
```

Symmetric example:
```Go
symkey, err := simplcrypto.GenerateSymKey()
if err != nil {
	log.Fatal(err)
}

plain := []byte("still pretty secret, tho!")

msg, err := symkey.Encrypt(plain)
if err != nil {
	log.Fatal(err)
}

newPlain, err := symkey.Decrypt(msg)
if err != nil {
	log.Fatal(err)
}

fmt.Println(string(newPlain))
```

Signature example:
```Go
keypair, err := simplcrypto.GenerateNewKeyPair()
if err != nil {
	log.Fatal(err)
}

data := []byte("very secret, do not share")

signature, err := keypair.Sign(data)
if err != nil {
	log.Fatal(err)
}

if err := keypair.Verify(data, signature); err != nil {
	log.Fatal(err)
}

fmt.Println("verified!")
```

This library is provided free of charge for use in any project, personal or professional. The maintainers of this library shall assume no liability for its use under any circumstance. By using this library, you agree to these terms.