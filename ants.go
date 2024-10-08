package main

import (
	"fmt"
)

type ant struct {
	id        int
	curRoom   *room
	pathOfAnt []*room
	step      int
}

type Path struct {
	id        int	// unique identifier for the path.
	paths     []*room	// the rooms in the path.
	intersect bool	// whether the path intersects with another path.
	queue     int	// the number of ants waiting to use the path.
	totalLen  int	// the total length of the path.
}
// EndRoomUsed represents whether the end room of a path has already been used.
type EndRoomUsed struct {
	whichPath []*room	// the end room of the path.
	used      bool		// whether the end room has been used.
}

var (
	ants                []ant
	roomspassed         int
	Combinations        [][]*room
	allways             [][]*room
	way                 []*room
	way2                []*room
	cleanway            [][]*room
	childrenOfFirstRoom int
	countloop           int
	appendWays          []*room
	CombinatedRooms     [][]*room
	BestCombinations    [][]Path
	intersectbool       bool = false
	BestPath            [][]*room
	counter             int
)
// CreatingAnts creates a slice of ants.
func CreatingAnts() []ant {
	if firstRm != nil {
		for i := 1; i < numOfAnt+1; i++ {
			var allants ant

			allants.id = i
			allants.curRoom = firstRm

			ants = append(ants, allants)
		}
	}

	return ants
}

func SwapFarm(Farm []room) {
	for i := 0; i < len(Farm)/2; i++ {
		Farm[i], Farm[(len(Farm)-1)-i] = Farm[(len(Farm)-1)-i], Farm[i]
	}
}
// This function simulates the movement of ants in the ant farm
// It takes a slice of ants as input
func walk(antfarm []ant) {
	var allpassed bool = true
	var EndRoomStr EndRoomUsed
	for i := 0; i < len(antfarm); i++ {
		
		// If an ant has already reached the last room, skip to the next ant
		if antfarm[i].curRoom.name == lastRm.name {
			continue
		}
		// If an ant is not in the first room, mark the previous room as unoccupied
		if antfarm[i].curRoom.name != firstRm.name {
			antfarm[i].curRoom.occupied = false
		}
		// If the next room an ant wants to move to is already occupied, skip to the next ant
		if antfarm[i].pathOfAnt[antfarm[i].step].occupied {
			continue
		}
		// If an ant has already used a particular path to reach the last room, skip to the next ant
		if EndRoomStr.used && EndRoomStr.whichPath[0].name == antfarm[i].pathOfAnt[0].name && antfarm[i].pathOfAnt[antfarm[i].step].name == lastRm.name {
			continue
		}
		// Move an ant to the next room in its path
		antfarm[i].curRoom = antfarm[i].pathOfAnt[antfarm[i].step]
		antfarm[i].step++
		if antfarm[i].curRoom.name != lastRm.name {
			antfarm[i].curRoom.occupied = true
		}
		// If an ant has reached the last room, store the path used by the ant
		if antfarm[i].curRoom.name == lastRm.name {
			EndRoomStr.used = true
			EndRoomStr.whichPath = antfarm[i].pathOfAnt
		}
		// Update the allpassed variable to false if an ant has not yet reached the last room
		allpassed = false
		// Print the movement of an ant in the ant farm
		fmt.Print("L", antfarm[i].id, "-", antfarm[i].curRoom.name, " ")

	}
	// If all ants have reached the last room, return
	if allpassed {
		return
	// If some ants have not yet reached the last room, set the end room used by an ant to false
	// and call the walk function recursively
	} else {
		EndRoomStr.used = false
		EndRoomStr.whichPath = nil

		fmt.Println(" ")
		walk(antfarm)
	}
}

func ShortestPath(rm *room) {
	if rm.name != lastRm.name {
		roomspassed++
		for i := 0; i < len(rm.children); i++ {
			ShortestPath(rm.children[i])
		}
	} else {

	}
}

func FindingPath(currentRoom *room) [][]*room {
	if len(currentRoom.children) != 0 {
		for i := 0; i < len(currentRoom.children); i++ {
			way2 = append(way2, currentRoom.children[i])
			FindingPath(currentRoom.children[i])
		}
	} else {

		allways = append(allways, way2)
		way2 = nil
	}

	return allways
}

// FindAllPossiblePaths is a recursive function that searches for all possible paths from the starting room to the ending room.
func FindAllPossiblePaths(path []*room, currentRoom room, paths *[][]*room, previousRoom *room) {
	// if the current room is the last room, append the current path to the paths slice
	if currentRoom.name == lastRm.name {
		// Check if the path goes back to the first room, and skip it if it does
		var skipPath bool
		for i := 0; i < len(path); i++ {
			if path[i].name == firstRm.name {
				skipPath = true
				break
			}
		}

		if len(*paths) == 0 {
			*paths = append((*paths), nil)
		} else if (*paths)[len(*paths)-1] != nil {
			*paths = append((*paths), nil)
		}

		for i := 0; i < len(path); i++ {
			if !skipPath {
				(*paths)[len(*paths)-1] = append((*paths)[len(*paths)-1], path[i])
			} else {
				break
			}
		}
	}
	// Recursively explore all possible paths from the current room to other rooms
	for i := 0; i < len(currentRoom.children); i++ {
		var toContinue bool

		for k := 0; k < len(path); k++ {
			if path[k].name == currentRoom.children[i].name {
				toContinue = true
				break
			}
		}

		if !toContinue {
			pathToPass := path
			pathToPass = append(pathToPass, currentRoom.children[i])
			FindAllPossiblePaths(pathToPass, *currentRoom.children[i], paths, &currentRoom)
			pathToPass = path
		}
	}
	// remove any empty paths from the paths slice
	for i := 0; i < len(*paths); i++ {
		if (*paths)[i] == nil {
			*paths = append((*paths)[:i], (*paths)[i+1:]...)
		}
	}
}

func SortPaths(ways [][]*room) [][]*room {
	for i := 0; i < len(ways)-1; i++ {
		if len(ways[i]) > len(ways[i+1]) {
			ways[i], ways[i+1] = ways[i+1], ways[i]
		}
	}

	for k := 0; k < len(ways)-1; k++ {
		if len(ways[len(ways)-1]) < len(ways[k]) {
			ways[len(ways)-1], ways[k] = ways[k], ways[len(ways)-1]
		}
	}
	return ways
}

func ClearPath(ways [][]*room) [][]*room {
	var somebool bool = false
	var anotherbool bool = false
	childrenOfFirstRoom = len(firstRm.children)
	if appendWays == nil {
		appendWays = ways[0]
	}
	if CombinatedRooms == nil {
		CombinatedRooms = append(CombinatedRooms, ways[0])
	}
	if countloop == len(ways)-1 {
		return CombinatedRooms
	}
	for i := 0; i < len(ways[countloop+1]); i++ {
		for k := 0; k < len(appendWays)-1; k++ {
			if ways[countloop+1][i].name == appendWays[k].name {
				somebool = true
			}
			if !somebool && i == len(ways[countloop+1])-1 && k == len(appendWays)-2 {
				appendWays = append(appendWays, ways[countloop+1]...)
				anotherbool = true
			}
		}
	}
	if anotherbool {
		CombinatedRooms = append(CombinatedRooms, ways[countloop+1])
	}
	countloop++
	ClearPath(ways)
	return CombinatedRooms
}

func FirstChildren(ways [][]*room) []Path {
	var PathStruct Path
	var PathStruct2 []Path
	childrenOfFirstRoom = len(firstRm.children)
	for i := 0; i < len(ways); i++ {
		for k := 0; k < childrenOfFirstRoom; k++ {
			if ways[i][0] == firstRm.children[k] {
				PathStruct.id = k
				PathStruct.paths = ways[i]
				PathStruct.intersect = true
				PathStruct2 = append(PathStruct2, PathStruct)

			}
		}
	}
	return PathStruct2
}

func SortedPaths(way []Path, idChildren int) []Path {
	var SepPath []Path
	for i := 0; i < len(way); i++ {
		if way[i].id == idChildren {
			SepPath = append(SepPath, way[i])
		}
	}
	return SepPath
}

func SortAgain(way [][]Path) [][]Path {
	for i := 0; i < len(way)-1; i++ {
		if len(way[i]) < len(way[i+1]) {
			way[i], way[i+1] = way[i+1], way[i]
		}
	}
	return way
}

func AllCombinations(way [][]Path) [][]Path {
	var AnotherPath []Path
	var AnotherPath2 [][]Path
	if childrenOfFirstRoom == 2 {
		for i := 0; i < len(way[0]); i++ {
			for k := 0; k < len(way[1]); k++ {
				AnotherPath = append(AnotherPath, way[0][i], way[1][k])
				AnotherPath2 = append(AnotherPath2, AnotherPath)
				AnotherPath = nil

			}
		}
	} else if childrenOfFirstRoom == 3 {
		for i := 0; i < len(way[0]); i++ {
			for k := 0; k < len(way[1]); k++ {
				for l := 0; l < len(way[2]); l++ {
					AnotherPath = append(AnotherPath, way[0][i], way[1][k], way[2][l])

					AnotherPath2 = append(AnotherPath2, AnotherPath)
					AnotherPath = nil
				}
			}
		}
	} else if childrenOfFirstRoom == 4 {
		for i := 0; i < len(way[0]); i++ {
			for k := 0; k < len(way[1]); k++ {
				for t := 0; t < len(way[3]); t++ {
					AnotherPath = append(AnotherPath, way[0][i], way[1][k], way[3][t])

					AnotherPath2 = append(AnotherPath2, AnotherPath)
					AnotherPath = nil
				}
			}
		}
	} else if childrenOfFirstRoom == 1 {
		AnotherPath2 = append(AnotherPath2, way[0])
	}

	return AnotherPath2
}

func FindIntersect(way [][]Path) [][]Path {
	for i := 0; i < len(way); i++ {
		for k := 0; k < len(way[i])-1; k++ {
			for l := 0; l < len(way[i][k].paths)-1; l++ {
				for t := 0; t < len(way[i][k+1].paths)-1; t++ {
					if way[i][k].paths[l].name == way[i][k+1].paths[t].name {
						way[i][k].intersect = false
						way[i][k+1].intersect = false

					}
				}
			}
		}
	}
	return way
}

func FindBestCombinations(way [][]Path) [][]Path {
	for i := 0; i < len(way); i++ {
		intersectbool = false
		for k := 0; k < len(way[i]); k++ {
			for t := 0; t < len(way[i][k].paths); t++ {
			}
			if !way[i][k].intersect {
				intersectbool = true
			} else if k == len(way[i])-1 && !intersectbool {
				BestCombinations = append(BestCombinations, way[i])
			}
		}
	}
	if BestCombinations == nil {
		BestCombinations = append(BestCombinations, way[0])
	}
	return BestCombinations
}

func PathtoRoom(way [][]Path) [][]*room {
	if len(way) > 1 {
		for i := 0; i < len(way)-1; i++ {
			countofpaths := 0
			for k := 0; k < len(way[i]); k++ {
				countofpaths += len(way[i][k].paths)
				way[i][k].totalLen = countofpaths
			}
		}
		for i := 0; i < len(way)-1; i++ {
			if way[i][len(way[i])-1].totalLen > way[i+1][len(way[i])-1].totalLen {
				way[i], way[i+1] = way[i+1], way[i]
			}
		}
	}
	for i := 0; i < len(way[0]); i++ {
		BestPath = append(BestPath, way[0][i].paths)
	}

	return BestPath
}

func SortBestPath(way [][]*room) [][]*room {
	for i := 0; i < len(way)-1; i++ {
		if len(way[i]) > len(way[i+1]) {
			way[i], way[i+1] = way[i+1], way[i]
		}
	}
	return way
}

func EqNum(ants []ant, Path [][]*room) []ant {
	countofants := len(ants)
	for i := 0; i < len(Path)-1; i++ {
		if len(Path[i])+Path[i][0].queue > len(Path[i+1])+Path[i+1][0].queue {
			Path[i], Path[i+1] = Path[i+1], Path[i]
		}
	}
	ants[counter].pathOfAnt = Path[0]
	Path[0][0].queue++
	counter++
	if counter != countofants {
		EqNum(ants, Path)
	}
	return ants
}

func Dist(ants []ant, Paths [][]*room) []ant {
	for i := 0; i < len(ants); i++ {
		ants[i].pathOfAnt = Paths[i]
	}
	if len(Paths[0]) == 1 {
		ants[len(ants)-1].pathOfAnt = Paths[0]
	}
	return ants
}

func FindAnotherIntersect(way [][]*room) [][]*room {
	for i := 0; i < len(way)-1; i++ {
		for k := 0; k < len(way[i])-1; k++ {
			for t := 0; t < len(way[i+1])-1; t++ {
				if way[i][k].name == way[i+1][t].name {
					way = append(way[:i+1], way[i+2:]...)
					break
				}
			}
		}
	}
	return way
}