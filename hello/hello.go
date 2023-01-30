package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("What is your name? ")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	if err != nil {
		fmt.Println("Enter a string")
	}

	fmt.Printf("Hello, %s, nice to meet you!\n", name)

}
