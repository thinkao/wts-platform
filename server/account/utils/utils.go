package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	uuid "github.com/satori/go.uuid"
	"path"
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

func ImgToUUID(img string) string {

	/*// 创建 UUID v4
	u1 := uuid.Must(uuid.NewV4())
	println(`生成的UUID v4：`)
	println(u1.String())*/

	// 创建可以进行错误处理的 UUID v4
	/*u2 := uuid.NewV4()
	if err1 != nil {
		println(`生成一个UUID v4时出现错误：`)
		panic(err1)
	}
	println(`生成的UUID v4：`)
	println(u2.String())*/
	//name := path.Base(img)

	ext := path.Ext(img)

	u2 := uuid.NewV4()

	return u2.String() + ext

}
