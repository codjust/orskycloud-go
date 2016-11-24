package comm

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"strconv"
)

func Md5_go(value interface{}) string {
	var str string
	switch value.(type) {
	case int:
		num, _ := value.(int)
		str = strconv.Itoa(num)
	case string:
		s, _ := value.(string)
		str = s
	}
	h := md5.New()
	io.WriteString(h, str)
	ret_md5 := h.Sum(nil)
	return hex.EncodeToString(ret_md5[:])
}
