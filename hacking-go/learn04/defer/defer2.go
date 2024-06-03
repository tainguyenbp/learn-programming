package main

import (
	"fmt"
	"os"
)

func openfile() {
	// Open a file
	file, err := os.Open("tainguyenbp.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Defer the closing of the file
	defer file.Close()

	// Perform operations on the file
	// For demonstration, let's just print its contents
	fmt.Println("File contents:")
	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println(string(data[:count]))
}

func openfileno() {
	// Open a file
	file, err := os.Open("tainguyenbp1.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Defer the closing of the file
	defer file.Close()

	// Perform operations on the file
	// For demonstration, let's just print its contents
	fmt.Println("File contents:")
	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println(string(data[:count]))
}

func main() {

	openfile()
	openfileno()
	// Any other operations on the file can go here
}
