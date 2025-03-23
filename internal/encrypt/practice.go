package encrypt
//
// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// 	"strconv"
// 	"strings"
// )
//
// type Enigma struct {
// 	message string
// 	key     string
// 	date    string
// }
//
// // NewEnigma creates a new instance of Enigma with the message, key, and date.
// func NewEnigma(message string, key string, date string) *Enigma {
// 	if key == "" {
// 		key = fmt.Sprintf("%d", rand.Intn(90000)+1000) // random key between 1000 and 99999
// 	}
// 	if date == "" {
// 		date = time.Now().Format("02-01-06")
// 	}
// 	return &Enigma{message: message, key: key, date: date}
// }
//
// // Encrypt encrypts the message.
// func (e *Enigma) Encrypt() string {
// 	return e.messageShift(true)
// }
//
// // Decrypt decrypts the message.
// func (e *Enigma) Decrypt() string {
// 	return e.messageShift(false)
// }
//
// // shiftKey generates the shift key based on the key.
// func (e *Enigma) shiftKey() map[string]string {
// 	keySplit := strings.Split(e.key, "")
// 	return map[string]string{
// 		"A": keySplit[0] + keySplit[1],
// 		"B": keySplit[1] + keySplit[2],
// 		"C": keySplit[2] + keySplit[3],
// 		"D": keySplit[3] + keySplit[4],
// 	}
// }
//
// // shiftDate generates the date-based shift.
// func (e *Enigma) shiftDate() map[string]string {
// 	square := e.squareDate()
// 	dateSplit := strings.Split(square, "")
// 	return map[string]string{
// 		"A": string(dateSplit[len(dateSplit)-4]),
// 		"B": string(dateSplit[len(dateSplit)-3]),
// 		"C": string(dateSplit[len(dateSplit)-2]),
// 		"D": string(dateSplit[len(dateSplit)-1]),
// 	}
// }
//
// // squareDate squares the date as an integer and returns it as a string.
// func (e *Enigma) squareDate() string {
// 	dateInt, _ := strconv.Atoi(e.date)
// 	return fmt.Sprintf("%d", dateInt*dateInt)
// }
//
// // finalKey calculates the final key by adding the shift key and date-based shift.
// func (e *Enigma) finalKey() map[string]int {
// 	final := make(map[string]int)
// 	shiftKey := e.shiftKey()
// 	shiftDate := e.shiftDate()
// 	for k, v := range shiftKey {
// 		keyInt, _ := strconv.Atoi(v)
// 		dateInt, _ := strconv.Atoi(shiftDate[k])
// 		final[k] = keyInt + dateInt
// 	}
// 	return final
// }
//
// // messageShift handles both encryption and decryption based on the shift.
// func (e *Enigma) messageShift(isEncrypt bool) string {
// 	letterKey := e.letters()
// 	final := e.finalKey()
// 	offset := e.offset(final, isEncrypt)
//
// 	var result []string
// 	for i := 0; i < len(e.message); i += 4 {
// 		slice := []rune(e.message[i:min(i+4, len(e.message))])
// 		if len(slice) > 0 {
// 			result = append(result, string(offset["A"][slice[0]]))
// 		}
// 		if len(slice) > 1 {
// 			result = append(result, string(offset["B"][slice[1]]))
// 		}
// 		if len(slice) > 2 {
// 			result = append(result, string(offset["C"][slice[2]]))
// 		}
// 		if len(slice) > 3 {
// 			result = append(result, string(offset["D"][slice[3]]))
// 		}
// 	}
// 	return strings.Join(result, "")
// }
//
// // offset calculates the letter shifts based on the final key and whether we're encrypting or decrypting.
// func (e *Enigma) offset(final map[string]int, isEncrypt bool) map[string]map[rune]rune {
// 	letterKey := e.letters()
// 	offset := make(map[string]map[rune]rune)
//
// 	for _, k := range []string{"A", "B", "C", "D"} {
// 		shift := final[k]
// 		if !isEncrypt {
// 			shift = -shift
// 		}
// 		offset[k] = make(map[rune]rune)
// 		for _, letter := range letterKey {
// 			offset[k][letter] = letterKey[(int(letter)-shift)%len(letterKey)]
// 		}
// 	}
//
// 	return offset
// }
//
// // letters returns the set of characters (lowercase, uppercase, digits, special chars).
// func (e *Enigma) letters() []rune {
// 	var letters []rune
// 	for i := 'a'; i <= 'z'; i++ {
// 		letters = append(letters, i)
// 	}
// 	for i := 'A'; i <= 'Z'; i++ {
// 		letters = append(letters, i)
// 	}
// 	for i := '0'; i <= '9'; i++ {
// 		letters = append(letters, i)
// 	}
// 	specialChars := []rune{'`', '~', '!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '-', '_', '=', '+', '[', '{', ']', '}', '|', ';', ',', '<', '.', '>', '/', '?'}
// 	letters = append(letters, specialChars...)
// 	return letters
// }
//
// // min returns the smaller of two integers.
// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }
