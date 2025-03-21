package encrypt

import (
	"fmt"
	"math/rand"
	// "strconv"
	// "strings"
)

type Level1 struct {
	password string
	key1 string
	key2 string
}

func NewLevel1(password string, key1 string) *Level1 {
	if password == "" {
		// TODO: Thow an error
	}
	if key1 == "" {
				key1 = fmt.Sprintf("%d", rand.Intn(90000)+1000)

	}
	return &Level1{password: password, key1: key1}
}
