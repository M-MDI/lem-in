package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Room struct {
	Name string
	X, Y int
	Start, End bool
}

type Tunnel struct {
	From, To string
}

type AntFarm struct {
	NumAnts int
	Rooms   map[string]*Room
	Tunnels []Tunnel
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <filename>")
		return
	}

	farm, err := parseFile(os.Args[1])
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return
	}

	// TODO: Implement path finding and ant movement simulation

	// For now, just print the parsed data
	printFarm(farm)
}

func parseFile(filename string) (*AntFarm, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("invalid data format, cannot open file: %v", err)
	}
	defer file.Close()

	farm := &AntFarm{
		Rooms: make(map[string]*Room),
	}

	scanner := bufio.NewScanner(file)
	lineNum := 0
	parsingRooms := true

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		lineNum++

		if lineNum == 1 {
			if err := parseNumAnts(farm, line); err != nil {
				return nil, err
			}
			continue
		}

		if line == "" || strings.HasPrefix(line, "#") {
			if line == "##start" || line == "##end" {
				if err := parseSpecialRoom(farm, scanner, line == "##start"); err != nil {
					return nil, err
				}
				lineNum++
			}
			continue
		}

		if parsingRooms {
			if strings.Contains(line, "-") {
				parsingRooms = false
			} else {
				if err := parseRoom(farm, line); err != nil {
					return nil, err
				}
				continue
			}
		}

		if err := parseTunnel(farm, line); err != nil {
			return nil, err
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("invalid data format, error reading file: %v", err)
	}

	return farm, nil
}

func parseNumAnts(farm *AntFarm, line string) error {
	num, err := strconv.Atoi(line)
	if err != nil || num <= 0 {
		return fmt.Errorf("invalid data format, invalid number of ants")
	}
	farm.NumAnts = num
	return nil
}

func parseSpecialRoom(farm *AntFarm, scanner *bufio.Scanner, isStart bool) error {
	if !scanner.Scan() {
		return fmt.Errorf("invalid data format, missing room after ##start or ##end")
	}
	line := strings.TrimSpace(scanner.Text())
	if err := parseRoom(farm, line); err != nil {
		return err
	}
	parts := strings.Fields(line)
	room := farm.Rooms[parts[0]]
	if isStart {
		room.Start = true
	} else {
		room.End = true
	}
	return nil
}

func parseRoom(farm *AntFarm, line string) error {
	parts := strings.Fields(line)
	if len(parts) != 3 {
		return fmt.Errorf("invalid data format, invalid room format")
	}
	name := parts[0]
	x, err := strconv.Atoi(parts[1])
	if err != nil {
		return fmt.Errorf("invalid data format, invalid room coordinate")
	}
	y, err := strconv.Atoi(parts[2])
	if err != nil {
		return fmt.Errorf("invalid data format, invalid room coordinate")
	}
	if strings.HasPrefix(name, "L") || strings.HasPrefix(name, "#") {
		return fmt.Errorf("invalid data format, invalid room name")
	}
	farm.Rooms[name] = &Room{Name: name, X: x, Y: y}
	return nil
}

func parseTunnel(farm *AntFarm, line string) error {
	parts := strings.Split(line, "-")
	if len(parts) != 2 {
		return fmt.Errorf("invalid data format, invalid tunnel format")
	}
	from, to := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
	if _, exists := farm.Rooms[from]; !exists {
		return fmt.Errorf("invalid data format, unknown room in tunnel: %s", from)
	}
	if _, exists := farm.Rooms[to]; !exists {
		return fmt.Errorf("invalid data format, unknown room in tunnel: %s", to)
	}
	farm.Tunnels = append(farm.Tunnels, Tunnel{From: from, To: to})
	return nil
}

func printFarm(farm *AntFarm) {
	fmt.Println(farm.NumAnts)
	for _, room := range farm.Rooms {
		prefix := ""
		if room.Start {
			prefix = "##start\n"
		} else if room.End {
			prefix = "##end\n"
		}
		fmt.Printf("%s%s %d %d\n", prefix, room.Name, room.X, room.Y)
	}
	for _, tunnel := range farm.Tunnels {
		fmt.Printf("%s-%s\n", tunnel.From, tunnel.To)
	}
}