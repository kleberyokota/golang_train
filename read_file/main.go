package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fileName := os.Args
	file, err := ioutil.ReadFile(fileName[1])
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(string(file))
}
