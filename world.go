package main

import (
	"errors"
	"fmt"
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

//Equals compares 2 cities
func (city *City) Equals(c City) bool {
	return city.name == c.name
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

//DestroyCity destroys the city
func (w *World) DestroyCity(c City) {
	len := len(w.cities)
	i := 0

	//Find the index of the city
	for i < len {
		if w.cities[i].Equals(c) {
			break
		}
		i++
	}

	//Delte the city at the index i
	if i != len {
		w.cities[i] = w.cities[len-1]
		w.cities[len-1] = City{}
		w.cities = w.cities[:len-1]
	}

	//When a city is distroyed the roads from and to to city get destroyed too
	w.DestroyRoads(c)

	//the aliens in the city are killed aswell
	w.aliens[c] = nil
}

//DestroyRoad deletes the road c1 --> c2
func (w *World) DestroyRoad(c1, c2 City) {
	len := len(w.roads[c1])
	i := 0

	//Find the index of the road
	for i < len {
		if w.roads[c1][i].Equals(c2) {
			break
		}
		i++
	}

	//Delete the road at the index i
	if i != len {
		w.roads[c1][i] = w.roads[c1][len-1]
		w.roads[c1][len-1] = City{}
		w.roads[c1] = w.roads[c1][:len-1]
	}
}

//DestroyRoads deletes all the roads from and to c <-->
func (w *World) DestroyRoads(c City) {
	//delete all roads to c <--
	for _, c2 := range w.roads[c] {
		w.DestroyRoad(c2, c)
	}

	//delete all the roads from c -->
	delete(w.roads, c)
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

//Fight finds cities with more than 1 alien, kills them and destroys the city
func (w *World) Fight() string {
	var str string
	for city, aliensInCity := range w.aliens {

		//More than 1 alien in a city
		if len(aliensInCity) > 1 {
			str += city.name + " has been destroyed by "
			for _, alienID := range aliensInCity {
				str += "Alien" + strconv.Itoa(alienID) + " "
			}

			//DestroyCity() destroys the city and kills aliens in the city
			w.DestroyCity(city)

			//Print the phrase : CityX has been destroyed by AlienY and AlienZ
			fmt.Println(str)
			str = ""
		}
	}
	return str
}

//MoveAliens moves each alien to a new city
func (w *World) MoveAliens() {
	rand.Seed(time.Now().Unix())

	//A new aliens map is filed and replaces the old aliens map
	newAliens := make(map[City][]int)

	//Fill the new aliens map
	for city, aliensInCity := range w.aliens {
		if aliensInCity != nil {

			//Aliens stuck in a city stay in the same city
			if len(w.roads[city]) == 0 {

				//His next city is the same city
				nextCity := city

				//Copy alien to the new aliens map
				newAliens[nextCity] = append(newAliens[nextCity], aliensInCity[0])

				//Aliens that are not stuck go to the next random city
			} else {

				//Pick a random road
				nextCity := w.roads[city][rand.Intn(len(w.roads[city]))]

				//Copy alien to the new aliens map
				newAliens[nextCity] = append(newAliens[nextCity], aliensInCity[0])
			}
		}
	}

	//Replace the old aliens map with the new aliens map
	w.aliens = newAliens
}
