package main

import (
	"fmt"
	"gorilla/relp"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is The Gorilla programming language!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")
	relp.Start(os.Stdin, os.Stdout)
}
