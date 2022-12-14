package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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
	locations, _ = readInput()

	rd2, err := round2(locations)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Round 2: %d", rd2)

}

func round1(locations [][]bool) (int, error) {
	sandNum := 0
	for {
		sand := Position{500, 0}
		old := Position{0, 0}
		for sand != old {
			old = sand
			sand = drop(locations, sand)
			if sand.y > 990 {
				return sandNum, nil
			}
		}
		locations[sand.y][sand.x] = true
		sandNum++
	}
}

func round2(locations [][]bool) (int, error) {
	highestY := 0
	for i, row := range locations {
		for _, elem := range row {
			if elem {
				highestY = i
			}
		}
	}

	fill(locations, Position{0, highestY + 2}, Position{999, highestY + 2})
	sandNum := 1
	for {
		sand := Position{500, 0}
		old := Position{0, 0}
		for sand != old {
			old = sand
			sand = drop(locations, sand)
		}
		locations[sand.y][sand.x] = true
		if sand.x == 500 && sand.y == 0 {
			return sandNum, nil
		}
		sandNum++
	}
}

type Position struct {
	x, y int
}

func readInput() ([][]bool, error) {
	locations := make([][]bool, 1000)
	for i := range locations {
		locations[i] = make([]bool, 1000)
	}
	f, err := os.Open("in.txt")
	if err != nil {
		return nil, err
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		coordinates := strings.Split(text, " -> ")
		points := make([]Position, len(coordinates))
		for i, point := range coordinates {
			split := strings.Split(point, ",")
			x, err := strconv.Atoi(split[0])
			if err != nil {
				return nil, err
			}
			y, err := strconv.Atoi(split[1])
			if err != nil {
				return nil, err
			}
			pt := Position{x, y}
			points[i] = pt
		}

		for i := 0; i < len(points)-1; i++ {
			fill(locations, points[i], points[i+1])
		}
	}
	return locations, nil
}

func fill(locations [][]bool, start Position, end Position) {
	if start.x == end.x {
		for i := min(start.y, end.y); i <= max(start.y, end.y); i++ {
			locations[i][start.x] = true
		}
	} else {
		for i := min(start.x, end.x); i <= max(start.x, end.x); i++ {
			locations[start.y][i] = true
		}
	}
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func drop(locations [][]bool, curr Position) Position {
	if !locations[curr.y+1][curr.x] {
		curr.y++
		return curr
	}
	if !locations[curr.y+1][curr.x-1] {
		curr.y++
		curr.x--
		return curr
	}
	if !locations[curr.y+1][curr.x+1] {
		curr.y++
		curr.x++
		return curr
	}
	return curr
}
