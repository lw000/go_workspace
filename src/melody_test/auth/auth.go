package auth

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
)

func Base64Encode(src []byte) string {
	if len(src) == 0 {
		return ""
	}

	return base64.StdEncoding.EncodeToString(src)
}

func Base64Decode(src string) ([]byte, error) {
	if len(src) == 0 {
		return nil, nil
	}

	return base64.StdEncoding.DecodeString(src)
}

func Md5Encode(src []byte) string {
	if len(src) == 0 {
		return ""
	}

	md5Ctx := md5.New()
	md5Ctx.Write(src)
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

func Sha1(src []byte) string {
	h := sha1.New()
	n, err := h.Write(src)
	if err != nil {
		return ""
	}
	if n < 0 {
		return ""
	}

	bs := h.Sum(nil)
	return hex.EncodeToString(bs)
}

func Sha224(src []byte) string {
	h := sha256.New224()
	n, err := h.Write(src)
	if err != nil {
		return ""
	}
	if n < 0 {
		return ""
	}

	bs := h.Sum(nil)
	return hex.EncodeToString(bs)
}

func Sha256(src []byte) string {
	h := sha256.New()
	n, err := h.Write(src)
	if err != nil {
		return ""
	}
	if n < 0 {
		return ""
	}

	bs := h.Sum(nil)
	return hex.EncodeToString(bs)
}

func Sha512(src []byte) string {
	h := sha512.New()
	n, err := h.Write(src)
	if err != nil {
		return ""
	}
	if n < 0 {
		return ""
	}

	bs := h.Sum(nil)
	return hex.EncodeToString(bs)
}
