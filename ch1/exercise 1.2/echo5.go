package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args[1:] {
		fmt.Print(i + 1)
		s := ": " + arg
		fmt.Println(s)
	}
}
