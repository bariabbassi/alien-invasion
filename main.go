package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	//There should be only 1 argument
	if len(os.Args) != 2 {
		fmt.Println("Use this format: aliens [N]")
		os.Exit(1)
	}

	//Read N the number of aliens
	N, err := strconv.Atoi(os.Args[1]) //the number of aliens
	if err != nil {
		fmt.Println("Argument N is not a number")
		os.Exit(1)
	}

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

	//Add N aliens to a random city
	w.AddAliens(N)

	fmt.Print(w)
}
