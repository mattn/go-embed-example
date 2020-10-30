package main

import (
	_ "embed"
	"fmt"
)

//go:embed message.txt
var message string

func main() {
	fmt.Println(message)
}
