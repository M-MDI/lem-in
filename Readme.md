# Lem-in Project

## Overview

Lem-in is a pathfinding project implemented in Go. The goal is to find the optimal path(s) for a colony of ants to traverse from a start room to an end room through a network of tunnels, minimizing the number of turns required.

## Algorithm

### 1. Input Parsing and Validation

```go
func parseInput(filename string) (antCount int, rooms map[string]Room, links []Link, start, end *Room, err error) {
    // Read file
    // Parse line by line
    // Validate data
    // Return parsed and validated data
}
```

### 2. Graph Representation

```go
type Room struct {
    Name string
    X, Y int
    Connections []*Room
}

type Graph struct {
    Rooms map[string]*Room
    Start, End *Room
}

func createGraph(rooms map[string]Room, links []Link) *Graph {
    // Create graph structure
    // Add rooms and connections
    // Return graph
}
```

### 3. Pathfinding Algorithm

```go
func findShortestPaths(graph *Graph) [][]Room {
    // Implement BFS
    // Find all shortest paths
    // Optimize path selection
    // Return optimal paths
}
```

### 4. Ant Movement Simulation

```go
type Ant struct {
    ID int
    CurrentRoom *Room
    Path []Room
}

func simulateAntMovement(paths [][]Room, antCount int) [][]string {
    // Initialize ants
    // Allocate ants to paths
    // Simulate moves
    // Return move history
}
```

### 5. Output Generation

```go
func generateOutput(input string, moves [][]string) string {
    // Format input echo
    // Format move output
    // Return formatted output
}
```

## Usage

```go
func main() {
    filename := os.Args[1]
    antCount, rooms, links, start, end, err := parseInput(filename)
    if err != nil {
        log.Fatalf("Error parsing input: %v", err)
    }

    graph := createGraph(rooms, links)
    paths := findShortestPaths(graph)
    moves := simulateAntMovement(paths, antCount)
    output := generateOutput(input, moves)

    fmt.Println(output)
}
```

## Error Handling

The program includes robust error handling for various scenarios:

- Invalid input format
- Missing start or end room
- No valid path from start to end
- Invalid room or link definitions

## Testing

Comprehensive unit tests are provided for each module:

```go
func TestParseInput(t *testing.T) { /* ... */ }
func TestCreateGraph(t *testing.T) { /* ... */ }
func TestFindShortestPaths(t *testing.T) { /* ... */ }
func TestSimulateAntMovement(t *testing.T) { /* ... */ }
```

## Bonus: Visualization

For bonus points, the program can output data in a format suitable for visualization:

```go
func generateVisualizationData(graph *Graph, moves [][]string) string {
    // Generate JSON or other formatted data for visualization
}
```

## Conclusion

This Lem-in implementation provides an efficient solution to the ant colony pathfinding problem, with clean, modular Go code and comprehensive error handling and testing.