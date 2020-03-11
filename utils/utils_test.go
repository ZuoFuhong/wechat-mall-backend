package utils

import (
	"container/list"
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

func Test_List(t *testing.T) {
	tmpList := list.New()
	tmpList.PushFront(1)
	tmpList.PushFront("dazuo")

	pushList(tmpList)

	for item := tmpList.Front(); item != nil; item = item.Next() {
		fmt.Println(item.Value)
	}
}

func pushList(tmp *list.List) *list.List {
	tmp.PushFront("age")
	return tmp
}

func Test_slice(t *testing.T) {
	tmp := []string{"one", "two"}

	appendSlice(&tmp)
	fmt.Println(tmp)
}

func appendSlice(tmp *[]string) {
	*tmp = append(*tmp, "three")
}

func Test_Map(t *testing.T) {
	resp := make(map[string]interface{}, 0)
	resp["bannerList"] = 23
	fmt.Println(resp)
	fmt.Println(len(resp))
}

func Test_switch(t *testing.T) {
	a := 1
	switch a {
	case 1:
		fmt.Println("1")
	case 2, 3:
		fmt.Println("2")
	}
}
