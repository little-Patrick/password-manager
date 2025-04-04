 <h1 align="center">🔐 password-manager</h1>
<p align="center">
  A terminal-based password manager built with Go and <a href="https://github.com/rivo/tview">tview</a>.
</p>

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.8+-00ADD8?style=for-the-badge&logo=go" alt="Go Version">
  <img src="https://img.shields.io/badge/Terminal%20UI-tview-5F4B8B?style=for-the-badge">
  <img src="https://img.shields.io/badge/CLI%20App-✔️-brightgreen?style=for-the-badge">
</p>

---

<h2 align="center"> ✨ Overview</h2>

<p align="center">This TUI was created as a password manager and password creator. It allows you to create multiple passwords from the same root password by just changing the pin making it easier to remember you passwords.</p>

<h2 align="center">❗️Disclaimer❗️</h2>

<p align="center">This application does not encrypt the local json file. It will display your password in plain text if someone gets access to it.</p>

<h2 align="center">Prerequisites</h2>

<p align="center">- Go 1.18+</p>

<h2 align="center">Insallation</h2>
<p>Clone the repo</p>

```bash
git clone https://github.com/yourusername/project-name.git
```
<h4>Option 1</h4>

<p>Run the program</p>

```bash
go run cmd/tui/main.go
```
<h4>Option 2</h4>

<p>Build the program as an executable</p>

```bash
go build -o password cmd/tui/main.go
```

<p>Run:</p>

```bash
./password
```

<h2 align="center"> Usage</h2>

<p>In the form on the left hand side input a password and 5 digit pin, then press enter. On the right hand side your original password, new password, and pin will appear!</p>

<h2 align="center"> Keybindings/Navigation</h2>

- `Tab:` Move to next form
- `Shift + Tab:` Move to previous form
- `Enter:` Will move to the next form. If the Encrypt button is selected it will create the password
- `Ctrl + c:` Exit the program

<h2 align="center"> 🛠️ TODO / Roadmap</h2>
  
- [ ] Add encryption/decryption for local json file
- [ ] Add password for entering the TUI interface
- [ ] Add decryption support
- [ ] Add password ID field and search
- [ ] Add UI option to delete a saved password
- [ ] Improve encryption levels
  - [ ] 2 additional handrolled encryptions
  - [ ] AES for encrypting json files
- [ ] Have the download be an executable to a directory in PATH
- [ ] Allow for mouse click to select fields



