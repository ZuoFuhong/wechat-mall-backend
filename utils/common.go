package utils

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"os"
	"strings"
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

func RandomNumberStr(lenth int) string {
	chars := [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
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

func PhoneMark(phone string) string {
	return phone[0:3] + "****" + phone[7:]
}

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// 检查目录是否存在，并创建
func CheckFileDirExists(filepath string) {
	index := strings.LastIndex(filepath, "/")
	dirPath := filepath[0:index]
	exists, e := PathExists(dirPath)
	if e != nil {
		panic(e)
	}
	if !exists {
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
}
