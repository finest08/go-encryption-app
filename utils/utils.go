package utils

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	cp "github.com/otiai10/copy"
)

func CopyDir(src string) string {
	// newCopy := fmt.Sprintf("%s%v", "playground/cp_", src)
	ns := strings.Replace(src, ".zip", "", 1)
	err := cp.Copy(src, ns)
	if err != nil {
		fmt.Println(err)
	}
	return src
}

func ReadFile(filename string) (content []byte) {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err.Error())
	}
	return
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}

func WriteToFile(data, file string) {
	os.WriteFile(file, []byte(data), 0777)
}

func WriteFile(content []byte, filename string) (err error) {
	// filepath := fmt.Sprintf("%s", filename)

	err = os.WriteFile(filename, content, 0644)
	if err != nil {
		return
	}
	return
}


