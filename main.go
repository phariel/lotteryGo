package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter command:")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		switch scanner.Text() {
		case "fetch":
			fetchDataToDb()
			fmt.Println("Enter command:")
			break
		default:
			fmt.Println("Command not found, enter again:")
		}
	}
}
