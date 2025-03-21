package encrypt

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

type Level2 struct {
	password string
	key1 string
	key2 string
}

func NewLevel2(password string, key1 string, key2 string) *Level2 {
	if password == "" {
		// TODO: Thow an error
	}
	if key1 == "" {
				key1 = fmt.Sprintf("%d", rand.Intn(90000)+1000)

	}
	if key2 == "" {
				key2 = fmt.Sprintf("%d", rand.Intn(90000)+1000)
	}
	return &Level2{password: password, key1: key1, key2: key2}
}
