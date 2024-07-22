package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)

	elapsed := time.Since(start).Nanoseconds()
	fmt.Println("Elapsed time in for version: ", elapsed)

	start = time.Now()

	fmt.Println(strings.Join(os.Args[1:], " "))

	elapsed = time.Since(start).Nanoseconds()
	fmt.Println("Elapsed time in Join version: ", elapsed)
}
