package process

import (
	"fmt"

	"github.com/finest08/go-encryption-app/encrypt"
	"github.com/finest08/go-encryption-app/zip"
)

func EncryptDir(src, secret string) {
	ns := zip.Compress(src)
	srcE := fmt.Sprintf("%s.zip", ns)
	encrypt.Encrypt(srcE, secret)
}

func DecryptDir(src, secret string) {

	encrypt.Decrypt(src, secret)
	zip.Decompress(src)
}
