package main

import (
	//"github.com/gin-gonic/gin"
	// "time"
	"crypto/rand"
	"fmt"
	"io"
)

type User struct {
	ID          int `json:"id"`
	UserName    string `json:"user_name"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	DisplayName string `json:"display_name"`
}
type UserCollection struct{}

func (col *UserCollection) NewUUID() (string, error) { //create token
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil

}
func (col *UserCollection) CreateUser(user *User) (int, error) {
	var id int
	n, err1 := col.NewUUID()
	if err1 != nil {
		panic(err1)
	}
	fmt.Println(n)
		res := fmt.Sprintf("insert into users() values(%s,%s,%s,%s)", user.UserName, user.Password, user.Email, user.DisplayName)
	err := DB.QueryRow(res)

	fmt.Printf("Last Inserted Id : %d", id)
	if err != nil {
		return 0, nil
	}
	return id, nil
}
