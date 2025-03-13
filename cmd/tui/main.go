package main

import (
	"fmt"
	"github.com/little-Patrick/password-manager"
)

func main() {
	// Create an Enigma instance, passing in the message and any other parameters.
	e := encryption.NewEnigma("hello", "", "")
	
	// Encrypt the message
	encrypted := e.Encrypt()
	fmt.Println("Encrypted:", encrypted)

	// Decrypt the message
	decrypted := e.Decrypt()
	fmt.Println("Decrypted:", decrypted)
}

