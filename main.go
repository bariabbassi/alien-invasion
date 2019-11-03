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
		fmt.Println("Argument N should be a number")
		os.Exit(1)
	}
	if N < 1 {
		fmt.Println("The number of aliens should be greater than 0")
		os.Exit(1)
	} else if N > 100 {
		fmt.Println("Try a smaller number of aliens")
		os.Exit(1)
	}

	//Read the map file
	fileName := "file.txt"
	fileLines, err := ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("********************")
	fmt.Println("** Peaceful world **")
	fmt.Println("********************")

	//Print the file
	PrintFile(fileLines)
	fmt.Println()

	//Create the world
	w := new(World)

	//Fill the world
	err = w.FillWorld(fileLines)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//Add N aliens to a random city
	w.AddAliens(N)

	fmt.Println("*************************************")
	fmt.Println("** OH NO! ALIENS ARE EVERYWHERE!!! **")
	fmt.Println("*************************************")

	var outputMove string
	var outputFight string

	//If more than 1 alien is added to a city they fight
	outputFight = w.Fight()
	fmt.Print(outputFight)

	//Move aliens until they get in a fight or get stuck
	counter := 0
	for len(w.aliens) > 0 && counter < 10001 {
		outputMove = w.MoveAliens()
		fmt.Print(outputMove)
		outputFight = w.Fight()
		fmt.Print(outputFight)
		counter++
	}
	fmt.Println()

	fmt.Println("*******************************")
	fmt.Println("** Post-alien-invasion world **")
	fmt.Println("*******************************")

	//Print what the file would look like now
	newFileLines := w.CreateFileLines()
	PrintFile(newFileLines)
}
