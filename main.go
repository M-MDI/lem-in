package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	startline           int
	endline             int
	slccontent          []string
	connectionStartLine int
	numOfAnt            int
	nameofStart         string
	nameofEnd           string
	numberofRooms       int
	antFarmRooms        *room
	startroombool       = false
	endroombool         = false
)

// readnote reads the contents of the input file and returns a slice of strings
// representing the rooms and connections
func readnote(textfile string) []string {
	content, err := ioutil.ReadFile("examples/" + textfile)
	if err != nil {
		log.Fatal(err)
	}
	// Store the contents of the input file as a slice of strings
	slccontent = strings.Split(string(content), "\n")
	// Remove comments from the input file
	for l := 0; l < len(slccontent); l++ {
		for t := 0; t < len(slccontent[l]); t++ {
			if string(slccontent[l][0]) == "#" {
				if string(slccontent[l][1]) != "#" {
					slccontent = append(slccontent[:l], slccontent[l+1:]...)
				}
				// Find the line numbers where the ##start and ##end room definitions begin
				for i := 0; i < len(slccontent); i++ {
					if slccontent[i] == "##start" {
						startline = i
						startroombool = true
					} else if !startroombool && i == len(slccontent)-1 {
						// If the start room is not defined, print an error message and exit the program
						fmt.Println("ERROR\nCheck the ##Start ")
						os.Exit(0)
					}
					if slccontent[i] == "##end" {
						endline = i
						endroombool = true
					} else if !endroombool && i == len(slccontent)-1 {
						// If the end room is not defined, print an error message and exit the program
						fmt.Println("ERROR\nCheck the ##End ")
						os.Exit(0)
					}
				}
			}
		}
	}
	RoomsandConnections := strings.Split(string(content), "\n")
	RoomsandConnections = append(RoomsandConnections[2:endline], RoomsandConnections[endline+1:]...)
	for l := 0; l < len(RoomsandConnections); l++ {
		for t := 0; t < len(RoomsandConnections[l]); t++ {
			if string(RoomsandConnections[l][0]) == "#" {
				if string(RoomsandConnections[l][1]) != "#" {
					RoomsandConnections = append(RoomsandConnections[:l], RoomsandConnections[l+1:]...)
				}
			}
		}
	}
	
	for m := len(RoomsandConnections) - 1; m >= 0; m-- {
		for k := 0; k < len(RoomsandConnections[m]); k++ {
			if string(RoomsandConnections[m][k]) == "-" {
				connectionStartLine = m
				break
			}
		}
	}
	stringfirstname := slccontent[startline+1]
	slicefirstname := strings.Split(stringfirstname, " ")
	nameofStart = slicefirstname[0]
	stringlastname := slccontent[endline+1]
	slicelastname := strings.Split(stringlastname, " ")
	nameofEnd = slicelastname[0]
	return RoomsandConnections
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("ERROR\nCheck the argument")
		os.Exit(0)
	}
	roomsandConnections := readnote(os.Args[1])
	numOfAnt, _ = strconv.Atoi(slccontent[0])
	if numOfAnt == 0 {
		fmt.Println("ERROR\nAnt Number 0 - Check The Ant Number")
		os.Exit(0)
	}
	// Create the ant farm based on the rooms and their connections
	antFarmRooms = Rooms(roomsandConnections)
	// Swap the ant farm to get the first and last room
	SwapFarm(Farm)
	// Add the last room to the ant farm
	Farm = append(Farm, *lastRm)
	// Get the total number of rooms in the ant farm
	numberofRooms = len(Farm)
	// Get the names of all the children of the first room
	var firstchildrennames []string
	for l := 0; l < len(firstRm.children); l++ {
		firstchildrennames = append(firstchildrennames, firstRm.children[l].name)
	}
	// Create the ants
	antfarm := CreatingAnts()
	// Find all possible paths between the rooms in the ant farm
	var allPaths [][]*room
	FindAllPossiblePaths(make([]*room, 0), Farm[0], &allPaths, &Farm[0])
	// Sort the paths based on their length
	for i := 0; i < len(allPaths); i++ {
		allPaths = SortPaths(allPaths)
	}
	// Clean the paths and remove any duplicate paths
	cleanway = ClearPath(allPaths)
	// Check if the last path is not in the appendWays slice and append it to the cleanway slice
	var anotherbool2 bool = false
	for i := 0; i < len(appendWays); i++ {
		for k := 0; k < len(allPaths[len(allPaths)-1]); k++ {
			if allPaths[len(allPaths)-1][k] == appendWays[i] {
				anotherbool2 = true
			}
		}
	}
	if !anotherbool2 {
		cleanway = append(CombinatedRooms, allPaths[len(allPaths)-1])
	}
	// Finding best paths for first children of the first room
	var Paths123 []Path
	var Paths1234 [][]Path
	Paths123 = FirstChildren(allPaths)
	for i := 0; i < childrenOfFirstRoom; i++ {
		Paths1234 = append(Paths1234, SortedPaths(Paths123, i))
	}
	Paths1234 = (SortAgain(Paths1234))
	Paths1234 = AllCombinations(Paths1234)
	
	// Finding intersecting paths
	FindIntersect(Paths1234)
	// Finding the best combinations of paths
	Paths1234 = FindBestCombinations(Paths1234)
	// Converting the best paths to rooms
	BestPath = PathtoRoom(Paths1234)
	BestPath = SortBestPath(BestPath)
	// Finding intersecting paths in the best path
	BestPath = FindAnotherIntersect(BestPath)
	BestPath = FindAnotherIntersect(BestPath)
	// Making sure that the number of ants in the antfarm is equal to the number of rooms in the best path
	antfarm = EqNum(antfarm, BestPath)
	// Making ants walk through the farm
	walk(antfarm)
}