package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
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
