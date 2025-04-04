package encrypt

import (
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
	// if password == "" {
	// 	// TODO: Handle error properly (e.g., return an error instead)
	// }
	if key1 == "" {
		key1 = strconv.Itoa(rand.Intn(100000-10000))
	}
	return &Level1{password: password, key1: key1}
}

func (e *Level1) Encrypt() (string, error) {
	encrypt, err := e.encryptOffset()
	if err != nil {
		return "", fmt.Errorf("Error: key: %s, err: %v", encrypt, err)

	}
	 
	return encrypt, nil
}

func (e *Level1) Decrypt() (string, error) {
	return "", nil
}

func (e *Level1) encryptOffset() (string, error) {
	finalKey, err  := e.finalKey()
	if err != nil {
		return "", fmt.Errorf("Error: key: %v, err: %v", finalKey, err)
	}
	password := e.passwordSplit()

	encryptedPassword := ""
	for i := 0; i < len(password); i++  {
		for j := 0; j < len(password[i]); {
			if len(password[i]) >= 1 {
				letterA := finalKey["A"][password[i][j]]
				encryptedPassword = encryptedPassword + letterA
				j++
			}
			if len(password[i]) >= 2 {
				letterB := finalKey["B"][password[i][j]]
				encryptedPassword = encryptedPassword + letterB
				j++
			}
			if len(password[i]) >= 3 {
				letterC := finalKey["C"][password[i][j]]
				encryptedPassword = encryptedPassword + letterC
				j++
			}
			if len(password[i]) >= 4 {
				letterD:= finalKey["D"][password[i][j]] 
				encryptedPassword = encryptedPassword + letterD
				j++
				
			}
		}
	}
	            
	return encryptedPassword, nil
}

func (e *Level1) finalKey() (map[string]map[string]string, error) {
	letters := e.letters()
	keyMap, err := e.keyMap()
	if err != nil {
		return nil, fmt.Errorf("Error: key: %v, err: %v", keyMap, err)
	}

	var aKey []string
	var bKey []string
	var cKey []string
	var dKey []string
	for index, _ := range letters {
		indexA := (index + keyMap["A"] -1) % len(letters)
		indexB := (index + keyMap["B"] -1) % len(letters)
		indexC := (index + keyMap["C"] -1) % len(letters)
		indexD := (index + keyMap["D"] -1) % len(letters)
		 
		aKey = append(aKey, letters[indexA])
		bKey = append(bKey, letters[indexB])
		cKey = append(cKey, letters[indexC])
		dKey = append(dKey, letters[indexD])
	}

	var aMap = make(map[string]string)
	var bMap = make(map[string]string)
	var cMap = make(map[string]string)
	var dMap = make(map[string]string)
	for i := 0; i < len(letters); i++ {
		aMap[letters[i]] = aKey[i]
		bMap[letters[i]] = bKey[i]
		cMap[letters[i]] = cKey[i]
		dMap[letters[i]] = dKey[i]
	}

	finalKey := make(map[string]map[string]string)
	finalKey["A"] = aMap
	finalKey["B"] = bMap
	finalKey["C"] = cMap
	finalKey["D"] = dMap

	return finalKey, nil
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

func (e *Level1) letters() []string {
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
	specialChars := []rune{'!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '-', '_', '=', '+', '[', '{', ']', '}', '|', ';', ',', '<', '.', '>', '/', '?'}
	letters = append(letters, specialChars...)

	var letterKey []string 

	for _, letter := range letters {
		letterKey = append(letterKey, string(letter))
	}

	return letterKey
}
