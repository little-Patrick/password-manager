package encrypt

import (
	"fmt"
	"math/rand"
)

type Level1 struct {
	password string
	key1     string
	key2     string
}

// NewLevel1 is a constructor function for Level1
func NewLevel1(password string, key1 string) *Level1 {
	if password == "" {
		// TODO: Handle error properly (e.g., return an error instead)
		panic("Password cannot be empty")
	}
	if key1 == "" {
		key1 = fmt.Sprintf("%d", rand.Intn(90000)+1000) // Fix formatting
	}
	return &Level1{password: password, key1: key1}
}

// Encrypt (Placeholder function)
func (l *Level1) Encrypt() string {
	return "Encrypted: " + l.password + l.key1
	// 
}

// Decrypt (Placeholder function)
func (l *Level1) Decrypt() string {
	return "Decrypted: " + l.password
}

