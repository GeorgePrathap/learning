package main

import (
	"fmt"
	"os"
)

func main() {

	// get environment variable
	envHome := os.Getenv("HOME")
	fmt.Printf("Home env variable :%v\n", envHome)

	// set environment variable
	err := os.Setenv("MY_VAR", "George")
	if err != nil {
		fmt.Printf("error in setting the environment variable: %v\n", err.Error())
		return
	}

	// set environment variable
	err = os.Setenv("MY_NAME", "George Pradap")
	if err != nil {
		fmt.Printf("error in setting the environment variable: %v\n", err.Error())
		return
	}

	// get environment variable
	envMyVar := os.Getenv("My_VAR")
	fmt.Printf("MyVar env variable: %v\n", envMyVar)

	// get all the env at once
	envVars := os.Environ()
	for _, envVar := range envVars {
		fmt.Printf("value : %v\n", envVar)
	}

	// Clear all environment variables, don't uncomment that you will loose all of your data
	// os.Clearenv()
}
