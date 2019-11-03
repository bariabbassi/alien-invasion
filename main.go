package main

import (
	"fmt"
	"os"
)

func main() {
	//Read the map file
	fileName := "file.txt"
	fileLines, err := ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//Print the file
	PrintFile(fileLines)
	fmt.Println()

	//Create the world
	w := new(World)

	//Fill the world
	if w.FillWorld(fileLines) != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Print(w.RoadsString())
}
