package main

import (
	"fmt"
	"bytes"
	"os"
	"io"
	"log"
)

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func LinesCount(file *os.File) int {
	stat, err := file.Stat()
	Check(err)

	buffer := make([]byte, stat.Size())

	n, err := file.Read(buffer)
	if n == 0 && err == io.EOF {
		return 0
	}
	Check(err)

	lineSeparator := []byte{'\n'}
	count := bytes.Count(buffer[:n], lineSeparator)

	return count
}

func main() {
	args := os.Args
	argumentsNumber := len(args)
	if argumentsNumber != 2 {
		log.Fatal("usage: go run path/lc.go path/FILE")
	}

	filepath := args[1]

	file, err := os.Open(filepath)
	Check(err)
	defer file.Close()

	count := LinesCount(file)
	fmt.Printf("%d " + filepath, count)
}


