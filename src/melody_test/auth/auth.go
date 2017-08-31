package auth

import (
	"crypto/md5"
	b64 "encoding/base64"
	"encoding/hex"
)

func Base64Encode(src []byte) string {
	return b64.StdEncoding.EncodeToString(src)
}

func Base64Decode(src string) ([]byte, error) {
	return b64.StdEncoding.DecodeString(src)
}

func Md5Encode(src []byte) string {
	md5Ctx := md5.New()
	md5Ctx.Write(src)
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
