package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var count int
	for {
		_, isPrefix, err := reader.ReadLine()
		if err != nil {
			break
		}
		if !isPrefix {
			count++
		}
	}

	fmt.Println(count)
}
