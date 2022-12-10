package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	rd1, err := round1()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Round 1: %d\n", rd1)

	rd2, err := round2()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Round 2: %d", rd2)

}

func round1() (int, error) {
	f, err := os.Open("in.txt")
	locations := make(map[Position]bool)

	if err != nil {
		return 0, err
	}

	defer f.Close()
	head := &Position{
		X: 0,
		Y: 0,
	}
	tail := &Position{
		X: 0,
		Y: 0,
	}
	locations[*tail] = true
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		direction := string(text[0])
		distance, err := strconv.Atoi(text[2:])
		if err != nil {
			return 0, err
		}
		// fmt.Println(direction, distance)
		move(locations, head, tail, direction, int(distance))

	}
	total := 0
	for _, val := range locations {
		if val {
			total++
		}
	}
	return total, nil
}

func round2() (int, error) {
	f, err := os.Open("in.txt")
	locations := make(map[Position]bool)

	if err != nil {
		return 0, err
	}

	defer f.Close()
	head := &Position{
		X: 0,
		Y: 0,
	}
	tails := make([]*Position, 9)
	for i := range tails {
		tails[i] = &Position{}
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		direction := string(text[0])
		distance, err := strconv.Atoi(text[2:])
		if err != nil {
			return 0, err
		}
		// fmt.Println(direction, distance)
		moveSlice(locations, head, tails, direction, int(distance))

	}
	total := 0
	for _, val := range locations {
		if val {
			total++
		}
	}
	return total, nil
}

type Position struct {
	X, Y int
}

func move(locations map[Position]bool, head *Position, tail *Position, direction string, distance int) error {
	for ; distance > 0; distance-- {
		switch direction {
		case "U":
			head.Y++
		case "D":
			head.Y--
		case "L":
			head.X--
		case "R":
			head.X++
		default:
			return fmt.Errorf("bad direction %s", direction)
		}

		if touching(head, tail) {

		} else if head.X == tail.X {
			if head.Y > tail.Y {
				tail.Y++
			} else if head.Y < tail.Y {
				tail.Y--
			}
		} else if head.Y == tail.Y {
			if head.X > tail.X {
				tail.X++
			} else if head.X < tail.X {
				tail.X--
			}
		} else {
			if head.X > tail.X {
				tail.X++
			} else {
				tail.X--
			}

			if head.Y > tail.Y {
				tail.Y++
			} else {
				tail.Y--
			}
		}
		locations[*tail] = true
	}
	return nil
}

func moveSlice(locations map[Position]bool, leader *Position, tails []*Position, direction string, distance int) error {
	for ; distance > 0; distance-- {
		switch direction {
		case "U":
			leader.Y++
		case "D":
			leader.Y--
		case "L":
			leader.X--
		case "R":
			leader.X++
		default:
			return fmt.Errorf("bad direction %s", direction)
		}
		for i, tail := range tails {
			head := &Position{}
			if i == 0 {
				head = leader
			} else {
				head = tails[i-1]
			}
			if touching(head, tail) {
			} else if head.X == tail.X {
				if head.Y > tail.Y {
					tail.Y++
				} else if head.Y < tail.Y {
					tail.Y--
				}
			} else if head.Y == tail.Y {
				if head.X > tail.X {
					tail.X++
				} else if head.X < tail.X {
					tail.X--
				}
			} else {
				if head.X > tail.X {
					tail.X++
				} else {
					tail.X--
				}

				if head.Y > tail.Y {
					tail.Y++
				} else {
					tail.Y--
				}
			}
		}

		locations[*tails[8]] = true
	}
	return nil
}

func touching(head *Position, tail *Position) bool {
	if *head == *tail {
		return true
	}
	if head.X == tail.X {
		return math.Abs(float64(head.Y)-float64(tail.Y)) <= 1
	}
	if head.Y == tail.Y {
		return math.Abs(float64(head.X)-float64(tail.X)) <= 1
	}
	return math.Abs(float64(head.Y)-float64(tail.Y)) <= 1 && math.Abs(float64(head.X)-float64(tail.X)) <= 1
}
