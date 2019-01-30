package lwauth

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(d []byte) string {
	m := md5.New()
	n, err := m.Write(d)
	if err != nil {
		return ""
	}

	if n < 0 {
		return ""
	}

	s := m.Sum(nil)
	return hex.EncodeToString(s)
}
