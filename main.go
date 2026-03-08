package main

import (
	"FileEncrypt/crypto"
	"FileEncrypt/data"
	"FileEncrypt/utils"
	"fmt"
	"log"
	"os"
)

const FileAlreadyEncryptedError = "File is already encrypted!"
const FileNotEncryptedError = "File not encrypted!"
const EncryptedSuccessfullyMessage = "File encrypted successfully"
const DecryptedSuccessfullyMessage = "File encrypted successfully"

func main() {
	if args := os.Args; len(args) < 4 {
		log.Fatal("Not enough arguments, provide command,filename and password")
	}
	cmd, filename, password := os.Args[1], os.Args[2], os.Args[3]
	switch cmd {
	case "encrypt":
		if utils.IsEncrypted(filename) {
			log.Fatal(FileAlreadyEncryptedError)
		}
		text := data.ReadFile(filename)
		encrypted := crypto.EncryptFile(password, text)
		filenameToSave := utils.TrimExtension(filename) + ".enc"
		err := os.WriteFile(filenameToSave, encrypted, 0o766)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(EncryptedSuccessfullyMessage)
	case "decrypt":
		if !utils.IsEncrypted(filename) {
			log.Fatal(FileNotEncryptedError)
		}
		text := data.ReadFile(filename)
		decrypted := crypto.DecryptFile(text, password)
		filenameToSave := utils.TrimExtension(filename) + "_dec.txt"
		err := os.WriteFile(filenameToSave, decrypted, 0o766)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Decrypted: %s", decrypted)

	default:
		log.Fatal("No such command")

	}

}
