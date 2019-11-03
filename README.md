# Aliens invasion
Alien invasion :alien: is a fun little simulation of an alien invasion.

## How to run the project
- Required: ```Go 1.11+```
- Clone this director within your ```gopath```.
```
go install github.com/bariabbassi/aliens
aliens [N]
```
- Replace ```[N]``` with a number between 1 and 100 in the previous line.

## What the project does
- The user chouses a number N of aliens.
- The program reads the file ```file.txt``` included in the project.
- This file describes the location of different cities.
- The file data is transformed into a directed graph called World where the cities are the nodes and neighboring cities define the roads between the cities.
- N aliens are unleashed into the world randomly and make a step to the next random city all at the same time.
- When 2 or more aliens land in the same city they fight and die and the city and its roads get destroyed.
- When an alien gets trapped in a city with no roads to get out he dies.
- The movement of the aliens stops when either all aliens have died or the number of moves has reached 10000 moves.

## Assumptions made
- The number of aliens has to be between 1 and 100 because unleashing too many aliens makes the game less fun to watch.
- An alien inside a city without roads to exit this city is killed because if he is left to live he will stay alive forever and the game won't stop until the cap of 10000 moves.
- The 4 directions north, east, south, and west are considered equal and are all just meaningless tags. A road is defined only by its origin and destination cities. This makes the world a UNWEIGHTED directed graph (state machine) with the roads being the arcs connecting the nodes (cities).

## Overview of the project
- ```file.txt``` consists of lines containing the name of a city followed by neighboring cities.
```
Amsterdam east=Berlin
Berlin south=Bern
…
```

This is what ```file.txt``` would look like visually.
```
          Amsterdam -----> Berlin
              |              |
Dublin ----> Paris <------> Bern <-----> Rome
              |
            Madrid
              |
            Lisbon
```

- ```io.go``` contains the 2 functions that read and print the file ```file.txt```. The name ```io``` stands for input and output.

- ```main.go``` calls functions from world.go and ```io.go``` needed to perform the flow of actions described in "What the project does" and deals with errors.

- ```world.go``` defines 2 structures ```World``` and ```City```. ```World``` is the directed graph and ```City``` is a node in the graph. The ```aliens``` in the world are defined just with a unique number without a name and ```roads``` defined by lists of destinations accessible from a city.
```
type City struct {
	name string
}
```
```
type World struct {
	cities []City
	roads  map[City][]City
	aliens map[City][]int
}
```
```world.go``` contains all the functions needed to perform the flow of actions needed in ```main.go```. The functions can, for example, add and delete cities, roads, and aliens.



