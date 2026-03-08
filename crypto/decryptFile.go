package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"log"

	"golang.org/x/crypto/argon2"
)

func DecryptFile(encrypted []byte, password string) []byte {
	salt := encrypted[:16]
	key := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatal(err)
	}
	nonce := encrypted[16 : 16+gcm.NonceSize()]
	toDecrypt := encrypted[16+gcm.NonceSize():]

	plaintext, err := gcm.Open(nil, nonce, toDecrypt, nil)
	if err != nil {
		log.Fatal(err)
	}
	return plaintext
}
