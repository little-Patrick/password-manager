package encrypt

import (
	"runtime"
	"fmt"
	"math/rand"
	"math"
	"strconv"
	"strings"
	"errors"
)

type Level1 struct {
	password string
	key1     string
	key2 	 string
}

func NewLevel1(password string, key1 string, key2 string) *Level1 {
	if password == "" {
		// TODO: Handle error properly (e.g., return an error instead)
		panic("Password cannot be empty")
	}
	if key1 == "" {
		key1 = fmt.Sprintf("%d", rand.Intn(90000)+1000)
	}
	if key2 == "" {
		key2 = fmt.Sprintf("%d", rand.Intn(90000)+1000)
	}
	return &Level1{password: password, key1: key1, key2: key2}
}

func (e *Level1) Encrypt() (string, error) {
	return e.passwordShift(true)
}

func (e *Level1) Decrypt() (string, error) {
	return e.passwordShift(false)
}

func (e *Level1) passwordShift(isEncrypt bool) (string, error) {
	if isEncrypt {
		key, err := e.keyMap()
		if err != nil {
			return "", err
		}
	} else {
	}
	return "", nil
}

func (e *Level1) keyMap() (map[string]int, error) {
	splitKey := strings.Split(e.key1, "")
	if len(splitKey) < 5 {
		return nil, errors.New("Can't have a pin less than 5")
	}
	mapKey := map[string]string {
		"A": splitKey[0] + splitKey[1],
		"B": splitKey[1] + splitKey[2],
		"C": splitKey[2] + splitKey[3],
		"D": splitKey[3] + splitKey[4],
	}
	// TODO: Make helper function ID: 1
	intKeyMap := make(map[string]int)
	for key, value := range mapKey {
		intValue, err := strconv.Atoi(value)
		if err != nil {
			return nil, fmt.Errorf("Something went wrong %s: %v", key, err)
		}
		intKeyMap[key] = intValue
	}
	runtime.Breakpoint()
	
	return intKeyMap, nil
}

func (e *Level1) letters() []rune {
	var letters []rune
	for i := 'a'; i <= 'z'; i++ {
		letters = append(letters, i)
	}
	for i := 'A'; i <= 'Z'; i++ {
		letters = append(letters, i)
	}
	for i := '0'; i <= '9'; i++ {
		letters = append(letters, i)
	}
	specialChars := []rune{'`', '~', '!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '-', '_', '=', '+', '[', '{', ']', '}', '|', ';', ',', '<', '.', '>', '/', '?'}
	letters = append(letters, specialChars...)
	return letters
}
