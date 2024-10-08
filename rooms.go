package main

import (
	"fmt"
	"os"
	"strings"
)

type room struct {
	parent   [](*room)
	name     string
	children [](*room)
	occupied bool
	queue    int
}

var (
	firstRm    *room
	lastRm     *room
	Farm       []room
	Farm2      []*room
	test       []*room
	count      int
	checkdup   []string
	firstendrm int
)
// Rooms function takes in an array of strings that contain the rooms and their connections
// and returns a pointer to the antFarmRooms structure, which represents the whole ant farm.
func Rooms(roomsandConnections []string) *room {
	var antFarmRooms *room
	// Initialize the slices to store the names of the rooms, connections, beginning and destination rooms
	roomNames := []string{}
	connections := []string{}
	beginConnRmNames := []string{}
	destConnRmNames := []string{}
	
	// Loop through the roomsandConnections array to extract the names of the rooms
	for i := 0; i < len(roomsandConnections); i++ {
		for j := 0; j < len(roomsandConnections[i]); j++ {
			// If there's a space, split the string and store the name of the room in roomNames slice
			if roomsandConnections[i][j] == ' ' {
				roomName := strings.Split(roomsandConnections[i], " ")[0]
				roomNames = append(roomNames, roomName)

				break // because there are 2 spaces
			}
		}
	}
	// Store the name of the starting and ending rooms in startRmName and endRmName variables
	startRmName := nameofStart
	endRmName := nameofEnd
	
	// Loop through the roomsandConnections array to extract the connections between the rooms
	for i := 0; i < len(roomsandConnections); i++ {
		for j := 0; j < len(roomsandConnections[i]); j++ {
			if roomsandConnections[i][j] == '-' {
				connections = append(connections, roomsandConnections[i])
				
				beginDestSlice := strings.Split(roomsandConnections[i], "-")
				// If the beginning room is the ending room, add the destination room to beginConnRmNames
				if beginDestSlice[0] == endRmName {
					beginConnRmNames = append(beginConnRmNames, beginDestSlice[1])
					destConnRmNames = append(destConnRmNames, beginDestSlice[0])
				}
				// Add the beginning and destination rooms to their respective slices
				beginConnRmNames = append(beginConnRmNames, beginDestSlice[0])
				destConnRmNames = append(destConnRmNames, beginDestSlice[1])
			}
		}
	}
	// Loop through the beginConnRmNames slice to check if any room is linked to itself
	for i := 0; i < len(beginConnRmNames); i++ {
		if beginConnRmNames[i] == destConnRmNames[i] {
			fmt.Println("ERROR\nSome Rooms Linked to Themselves")
			os.Exit(0)

		}
	}
	// Add the start room to the antFarmRooms
	antFarmRooms = addRoom(antFarmRooms, startRmName, startRmName, endRmName, beginConnRmNames, destConnRmNames)
	return antFarmRooms
}

func findChildren(roomToAdd *room, rmToAddName string, startRmName, endRmName string, beginConnRmNames, destConnRmNames []string) {
	for c := 0; c < len(beginConnRmNames); c++ {
		beginRmName := beginConnRmNames[c]
		if startRmName == destConnRmNames[c] {
			beginConnRmNames[c], destConnRmNames[c] = destConnRmNames[c], beginConnRmNames[c]
		}
		if rmToAddName != startRmName {
			for m := 0; m < len(firstRm.children); m++ {
				if firstRm.children[m].name == destConnRmNames[c] && startRmName != beginConnRmNames[c] {
					beginConnRmNames[c], destConnRmNames[c] = destConnRmNames[c], beginConnRmNames[c]
				}
			}
		}
		// Check if the beginning room of the current connection is the room to add.
		if beginRmName == rmToAddName {
			// If the above condition is true, retrieve the name of the destination room and add a new room to the ant farm.
			destRmName := destConnRmNames[c]
			addRoom(roomToAdd, destRmName, startRmName, endRmName, beginConnRmNames, destConnRmNames)
		}
	}
}

func addRoom(root *room, rmToAddName string, startRmName, endRmName string, beginConnRmNames, destConnRmNames []string) *room {
	var roomToAdd *room
	var endroomparent bool = false
	var anotherbool bool = false
	checkdup = append(checkdup, rmToAddName)
	if rmToAddName == endRmName { // end room / base case
		test = nil
		test = append(test, root)
		if firstendrm == 0 {
			lastRm = &room{
				parent:   test,
				name:     rmToAddName,
				children: nil,
				occupied: false,
			}
			firstendrm = 1
		} else {
			parentnumber := len(lastRm.parent)
			for t := 0; t < parentnumber; t++ {
				if lastRm.parent[t].name == root.name {
					lastRm.parent[t].children = root.children
					lastRm.parent[t].parent = root.parent
					endroomparent = true
				} else if t == parentnumber-1 && !endroomparent {
					lastRm.parent = append(lastRm.parent, root)
				}
			}
		}

		countofparents := len(lastRm.parent)
		if !endroomparent {
			lastRm.parent[countofparents-1].children = append(lastRm.parent[countofparents-1].children, lastRm)
		}


		return lastRm

	} else if rmToAddName == startRmName { // start Room special case
		firstRm = &room{
			parent:   nil,
			name:     rmToAddName,
			occupied: true,
		}
		findChildren(firstRm, rmToAddName, startRmName, endRmName, beginConnRmNames, destConnRmNames)
		Farm = append(Farm, *firstRm)
	} else {

		test = nil
		test = append(test, root)
		if len(dup(checkdup)) != 0 {
			for i := 0; i < len(dup(checkdup)); i++ {
				if (dup(checkdup)[i]) == rmToAddName {
					for t := 0; t < len(Farm); t++ {
						if Farm[t].name == rmToAddName {
							if roomToAdd == nil {
								roomToAdd = &room{
									parent:   test,
									name:     rmToAddName,
									occupied: false,
								}
							}
							roomToAdd.parent = Farm[t].parent
							roomToAdd.children = Farm[t].children
							numberofparents := len(roomToAdd.parent)
							for u := 0; u < len(roomToAdd.parent); u++ {
								if roomToAdd.parent[u].name == root.name {
									roomToAdd.parent[u].parent = root.parent
									roomToAdd.parent[u].children = root.children
									anotherbool = true

								} else if u == numberofparents-1 && !anotherbool {
									roomToAdd.parent = append(roomToAdd.parent, root)
									anotherbool = true
									break

								}
							}

							break

						}
					}
					break
				} else {
					roomToAdd = &room{
						parent:   test,
						name:     rmToAddName,
						occupied: false,
					}
				}
			}
		} else {
			roomToAdd = &room{
				parent:   test,
				name:     rmToAddName,
				occupied: false,
			}
		}

		countofparents := len(roomToAdd.parent)
		countofchildrens := 0
		var somebool bool = true
		countofchildrens = len(roomToAdd.parent[countofparents-1].children)
		if len(roomToAdd.parent[countofparents-1].children) == 0 {
			roomToAdd.parent[countofparents-1].children = append(roomToAdd.parent[countofparents-1].children, roomToAdd)
		} else {
			for r := 0; r < countofchildrens; r++ {

				if roomToAdd.parent[countofparents-1].children[r].name == roomToAdd.name {
					roomToAdd.parent[countofparents-1].children[r] = roomToAdd
					somebool = false
				}
			}
			if somebool {
				roomToAdd.parent[countofparents-1].children = append(roomToAdd.parent[countofparents-1].children, roomToAdd)
			}

		}

		findChildren(roomToAdd, rmToAddName, startRmName, endRmName, beginConnRmNames, destConnRmNames)
		Farm = CheckFarmDup(Farm, roomToAdd.name)
		Farm = append(Farm, *roomToAdd)
	}

	return roomToAdd
}

func dup(s []string) []string {
	var result []string
	duplicate := make(map[string]bool)
	for i := 0; i < len(s); i++ {
		if duplicate[s[i]] {
			result = append(result, s[i])
		} else {
			duplicate[s[i]] = true
		}
	}
	return result
}

func RemoveElement(s []room, i int) []room {
	return append(s[:i], s[i+1:]...)
}

func CheckFarmDup(s []room, t string) []room {
	for k := 0; k < len(s); k++ {
		if s[k].name == t {
			s = RemoveElement(s, k)
			CheckFarmDup(s, t)
		}
	}

	return s
}
