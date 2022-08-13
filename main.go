package main

import (
	"github.com/finest08/go-encryption-app/process"
	"github.com/finest08/go-encryption-app/utils"
	// "github.com/finest08/go-encryption-app/encrypt"
)

func main() {

	// encrypt.Generate32Key()

	secret := string(utils.ReadFile("secretKey.txt"))

	process.EncryptDir("/Users/mitchwilson/go/src/github.com/finest08/go-encryption-app/playground/dir", secret)
	// process.DecryptDir("/Users/mitchwilson/go/src/github.com/finest08/go-encryption-app/playground/dir&mwx&&.zip", secret)

}