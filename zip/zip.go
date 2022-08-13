package zip

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/finest08/go-encryption-app/utils"
)

func ZipSource(source string) error {
	// 1. Create a ZIP file and zip.Writer
	f, err := os.Create(source + ".zip")
	if err != nil {
		return err
	}
	defer f.Close()

	writer := zip.NewWriter(f)
	defer writer.Close()

	// 2. Go through all the files of the source
	return filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 3. Create a local file header
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		// set compression
		header.Method = zip.Deflate

		// 4. Set relative path of a file as the header name
		header.Name, err = filepath.Rel(filepath.Dir(source), path)
		if err != nil {
			return err
		}
		if info.IsDir() {
			header.Name += "/"
		}

		// 5. Create writer for the file header and save content of the file
		headerWriter, err := writer.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = io.Copy(headerWriter, f)
		return err
	})
}

func UnzipSource(src string) error {
	dst := "playground"
	archive, err := zip.OpenReader(src)
	if err != nil {
		panic(err)
	}
	defer archive.Close()

	for _, f := range archive.File {
		filePath := filepath.Join(dst, f.Name)
		// fmt.Println("unzipping file ", filePath)

		if !strings.HasPrefix(filePath, filepath.Clean(dst)+string(os.PathSeparator)) {
			fmt.Println("invalid file path")
			return err
		}
		if f.FileInfo().IsDir() {
			// fmt.Println("creating directory...")
			os.MkdirAll(filePath, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			panic(err)
		}

		dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			panic(err)
		}

		fileInArchive, err := f.Open()
		if err != nil {
			panic(err)
		}

		if _, err := io.Copy(dstFile, fileInArchive); err != nil {
			panic(err)
		}

		dstFile.Close()
		fileInArchive.Close()
	}
	return err
}

func Compress(src string) string {
	alExis := utils.Exists(src)
	srcExists := utils.Exists(src)
	if !srcExists {
		fmt.Println("Source does not exist")
		return ""
	}

	if !alExis {
		fmt.Println("File is already compressed")
	} else {
		np := fmt.Sprintf("%s&mwx&&", src)
		err := os.Rename(src, np)
		if err != nil {
			fmt.Println(err)
		}
		err = ZipSource(np)
		if err != nil {
			fmt.Println(err)
		}

		err = os.RemoveAll(np)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Compression complete.")
		return np
	}
	return ""
}

func Decompress(src string) {
	srcExists := utils.Exists(src)
	if !srcExists {
		fmt.Println("Source does not exist")
		return
	} else {

		err := UnzipSource(src)
		if err != nil {
			fmt.Println(err)
		}
		err = os.RemoveAll(src)
		if err != nil {
			fmt.Println(err)
		}
		ns := strings.Replace(src, ".zip", "", -1)
		s := strings.Replace(ns, "&mwx&&", "", -1)

		err = os.Rename(ns, s)
		if err != nil {
			fmt.Println(err, "Error renaming file")
		}
		fmt.Println("Decompression complete.")
	}

}
