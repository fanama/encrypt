package encrypt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
)

type RSA struct {
	privateKey rsa.PrivateKey
}

func BuildRSA() (RSA, error) {

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)

	if err != nil {
		return RSA{}, nil
	}

	return RSA{privateKey: *privateKey}, err
}

func (this *RSA) Encrypt(secretMessage string) (string, error) {
	key := this.privateKey.PublicKey
	label := []byte("OAEP Encrypted")
	rng := rand.Reader
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rng, &key, []byte(secretMessage), label)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func (this *RSA) Decrypt(cipherText string) (string, error) {
	ct, _ := base64.StdEncoding.DecodeString(cipherText)
	label := []byte("OAEP Encrypted")
	rng := rand.Reader
	plaintext, err := rsa.DecryptOAEP(sha256.New(), rng, &this.privateKey, ct, label)

	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
