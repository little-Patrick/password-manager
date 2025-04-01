package encrypt

import (
	"runtime"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"errors"
)

type Level1 struct {
	password string
	key1     string
}

func NewLevel1(password string, key1 string) *Level1 {
	if password == "" {
		// TODO: Handle error properly (e.g., return an error instead)
		panic("Password cannot be empty")
	}
	if key1 == "" {
		key1 = fmt.Sprintf("%d", rand.Intn(90000)+1000)
	}
	return &Level1{password: password, key1: key1}
}

func (e *Level1) Encrypt() (string, error) {
	return e.passwordShift(true)
}

func (e *Level1) Decrypt() (string, error) {
	return e.passwordShift(false)
}

func (e *Level1) passwordShift(isEncrypt bool) (string, error) {
	if isEncrypt {
		e.offset()
		// key, err := e.keyMap()
		// if err != nil {
		// 	return "", err
		// }

		// split the password into array each 4 index's
		// offset first 4 index with A key repeat for rest of keys
		// 
		// 
		
	} else {
	}
	return "", nil
}

func (e *Level1) offset() (string, error) {
	split := e.passwordSplit()
	key, err := e.keyMap()
	if err != nil {
		return "", fmt.Errorf("Error: password: %s, key: %v, err: %v", split, key, err)
	}
	runtime.Breakpoint()
	return "", nil
}

func (e *Level1) keyShift() map[rune]rune {

}

func (e *Level1) passwordSplit() [][]string {
	splitPassword := strings.Split(e.password, "")

	var offsetPassword [][]string
	for i := 0; i < len(splitPassword); i += 4 {
		sliceRange := i + 4
		if sliceRange >= len(splitPassword) {
			sliceRange = (i + (len(splitPassword) % 4))
		}
		slice := splitPassword[i:sliceRange]
	    offsetPassword = append(offsetPassword, slice)
	}
	return offsetPassword

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
	intKeyMap := make(map[string]int)
	for key, value := range mapKey {
		intValue, err := strconv.Atoi(value)
		if err != nil {
			return nil, fmt.Errorf("Something went wrong %s: %v", key, err)
		}
		intKeyMap[key] = intValue
	}
	return intKeyMap, nil
}

func (e *Level1) letters() map[rune]rune {
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

	letterKey := map[rune]rune {}

	for _, letter := range letters {
		letterKey[letter] = letter
	}

	return letterKey
}
