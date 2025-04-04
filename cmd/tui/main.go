package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"github.com/rivo/tview"
	"github.com/little-Patrick/password-manager/internal/encrypt"
	"strconv"
)

type PasswordTUI struct {
	Password string `json:"password"`
	Encryption string `json:"encryption"`
	Pin int `json:"pin"`
}

var (
	passwords = []PasswordTUI{}
	passwordsFile = "data/password-manager.json"
)

func main() {
	app := tview.NewApplication()
	loadPasswords()
	passwordList := tview.NewTextView().SetDynamicColors(true).SetWordWrap(true)
	
	passwordList.SetBorder(true).SetTitle("Passwords")

	refreshPassword := func() {
		passwordList.Clear()
		if len(passwords) == 0 {
			fmt.Fprintln(passwordList, "No passwords")
		} else {
			for i, password := range passwords {
				fmt.Fprintf(passwordList, "[%d]\n Password: %s\n Encrypted: %s\n Pin: %d", i+1, password.Password, password.Encryption, password.Pin)
			}
		}
	}

	passwordInput := tview.NewInputField().SetLabel("Password: ")
	pinInput := tview.NewInputField().SetLabel("Pin: ")
	// passwordIdInput := tview.NewInputField().SetLabel("Item Name: ")

	form := tview.NewForm().
		AddFormItem(passwordInput).
		AddFormItem(pinInput).
		// AddFormItem(passwordIdInput).
		AddButton("Encrypt", func () {
			password := passwordInput.GetText()
			pin:= pinInput.GetText()
			
			if password != "" {
				encryption := encrypt.NewLevel1(password, pin)
				encrypted, err := encryption.Encrypt()
				if err != nil {
					fmt.Errorf("Bad Encryption: %v", err)
					return
				}
				convPin, err := strconv.Atoi(pin)
				if err != nil {
					fmt.Errorf("Invalid Pin: %v", err)
					return
				}

				passwords = append(passwords, PasswordTUI{Password: password, Encryption: encrypted, Pin: convPin})

				savePassord()
				refreshPassword()

				passwordInput.SetText("")
				pinInput.SetText("")
				
			}
		})

		form.SetBorder(true).SetTitle("Manage Password").SetTitleAlign(tview.AlignLeft)
	flex := tview.NewFlex().
		AddItem(form, 0, 1, true).
		AddItem(passwordList, 0, 1, false)

	refreshPassword()

	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}


func loadPasswords() {
	_, err := os.Stat(passwordsFile)
	if err == nil {
		data, err := os.ReadFile(passwordsFile)
		if err != nil {
			log.Fatal("Error loading password file:", err)
		}
		json.Unmarshal(data, &passwords)
	}
}

func savePassord() {
	data, err := json.MarshalIndent(passwords, "", " ")
	if err != nil {
		log.Fatal("Error saving password:", err)
	}
	os.WriteFile(passwordsFile, data, 0644)
}

func deletePassword(index int) {
	if index < 0 || index >= len(passwords) {
		fmt.Printf("Invalid password index")
		return
	}
	passwords = append(passwords[:index], passwords[index + 1:]...)
}
