package main

import (
	"fmt"
	"os"
)

func main() {

	// creating the file
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Printf("error in creating the file: %v\n", err)
		return
	}
	defer file.Close()

	// write the file
	_, err = file.WriteString("Hello World!!\n")
	if err != nil {
		fmt.Printf("error in writing the file: %v\n", err)
		return
	}

	// close the file
	file.Close()

	// open the file
	file, err = os.Open("example.txt")
	if err != nil {
		fmt.Printf("error in open the file: %v\n", err)
		return
	}
	defer file.Close()

	// get the state of the file
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("error in getting the file state: %v\n", err)
		return
	}

	// read the file
	buffer := make([]byte, fileInfo.Size())
	_, err = file.Read(buffer)
	if err != nil {
		fmt.Printf("error in read the file: %v\n", err)
		return
	}

	file.Close()

	fmt.Printf("%s\n", buffer)

	// open the file with append mode
	file, err = os.OpenFile("example.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("error in open the file: %v\n", err)
		return
	}

	// append the file
	content := []byte("New content to append\n")
	if _, err = file.Write(content); err != nil {
		fmt.Printf("error in append the file: %v\n", err)
		return
	}

	// remove the file
	// err = os.Remove("example.txt")
	// if err != nil {
	// 	fmt.Printf("error in remove the file: %v\n", err)
	// 	return
	// }

	fmt.Printf("Success")

}
