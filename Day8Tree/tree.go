package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	trees := make([][]int, 0)
	f, err := os.Open("in.txt")

	if err != nil {
		panic(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	i := 0
	for scanner.Scan() {
		text := scanner.Text()
		trees = append(trees, make([]int, 0))
		for _, char := range text {
			val := int(char - '0')
			trees[i] = append(trees[i], val)
		}
		i++
	}
	rd1, err := round1(trees)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Round 1: %d\n", rd1)

	rd2, err := round2(trees)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Round 2: %d", rd2)

}

func round2(trees [][]int) (int, error) {
	total := 0
	for y := range trees {
		for x := range trees[y] {
			tmp := score(trees, x, y)
			if tmp > total {
				total = tmp
			}
		}
	}
	return total, nil
}

func round1(trees [][]int) (int, error) {
	total := 0
	for y := range trees {
		for x := range trees[y] {
			if check(trees, x, y) {
				total++
			}
		}
	}
	return total, nil
}

func check(trees [][]int, x int, y int) bool {
	north := true
	south := true
	east := true
	west := true
	for i := x - 1; i >= 0; i-- {
		if trees[y][i] >= trees[y][x] {
			west = false
		}
	}
	for i := y - 1; i >= 0; i-- {
		if trees[i][x] >= trees[y][x] {
			north = false
		}
	}
	for i := x + 1; i < len(trees[y]); i++ {
		if trees[y][i] >= trees[y][x] {
			east = false
		}
	}
	for i := y + 1; i < len(trees); i++ {
		if trees[i][x] >= trees[y][x] {
			south = false
		}
	}
	return north || south || east || west
}

func score(trees [][]int, x int, y int) int {
	north := 0
	south := 0
	east := 0
	west := 0
	for i := x - 1; i >= 0; i-- {
		if trees[y][i] >= trees[y][x] {
			west = x - i
			break
		}
		if i == 0 {
			west = x
			break
		}
	}
	for i := y - 1; i >= 0; i-- {
		if trees[i][x] >= trees[y][x] {
			north = y - i
			break
		}
		if i == 0 {
			north = y
			break
		}
	}
	for i := x + 1; i < len(trees[y]); i++ {
		if trees[y][i] >= trees[y][x] {
			east = i - x
			break
		}
		if i == len(trees[y])-1 {
			east = len(trees[y]) - x - 1
			break
		}
	}
	for i := y + 1; i < len(trees); i++ {
		if trees[i][x] >= trees[y][x] {
			south = i - y
			break
		}
		if i == len(trees)-1 {
			south = len(trees) - y - 1
			break
		}
	}
	//fmt.Println(north, east, south, west)
	return north * south * east * west
}
