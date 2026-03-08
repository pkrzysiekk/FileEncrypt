package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"log"

	"golang.org/x/crypto/argon2"
)

func EncryptFile(password string, fileToEncrypt []byte) []byte {
	salt := make([]byte, 16)
	rand.Read(salt)
	key := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatal(gcm)
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = rand.Read(nonce)
	if err != nil {
		log.Fatal(err)
	}
	encrypted := gcm.Seal(nil, nonce, fileToEncrypt, nil)
	dataToSave := append(salt, nonce...)
	dataToSave = append(dataToSave, encrypted...)
	return dataToSave
}
