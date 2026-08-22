package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/tushar434434/tushar/auth"
	"github.com/tushar434434/tushar/user"
)

func main() {
	auth.Loginwith("tushar")

	session := auth.Getsession()
	fmt.Println("session", session)

	u := user.User{
		Email: "tushar2gmail.com",
		Name:  "Tushar",
	}

	fmt.Println(u.Email, u.Name)

	color.Red(u.Email)
}
