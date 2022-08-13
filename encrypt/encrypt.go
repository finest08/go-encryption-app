package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	mathRand "math/rand"
	"os"
	"time"

	"github.com/finest08/go-encryption-app/utils"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*();{}}?/'_+=-"

type Client struct{}

func NewClient() (c *Client) {
	mathRand.Seed(time.Now().UnixNano())
	c = &Client{}
	return
}

func Generate32Key() {
	fmt.Print("Generating 32bit key... \n***\n")
	passkey := NewClient().GenerateRandomString(32)
	utils.WriteToFile(passkey, "secretKey.txt")
	fmt.Print("\nKey saved to `secretKey.txt`\n\n** STORE KEY IN SECURE PLACE **\n")

	os.Exit(0)
}

func (c *Client) GenerateRandomString(length int) (result string) {
	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[mathRand.Intn(len(letterBytes))]
	}
	result = string(b)
	return
}

func Encrypt(dir, secret string) {
	file := utils.ReadFile(dir)
	key := []byte(secret)
	cphr, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}
	gcm, err := cipher.NewGCM(cphr)
	if err != nil {
		fmt.Println(err)
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
	}
	
	err = os.WriteFile(dir, gcm.Seal(nonce, nonce, file, nil), 0777)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Encryption complete.")
}

func Decrypt(src, secret string) {
	key := []byte(secret)
	cipherData, err := os.ReadFile(src)
	if err != nil {
		fmt.Println(err)
	}
	c, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}
	gcmDecrypt, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err)
	}
	nonceSize := gcmDecrypt.NonceSize()
	if len(cipherData) < nonceSize {
		fmt.Println(err)
	}
	nonce, encryptedMessage := cipherData[:nonceSize], cipherData[nonceSize:]
	plainData, err := gcmDecrypt.Open(nil, nonce, encryptedMessage, nil)
	if err != nil {
		fmt.Println(err)
	}

	err = os.WriteFile(src, plainData, 0777)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Decryption complete.")

}
