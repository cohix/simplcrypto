## SimplCrypto

This library aims to add a friendlier abstraction to the Go `crypto` package.

Currently, it supports AES-256 GCM symmetric key encryption, RSA-OEAP asymmetric encryption, and RSA PKCS1v15 Digital Signatures. It also includes some helper functions for base64 operations.

Plaintext `[]byte`s are encrypted by a SymKey or KeyPair into a `Message` struct, which includes the ciphertext and some metadata. Message structs can be serialized to JSON or Protobuf.

`[]byte`s can also be signed by a KeyPair that has a private key, creating a `Signature` struct. `Siguature`s can be verified by a KeyPair with a public key.

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