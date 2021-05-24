package golang_demo

import (
	"encoding/json"
	"fmt"
	"testing"
)

type User struct {
	UserName string `json:"username"`
	NickName string `json:"nickname"`
	Age      int
	Birthday string
	Sex      string
	Email    string
	Phone    string
}

type Users struct {
	All []*User
}

func TestTree(test *testing.T) {
	user1 := &User{
		UserName: "user1",
		NickName: "上课看似",
		Age:      18,
		Birthday: "2008/8/8",
		Sex:      "男",
		Email:    "mahuateng@qq.com",
		Phone:    "110",
	}

	var all []*User

	all = append(all, user1)

	data, err := json.Marshal(all)
	if err != nil {
		fmt.Printf("json.marshal failed, err: %s", err)
		return
	}

	fmt.Printf("%s\n", string(data))

}
