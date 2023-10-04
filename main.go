package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/agp745/Interpreter-Go/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s, welcome to Monkey programming language!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n\n")
	repl.Start(os.Stdin, os.Stdout)
}