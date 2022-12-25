package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	locations, err = readInput()
	if err != nil {
		log.Fatal(err)
	}

	rd2, err := round2(locations)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Round 2: %d\n", rd2)

}

func round1(locations map[Position]*Elf) (int, error) {
	moves := make([]move, 4)
	moves[0] = north
	moves[1] = south
	moves[2] = west
	moves[3] = east
	for i := 1; i <= 10; i++ {
		newLocations := make(map[Position]int)
		updating := make([]*Elf, 0)
		for _, elf := range locations {
			if isNear(*elf.location, locations) {
				updating = append(updating, elf)
			}
		}
		startingMove := i - 1%4
		for _, elf := range updating {
			for try := 0; try < 4; try++ {
				func1 := moves[(startingMove+try)%4]
				updatedElf, err := func1(*elf, locations)
				if err == nil {
					elf.newLocation = updatedElf.location
					newLocations[*updatedElf.location]++
					break
				}
				if try == 3 {
					elf.newLocation = elf.location
				}
			}
		}

		for _, elf := range updating {
			if newLocations[*elf.newLocation] == 1 {
				delete(locations, *elf.location)
				elf.location = elf.newLocation
				elf.newLocation = nil
				locations[*elf.location] = elf
			} else {
				elf.newLocation = nil
			}
		}
		fmt.Printf("After round %d\n", i)
		print(locations)
	}

	xMin := 99999999
	xMax := -99999999
	yMin := 99999999
	yMax := -99999999
	points := 0
	for _, elf := range locations {
		x := elf.location.X
		if x < xMin {
			xMin = x
		}
		if x > xMax {
			xMax = x
		}
		y := elf.location.Y

		if y < yMin {
			yMin = y
		}
		if y > yMax {
			yMax = y
		}
		points++
	}

	return (yMax-yMin+1)*(xMax-xMin+1) - points, nil
}

func readInput() (map[Position]*Elf, error) {
	f, err := os.Open("in.txt")
	locations := make(map[Position]*Elf)

	if err != nil {
		return nil, err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	y := 0
	for scanner.Scan() {
		text := scanner.Text()
		for x := range text {
			char := string(text[x])
			if char == "#" {
				pos := Position{
					X: x,
					Y: y,
				}
				elf := Elf{
					location: &pos,
				}
				locations[pos] = &elf
			}
		}
		y++
	}
	print(locations)
	return locations, nil
}

func isNear(loc Position, locations map[Position]*Elf) bool {
	tmp := Position{}
	for x := loc.X - 1; x <= loc.X+1; x++ {
		for y := loc.Y - 1; y <= loc.Y+1; y++ {
			if x != loc.X || y != loc.Y {
				tmp.X = x
				tmp.Y = y
				if locations[tmp] != nil {
					return true
				}
			}
		}
	}
	return false
}

func round2(locations map[Position]*Elf) (int, error) {
	moves := make([]move, 4)
	moves[0] = north
	moves[1] = south
	moves[2] = west
	moves[3] = east
	for i := 1; true; i++ {
		newLocations := make(map[Position]int)
		updating := make([]*Elf, 0)
		for _, elf := range locations {
			if isNear(*elf.location, locations) {
				updating = append(updating, elf)
			}
		}
		startingMove := i - 1%4
		if len(updating) == 0 {
			return i, nil
		}
		for _, elf := range updating {
			for try := 0; try < 4; try++ {
				func1 := moves[(startingMove+try)%4]
				updatedElf, err := func1(*elf, locations)
				if err == nil {
					elf.newLocation = updatedElf.location
					newLocations[*updatedElf.location]++
					break
				}
				if try == 3 {
					elf.newLocation = elf.location
				}
			}
		}

		for _, elf := range updating {
			if newLocations[*elf.newLocation] == 1 {
				delete(locations, *elf.location)
				elf.location = elf.newLocation
				elf.newLocation = nil
				locations[*elf.location] = elf
			} else {
				elf.newLocation = nil
			}
		}
		//fmt.Printf("After round %d\n", i)
		//print(locations)
	}

	return 0, fmt.Errorf("not found ")
}

type Position struct {
	X, Y int
}

type Elf struct {
	location    *Position
	newLocation *Position
}

type move func(elf Elf, locations map[Position]*Elf) (*Elf, error)

func north(elf Elf, locations map[Position]*Elf) (*Elf, error) {
	tmp := Position{
		X: elf.location.X,
		Y: elf.location.Y,
	}
	tmp.Y--
	if locations[tmp] != nil {
		return nil, fmt.Errorf("taken")
	}
	tmp.X++
	if locations[tmp] != nil {
		return nil, fmt.Errorf("taken")
	}
	tmp.X -= 2
	if locations[tmp] != nil {
		return nil, fmt.Errorf("taken")
	}
	tmp.X++

	elf.location = &tmp
	return &elf, nil
}

func south(elf Elf, locations map[Position]*Elf) (*Elf, error) {
	tmp := Position{
		X: elf.location.X,
		Y: elf.location.Y,
	}
	tmp.Y++
	if locations[tmp] != nil {
		return nil, fmt.Errorf("taken")
	}
	tmp.X++
	if locations[tmp] != nil {
		return nil, fmt.Errorf("taken")
	}
	tmp.X -= 2
	if locations[tmp] != nil {
		return nil, fmt.Errorf("taken")
	}
	tmp.X++

	elf.location = &tmp
	return &elf, nil
}

func east(elf Elf, locations map[Position]*Elf) (*Elf, error) {
	tmp := Position{
		X: elf.location.X,
		Y: elf.location.Y,
	}
	tmp.X++
	if locations[tmp] != nil {
		return nil, fmt.Errorf("taken")
	}
	tmp.Y++
	if locations[tmp] != nil {
		return nil, fmt.Errorf("taken")
	}
	tmp.Y -= 2
	if locations[tmp] != nil {
		return nil, fmt.Errorf("taken")
	}
	tmp.Y++

	elf.location = &tmp
	return &elf, nil
}

func west(elf Elf, locations map[Position]*Elf) (*Elf, error) {
	tmp := Position{
		X: elf.location.X,
		Y: elf.location.Y,
	}
	tmp.X--
	if locations[tmp] != nil {
		return nil, fmt.Errorf("taken")
	}
	tmp.Y++
	if locations[tmp] != nil {
		return nil, fmt.Errorf("taken")
	}
	tmp.Y -= 2
	if locations[tmp] != nil {
		return nil, fmt.Errorf("taken")
	}
	tmp.Y++

	elf.location = &tmp
	return &elf, nil
}

func print(locations map[Position]*Elf) {
	xMin := 99999999
	xMax := -99999999
	yMin := 99999999
	yMax := -99999999
	for _, elf := range locations {
		x := elf.location.X
		if x < xMin {
			xMin = x
		}
		if x > xMax {
			xMax = x
		}
		y := elf.location.Y

		if y < yMin {
			yMin = y
		}
		if y > yMax {
			yMax = y
		}
	}

	for y := yMin; y <= yMax; y++ {
		for x := xMin; x <= xMax; x++ {
			pos := Position{
				X: x,
				Y: y,
			}
			if locations[pos] != nil {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
}
