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

	err = round2()

	if err != nil {
		log.Fatal(err)
	}

}

func round1() (int, error) {
	f, err := os.Open("in.txt")

	if err != nil {
		return 0, err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	x := 1
	queue := make([]int, 0)
	total := 0
	cycle := 1
	for scanner.Scan() {
		text := scanner.Text()
		for true {
			if cycle%40 == 20 {
				total += x * cycle
			}
			if len(queue) == 1 {
				x += queue[0]
				queue = queue[:0]
				cycle++
				break
			} else if text[:4] == "noop" {
				cycle++
				break
			} else if text[:4] == "addx" {
				val, err := strconv.Atoi(text[5:])
				if err != nil {
					return 0, err
				}
				queue = append(queue, val)
				cycle++
			}
		}
	}

	return total, nil
}

func round2() error {
	f, err := os.Open("in.txt")

	if err != nil {
		return err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	x := 1
	queue := make([]int, 0)
	cycle := 1
	for scanner.Scan() {
		text := scanner.Text()
		for {
			if math.Abs(float64(cycle%40)-float64(x+1)) <= 1 {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
			if cycle%40 == 0 {
				fmt.Printf("\n")
			}
			if len(queue) == 1 {
				x += queue[0]
				queue = queue[:0]
				cycle++
				break
			} else if text[:4] == "noop" {
				cycle++
				break
			} else if text[:4] == "addx" {
				val, err := strconv.Atoi(text[5:])
				if err != nil {
					return err
				}
				queue = append(queue, val)
				cycle++
			}
		}
	}

	return nil
}
