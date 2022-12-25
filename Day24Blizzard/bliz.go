package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	locations, err := readInput()
	if err != nil {
		log.Fatal(err)
	}

	rd1, err := round1(locations)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Round 1: %d\n", rd1)

	// locations, err = readInput()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	rd2, err := round2()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Round 2: %d\n", rd2)

}

func round1(locations *state) (int, error) {
	queue := make([]state, 1)
	queue[0] = *locations
	out := 0
	var err error
	for len(queue) != 0 {
		curr := queue[0]
		queue = queue[1:]
		fmt.Printf("Starting with \n")
		print(curr)
		out, queue, err = operate(curr, queue)
		if out != 0 {
			return out, nil
		}
	}
	return 0, err
}

func round2() (int, error) {
	return 0, nil
}

type Position struct {
	X, Y int
}

type location struct {
	points []*point
}

type point struct {
	location    *Position
	newLocation *Position
	direction   string
	isExplorer  bool
	isEnd       bool
}

type state struct {
	locations map[Position]*location
	moves     int
	xMax      int
	yMax      int
	xEnd      int
	yEnd      int
}

func print(curr state) {
	return
	for y := 0; y <= curr.yMax; y++ {
		for x := 0; x <= curr.xMax; x++ {
			pos := Position{
				X: x,
				Y: y,
			}
			if curr.locations[pos] == nil {
				fmt.Printf(".")
			} else {
				loc := curr.locations[pos]
				if len(loc.points) > 1 {
					fmt.Printf("%d", len(loc.points))
				} else {
					pos := loc.points[0]
					if pos.isExplorer {
						fmt.Printf("E")
					} else if pos.isEnd {
						fmt.Printf("F")
					} else {
						fmt.Printf(pos.direction)
					}
				}
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func operate(curr state, queue []state) (int, []state, error) {
	explorer := &point{}
	for _, loc := range curr.locations {
		for _, pt := range loc.points {
			if pt.isExplorer {
				explorer = pt
			} else {
				pt.newLocation = &Position{
					X: pt.location.X,
					Y: pt.location.Y,
				}
				if pt.direction == ">" {
					pt.newLocation.X += 1
					if pt.newLocation.X == curr.xMax {
						pt.newLocation.X = 1
					}
				} else if pt.direction == "<" {
					pt.newLocation.X--
					if pt.newLocation.X == 0 {
						pt.newLocation.X = curr.xMax - 1
					}
				} else if pt.direction == "^" {
					pt.newLocation.Y--
					if pt.newLocation.Y == 0 {
						pt.newLocation.Y = curr.yMax - 1
					}
				} else if pt.direction == "v" {
					pt.newLocation.Y++
					if pt.newLocation.Y == curr.yMax {
						pt.newLocation.Y = 1
					}
				}
			}
		}
	}
	new := state{
		locations: make(map[Position]*location),
		xMax:      curr.xMax,
		yMax:      curr.yMax,
		xEnd:      curr.xEnd,
		yEnd:      curr.yEnd,
		moves:     curr.moves + 1,
	}

	for _, loc := range curr.locations {
		for _, pt := range loc.points {
			if pt.newLocation != nil {
				if newLoc := new.locations[*pt.newLocation]; newLoc != nil {
					ptClone := pt.clone()
					newLoc.points = append(newLoc.points, ptClone)
					ptClone.location = ptClone.newLocation
					ptClone.newLocation = nil
				} else {
					ptClone := pt.clone()
					ptClone.location = ptClone.newLocation
					ptClone.newLocation = nil
					points := []*point{ptClone}
					newLocation := location{
						points: points,
					}
					new.locations[*ptClone.location] = &newLocation
				}
			}
		}
	}
	if explorer.location.X == curr.xEnd {
		if (explorer.location.Y + 1) == curr.yEnd {
			return curr.moves + 1, nil, nil
		}
	}
	fmt.Printf("updated to\n")
	fmt.Printf("(%d, %d) %d\n", explorer.location.X, explorer.location.Y, curr.moves)
	print(new)

	if isAvailable(new, explorer.location.X, explorer.location.Y) {
		newExplorer := explorer.clone()
		newState := generateState(curr)
		newState.locations[*newExplorer.location] = &location{[]*point{newExplorer}}
		print(newState)
		if !positionExists(*newExplorer, queue, new) {
			queue = append(queue, newState)
		}

	}
	if isAvailable(new, explorer.location.X+1, explorer.location.Y) {
		newExplorer := explorer.clone()
		newExplorer.location.X = newExplorer.location.X + 1
		newState := generateState(curr)
		newState.locations[*newExplorer.location] = &location{[]*point{newExplorer}}
		print(newState)
		if !positionExists(*newExplorer, queue, new) {
			queue = append(queue, newState)
		}
	}
	if isAvailable(new, explorer.location.X-1, explorer.location.Y) {
		newExplorer := explorer.clone()
		newExplorer.location.X = newExplorer.location.X - 1
		newState := generateState(curr)
		newState.locations[*newExplorer.location] = &location{[]*point{newExplorer}}
		print(newState)

		if !positionExists(*newExplorer, queue, new) {
			queue = append(queue, newState)
		}
	}
	if isAvailable(new, explorer.location.X, explorer.location.Y+1) {
		newExplorer := explorer.clone()
		newExplorer.location.Y = newExplorer.location.Y + 1
		newState := generateState(curr)
		newState.locations[*newExplorer.location] = &location{[]*point{newExplorer}}
		print(newState)
		if !positionExists(*newExplorer, queue, new) {
			queue = append(queue, newState)
		}
	}
	if isAvailable(new, explorer.location.X, explorer.location.Y-1) {
		newExplorer := explorer.clone()
		newExplorer.location.Y = newExplorer.location.Y - 1
		newState := generateState(curr)
		newState.locations[*newExplorer.location] = &location{[]*point{newExplorer}}
		print(newState)
		if !positionExists(*newExplorer, queue, new) {
			queue = append(queue, newState)
		}
	}

	return 0, queue, nil
}

func isAvailable(curr state, x int, y int) bool {
	temp := Position{
		X: x,
		Y: y,
	}
	if x == 1 && y == 0 {
		return true
	}

	if curr.locations[temp] != nil && len(curr.locations[temp].points) == 1 && curr.locations[temp].points[0].isExplorer {
		return true
	}

	if curr.locations[temp] != nil {
		return false
	}

	if x <= 0 || y <= 0 || x == curr.xMax || y == curr.yMax {
		return false
	}
	return true
}

func generateState(curr state) state {
	new := state{
		locations: make(map[Position]*location),
		xMax:      curr.xMax,
		yMax:      curr.yMax,
		xEnd:      curr.xEnd,
		yEnd:      curr.yEnd,
		moves:     curr.moves + 1,
	}

	for _, loc := range curr.locations {
		for _, pt := range loc.points {
			if pt.newLocation != nil {
				if newLoc := new.locations[*pt.newLocation]; newLoc != nil {
					ptClone := pt.clone()
					newLoc.points = append(newLoc.points, ptClone)
					ptClone.location = ptClone.newLocation
					ptClone.newLocation = nil
				} else {
					ptClone := pt.clone()
					ptClone.location = ptClone.newLocation
					ptClone.newLocation = nil
					points := []*point{ptClone}
					newLocation := location{
						points: points,
					}
					new.locations[*ptClone.location] = &newLocation
				}
			}

		}
	}
	return new
}

func (p *point) clone() *point {
	new := &point{
		direction:  p.direction,
		isExplorer: p.isExplorer,
		isEnd:      p.isEnd,
	}
	pos := Position{
		X: p.location.X,
		Y: p.location.Y,
	}
	if p.newLocation != nil {
		newPos := Position{
			X: p.newLocation.X,
			Y: p.newLocation.Y,
		}
		new.newLocation = &newPos

	}

	new.location = &pos
	return new
}

func readInput() (*state, error) {
	f, err := os.Open("in1.txt")
	start := state{
		locations: make(map[Position]*location),
	}

	if err != nil {
		return nil, err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	text := scanner.Text()
	startX := strings.Index(text, ".")
	position := Position{
		Y: 0,
		X: startX,
	}

	startPoint := &point{
		location:   &position,
		isExplorer: true,
	}
	loc := &location{
		[]*point{startPoint},
	}
	start.locations[position] = loc
	y := 1
	last := ""
	for scanner.Scan() {
		text := scanner.Text()
		last = text
		for x := range text {
			char := string(text[x])
			if char == ">" || char == "<" || char == "v" || char == "^" {
				pos := Position{
					X: x,
					Y: y,
				}
				pt := &point{
					location:   &pos,
					isExplorer: false,
					direction:  char,
				}
				loc := &location{
					[]*point{pt},
				}
				start.locations[pos] = loc
			}
			if x > start.xMax {
				start.xMax = x
			}
		}
		y++
		if y > start.yMax {
			start.yMax = y
		}
	}
	endX := strings.Index(last, ".")
	endPosition := Position{
		Y: y,
		X: endX,
	}

	endPoint := &point{
		location: &endPosition,
		isEnd:    true,
	}
	loc = &location{
		[]*point{endPoint},
	}
	start.locations[endPosition] = loc
	start.xEnd = endX
	start.yEnd = y
	print(start)
	return &start, nil
}

func positionExists(explorer point, queue []state, curr state) bool {
	for _, st := range queue {
		if val := st.locations[*explorer.location]; val != nil && len(val.points) == 1 && val.points[0].isExplorer && curr.moves == st.moves {
			return true
		}
	}
	return false
}
