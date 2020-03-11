package utils

import (
	"fmt"
	"testing"
)

func Test_jwt_create(t *testing.T) {
	token, err := CreateToken(1001, 10)
	if err != nil {
		panic(err)
	}
	println(token)
}

func Test_jwt_parse_and_validate(t *testing.T) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODMxNTQxODcsImlhdCI6MTU4MzExODE4NywiaXNzIjoiZGF6dW8iLCJuYmYiOjE1ODMxMTgxODcsInVpZCI6MTAwMX0.mT9GVNVkflGj1XxRgYmt6xToJPAWqB_A_fitumt4oqM"
	println(ValidateToken(tokenString))

	payload, err := ParseToken(tokenString)
	if err != nil {
		panic(err)
	}
	println(payload)
}

func Test_MD5(t *testing.T) {
	val := Md5Encrpyt("123456")
	fmt.Println(val)
}

func Test_bean(t *testing.T) {

}
