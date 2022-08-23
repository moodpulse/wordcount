package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println(0)
		return
	}

	src := strings.Trim(os.Args[1], " ")

	if src == "" {
		fmt.Println(0)
		return
	}

	words := strings.Split(src, " ")
	fmt.Println(len(words))
}
