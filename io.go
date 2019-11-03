package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

//ReadFile reads the file
func ReadFile(fileName string) ([]string, error) {
	file, err := os.Open(fileName)

	if err != nil {
		return nil, errors.New("Can't read the file " + fileName)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var fileLines []string

	for scanner.Scan() {
		fileLines = append(fileLines, scanner.Text())
	}

	file.Close()
	return fileLines, nil
}

//PrintFile prints the file lines
func PrintFile(fileLines []string) {
	for _, line := range fileLines {
		fmt.Println(line)
	}
}
