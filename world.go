package main

import (
	"errors"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

////////////////////////////// City /////////////////////////////////////////////////

//City defines a node of the world
type City struct {
	name string
}

////////////////////////////// World /////////////////////////////////////////////////

//World defines a new world that contains cities roads and aliens
type World struct {
	cities []City
	roads  map[City][]City
	aliens map[City][]int
}

//RoadsString converts roads to a string
func (w *World) RoadsString() string {
	var str string
	str += "Roads:\n"
	for c1, roadsFromC1 := range w.roads {
		str += c1.name + " --> "
		for _, c2 := range roadsFromC1 {
			str += c2.name + " "
		}
		str += "\n"
	}
	return str
}

//AliensString converts aliens to a string
func (w *World) AliensString() string {
	var str string
	str += "Aliens:\n"
	for city, aliensInCity := range w.aliens {
		str += city.name + " <=< "
		for _, alienID := range aliensInCity {
			str += "Alien" + strconv.Itoa(alienID) + " "
		}
		str += "\n"
	}
	return str
}

func (w *World) String() string {
	return w.RoadsString() + w.AliensString()
}

//AddCity addds the city
func (w *World) AddCity(c City) {
	w.cities = append(w.cities, c)
}

//AddRoad destroys the road c1 --> c2
func (w *World) AddRoad(c1, c2 City) {
	w.roads[c1] = append(w.roads[c1], c2)
}

//FillWorld fills the world with fileLines
func (w *World) FillWorld(fileLines []string) error {
	w.roads = make(map[City][]City)
	w.aliens = make(map[City][]int)
	var c1 City
	var c2 City

	for _, line := range fileLines {
		fields := strings.Fields(line)

		//A file with a line with more than 5 fields is not accepted
		if len(fields) > 5 { //1 name + 4 fieldPrefixes = 5 fields
			return errors.New("Your file contains a line with more than 5 fields")
		}

		//Add cities
		c1 = City{fields[0]}
		w.AddCity(c1)

		//Add directions
		for i := 1; i < len(fields); i++ {
			direction := strings.Split(fields[i], "=")

			if len(direction) != 2 {
				return errors.New("The name of a city can not include =")
			}

			//A file with a line with a field thats not a direction is not accepted
			if direction[0] != "north" && direction[0] != "east" &&
				direction[0] != "south" && direction[0] != "west" {
				return errors.New("Your file contains a line with a field that's not a direction")
			}

			c2 = City{direction[1]}
			w.AddRoad(c1, c2)
		}
	}
	return nil
}

//AddAlien adds an alien to a city
func (w *World) AddAlien(alienID int, c City) {
	w.aliens[c] = append(w.aliens[c], alienID)
}

//AddAliens adds N aliens to random cities
func (w *World) AddAliens(N int) {
	rand.Seed(time.Now().Unix())
	for alienID := 1; alienID < N+1; alienID++ {
		randomCity := w.cities[rand.Intn(len(w.cities))]
		w.AddAlien(alienID, randomCity)
	}
}
