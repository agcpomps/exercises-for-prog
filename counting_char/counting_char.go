package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("What is the input string? ")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	if err != nil {
		fmt.Println("Error")
	}

	fmt.Printf("%s has %d characters\n", name, len(name))

}
