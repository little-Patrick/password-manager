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

func (e *Level1) Encrypt() string {
	return e.passwordShift(true)
}
	// Pin Key
	// square key
	// LATER: Divide key to go down different paths
	// take last 5 digits of sqyared key
	// A key: first two digits (02)
	// B key: second and third digits (27)
	// C key: third and fourth digits (71)
	// D key: fourth and fifth digits (15)

	// password key
	// multiply the ord of each letter together
	// take the last four digits and make a map 
	//

func (e *Level1) Decrypt() string {
	return e.passwordShift(false)
}

func (e *Level1) passwordShift(isEncrypt bool) string {
	if isEncrypt {
		e.finalKey()
	} else {
	}
	return ""
}

func (e *Level1) finalKey() map[string]int  {
	passwordKey, _ := e.passwordKeyMap()
	pinKey, _ := e.keyMap()

	finalKey := map[string]int {
		"A": pinKey["A"] + passwordKey["A"],
		"B": pinKey["B"] / passwordKey["B"],
		"C": pinKey["C"] * passwordKey["C"],
		"D": pinKey["D"] - passwordKey["D"],
	}

	return finalKey
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

func (e *Level1) passwordKeyMap() (map[string]int, error) {
	splitPassword := strings.Split(e.password, "")
	multiplyKey := 1.0
	for i := 0; i < len(splitPassword); i++ {
		if len(splitPassword[i]) > 0 {
			char := int([]rune(splitPassword[i])[0]) 
			key := float64(char)
			multiplyKey *= key
		}
	}
	square := math.Pow(multiplyKey, 2) 
	keyToS := strconv.FormatFloat(square, 'f', -1, 64)
	splitKey := strings.Split(keyToS, "")
	// TODO: Make key different from keyMap
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
