package utils

import (
	"bytes"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	uuid "github.com/satori/go.uuid"
	"io"
	"log"
	"mime/multipart"
)

func Encrypt(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	sha256_str := hex.EncodeToString(h.Sum(nil))

	length := len(sha256_str) / 8
	slice_str := sha256_str[:length] + sha256_str[4*length:5*length] + sha256_str[7*length:]

	m := md5.New()
	m.Write([]byte(slice_str))
	return hex.EncodeToString(m.Sum(nil))
}

func ImgToUUID(img *multipart.FileHeader) string {


	u2 := uuid.NewV4()



	b := new(bytes.Buffer)
	w := multipart.NewWriter(b)


	fileWriter,err := w.CreateFormFile("file",img.Filename)

	f, err := img.Open()
	if err != nil {
		log.Fatal("opening file:", err)
	}


	_, err = io.Copy(fileWriter, f)


	return u2.String()

}
