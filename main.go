package main

import (
	"fmt"
	"os"
	"os/user"
	"github.com/codytheroux96/go-interpreter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Printf("Feel free to tupe in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}