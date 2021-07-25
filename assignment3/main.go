package main

import (
	"fmt"
	"io"
	"os"
)

type logWriter struct{}

func main() {
	fmt.Println(os.Args)
	fileName := os.Args[1]
	fmt.Println(fileName)

	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	io.Copy(os.Stdout, f)
}
