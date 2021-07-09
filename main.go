package main

import (
	"fmt"
	"os"
	"simple/lib"
	"simple/repl"
)

func main() {
	fmt.Printf("This is the Simple programming language!\n")
	fmt.Printf("Hi %s! Feel free to type in commands.\n", lib.GetUser())
	repl.Start(os.Stdin, os.Stdout)
}
