package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	pass := `password1001`
	bs, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(pass)
	fmt.Println(bs)

	login := `password1001`
	err = bcrypt.CompareHashAndPassword(bs, []byte(login))
	if err != nil {
		fmt.Println("You can't login")
		return
	}
	fmt.Println("You're Logged In")
}
