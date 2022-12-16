package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type Valve struct {
	label     string
	flowRate  int
	pressure  int
	neighbors map[*Valve]int
}

func (v *Valve) addNeighbor(neighbor *Valve, flowRate int) {
	v.neighbors[neighbor] = flowRate
}

func createValves() (map[string]*Valve, error) {

	f, err := os.Open("in2.txt")
	if err != nil {
		return nil, err
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	valves := make(map[string]*Valve)

	for scanner.Scan() {
		line := scanner.Text()
		// Parse the valve information from the input line
		var label string
		var flowRate int
		fmt.Sscanf(line, "Valve %s has flow rate=%d;", &label, &flowRate)

		// Create a new valve and add it to the map
		valves[label] = &Valve{
			label:     label,
			flowRate:  flowRate,
			pressure:  0,
			neighbors: make(map[*Valve]int),
		}
	}

	f2, err := os.Open("in2.txt")
	if err != nil {
		return nil, err
	}

	defer f2.Close()
	scanner2 := bufio.NewScanner(f2)
	for scanner2.Scan() {
		line := scanner2.Text()
		// Add the tunnels to the graph
		var label string
		var flowRate int
		fmt.Sscanf(line, "Valve %s has flow rate=%d;", &label, &flowRate)
		tunStart := strings.Index(line, "tunnels lead to valves")
		if tunStart == -1 {
			tunStart = strings.Index(line, "tunnel leads to valve") - 1
		}
		tunnels := line[tunStart+23:]
		for _, tunnel := range strings.Split(tunnels, ", ") {
			neighbor := valves[tunnel]
			valves[label].addNeighbor(neighbor, neighbor.flowRate)
		}
	}

	return valves, nil
}

func dijkstra(valves map[string]*Valve, start string) {
	queue := make([]*Valve, 0)
	queue = append(queue, valves[start])

	for len(queue) > 0 {
		// Find the valve with the lowest pressure
		minPressure := math.MaxInt32
		minValve := queue[0]
		for _, valve := range queue {
			if valve.pressure < minPressure {
				minPressure = valve.pressure
				minValve = valve
			}
		}

		// Remove the valve from the queue and update its neighbors
		queue = remove(queue, minValve)
		for neighbor, flowRate := range minValve.neighbors {
			newPressure := minValve.pressure + flowRate
			if newPressure > neighbor.pressure {
				neighbor.pressure = newPressure
				queue = append(queue, neighbor)
			}
		}
	}
}

func remove(slice []*Valve, valve *Valve) []*Valve {
	for i, v := range slice {
		if v == valve {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

func main() {
	valves, err := createValves()
	if err != nil {
		panic(err)
	}
	fmt.Println(valves)
}
