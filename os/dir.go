package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Mkdir("temp", 0755)
	if err != nil {
		fmt.Printf("error in creating the directory: %v\n", err.Error())
		return
	}

	err = os.Chdir("temp")
	if err != nil {
		fmt.Printf("error in changing the directory: %v\n", err.Error())
		return
	}

	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Printf("error in creating the file: %v\n", err)
		return
	}

	defer file.Close()

	buffer := []byte("Hello George!")
	_, err = file.Write(buffer)
	if err != nil {
		fmt.Printf("error writing the data: %v\n", err.Error())
		return
	}

	files, err := os.ReadDir("temp")
	if err != nil {
		fmt.Printf("error in reading the directory: %v\n", err.Error())
		return
	}

	for _, fileReference := range files {
		fmt.Println(fileReference)
	}

	err = os.Remove("temp")
	if err != nil {
		fmt.Printf("error in remove the directory: %v\n", err)
		return
	}

}
