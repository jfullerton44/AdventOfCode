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

	fmt.Printf("Round 2: %d\n", rd2)

}

func round1(locations []Pair) (int, error) {
	ref := Position{0, 2000000}
	total := 0
	for i := -10000000; i < 10000000; i++ {
		ref.x = i
		if !canContainBeacon(ref, locations) {
			total++
		}
	}
	return total, nil
}

func round2(locations []Pair) (int, error) {
	pos := Position{0, 0}
	for i := 0; i <= 4000000; i++ {
		for j := 0; j <= 4000000; j++ {
			pos.x = j
			pos.y = i
			if canContainBeacon(pos, locations) {
				return pos.x*4000000 + pos.y, nil
			} else {
				pair, err := coveredBySensor(pos, locations)
				if err != nil {
					return 0, err
				}

				dist := distance(pair.Sensor, pair.Beacon)

				j += dist - abs(pair.Sensor.y-pos.y) + abs(pair.Sensor.x-pos.x)

			}
		}
	}
	return 0, nil
}

type Position struct {
	x, y int
}

type Pair struct {
	Sensor, Beacon Position
}

func readInput() ([]Pair, error) {
	locations := make([]Pair, 0)

	f, err := os.Open("in.txt")
	if err != nil {
		return nil, err
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		split := strings.Split(text, ":")
		sensor, err := textToPosition(split[0])
		if err != nil {
			return nil, err
		}
		beacon, err := textToPosition(split[1])
		if err != nil {
			return nil, err
		}
		locations = append(locations, Pair{*sensor, *beacon})
	}
	return locations, nil
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

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func textToPosition(in string) (*Position, error) {
	xLocation := strings.Index(in, "x=")
	commaLocation := strings.Index(in, ",")
	xString := in[xLocation+2 : commaLocation]
	yString := in[commaLocation+4:]
	x, err := strconv.Atoi(xString)
	if err != nil {
		return nil, err
	}
	y, err := strconv.Atoi(yString)
	if err != nil {
		return nil, err
	}
	return &Position{x, y}, nil
}

func distance(a Position, b Position) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func canContainBeacon(a Position, locations []Pair) bool {
	for _, pair := range locations {
		if a == pair.Beacon {
			return true
		}
		if distance(a, pair.Sensor) <= distance(pair.Beacon, pair.Sensor) {
			return false
		}
	}
	return true
}

func coveredBySensor(a Position, locations []Pair) (Pair, error) {
	for _, pair := range locations {
		if distance(a, pair.Sensor) <= distance(pair.Beacon, pair.Sensor) {
			return pair, nil
		}
	}
	return Pair{}, fmt.Errorf("no sensor covers %v", a)
}
