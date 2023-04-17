package main

import (
	"fmt"
	"os"
)

func main() {
	//
	// func WriteFile(name string, data []byte, perm FileMode) error

	str := "helle world !!!!!!"
	err := os.WriteFile("./hello", []byte(str), 0666)
	if err != nil {
		// handle error
		fmt.Println(err)
	}
}
