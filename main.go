package main

import (
	"fanama/encrypt/infra/service/encrypt"
	"fmt"
)

func main() {

	crypt, err := encrypt.BuildRSA()

	if err != nil {
		fmt.Println("err", err)
		return
	}

	secretMessage := "This is super secret message!"

	encryptedMessage, err := crypt.Encrypt(secretMessage)

	if err != nil {
		fmt.Println("err", err)
		return
	}

	fmt.Println("Cipher Text:", encryptedMessage)

	result, err := crypt.Decrypt(encryptedMessage)

	if err != nil {
		fmt.Println("err", err)
		return
	}

	fmt.Println("RESULT", result)
}
