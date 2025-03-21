package main

import (
	"fmt"
	"github.com/little-Patrick/password-manager/internal/encrypt"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	e := encrypt.NewLevel1("abc", "")

	encrypted := e.Encrypt()
	fmt.Println("Encrypted:", encrypted)

	decrypted := e.Decrypt()
	fmt.Println("Decrypted:", decrypted)

	if err := app.Run(); err != nil {
	    panic(err)
	}
}
