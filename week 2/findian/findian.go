package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func main () {
	
	var str string
	var first string
	var last string

	fmt.Printf("Enter a string:\n")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	str = scanner.Text()
	str = strings.ToUpper(str)

	first = str[0:1]
	last = str[len(str)-1:]
	
	if (first == "I" && last == "N" && strings.Contains(str, "A")) {
		fmt.Printf("Found!\n")
	} else {
		fmt.Printf("Not Found!\n")
	}
}

