package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func ReadFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], nil
	}
	return p, err
}

func sampleReadFrom() {
	data, _ := ReadFrom(strings.NewReader("from string"), 12)
	fmt.Println(string(data))
}

func sampleReadStdin() {
	fmt.Println("please input from stdion")
	data, _ := ReadFrom(os.Stdin, 11)
	fmt.Println(string(data))
}

func sampleReadFile() {
	file, _ := os.Open("test.go")
	defer file.Close()
	data, _ := ReadFrom(file, 12)
	fmt.Println(string(data))
}

func main() {
	sampleReadFrom()
	sampleReadStdin()
	sampleReadFile()
}   
