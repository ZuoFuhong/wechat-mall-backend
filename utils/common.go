package utils

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

func RandomStr(lenth int) string {
	chars := [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H',
		'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b',
		'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v',
		'w', 'x', 'y', 'z'}

	tmpArr := make([]byte, 0, lenth)
	rand.Seed(time.Now().Unix())
	for i := 0; i < lenth; i++ {
		n := rand.Intn(len(chars))
		tmpArr = append(tmpArr, chars[n])
	}
	return string(tmpArr)
}

func Md5Encrpyt(passwd string) string {
	ctx := md5.New()
	ctx.Write([]byte(passwd))
	ctx.Write([]byte("salt123"))
	return hex.EncodeToString(ctx.Sum(nil))
}
