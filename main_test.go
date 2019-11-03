package main

import (
	"testing"
)

//TestAlienInvasion tests all functions
func TestAlienInvasion(t *testing.T) {
	N := 5
	fileName := "file.txt"
	fileLines, err := ReadFile(fileName)
	if err != nil {
		t.Error(err)
	}
	if len(fileLines) != 9 {
		t.Error("Not the whole file has been read")
	}

	w := new(World)

	//Test that the world was filled
	err = w.FillWorld(fileLines)
	if err != nil {
		t.Error(err)
	}

	w.AddAliens(N)
	w.Fight()
	counter := 0
	for len(w.aliens) > 0 && counter < 10001 {
		w.MoveAliens()
		w.Fight()
		counter++
	}

	//Test that aliens have died
	if counter != 10001 && len(w.aliens) != 0 {
		t.Error("Aliens did not die during fights")
	}

	newFileLines := w.CreateFileLines()
	if len(newFileLines) >= 9 {
		t.Error("The file created is too short")
	}
}
