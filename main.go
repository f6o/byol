package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("lispy> ")
		if scanner.Scan() {
			input := scanner.Text()
			fmt.Printf("No you're a %s\n", input)
		} else {
			break
		}
		if scanner.Err() != nil {
			break
		}
	}
}
