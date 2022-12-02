package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func round2() (int, error) {
	f, err := os.Open("in.txt")

	if err != nil {
		return 0, err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	total := 0
	for scanner.Scan() {
		text := scanner.Text()
		switch text {
		case "A X":
			total += 3
		case "A Y":
			total += 4
		case "A Z":
			total += 8
		case "B X":
			total += 1
		case "B Y":
			total += 5
		case "B Z":
			total += 9
		case "C X":
			total += 2
		case "C Y":
			total += 6
		case "C Z":
			total += 7
		}
	}

	return total, nil
}

func round1() (int, error) {
	f, err := os.Open("in.txt")

	if err != nil {
		return 0, err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	total := 0
	for scanner.Scan() {
		text := scanner.Text()
		switch text {
		case "A X":
			total += 4
		case "A Y":
			total += 8
		case "A Z":
			total += 3
		case "B X":
			total += 1
		case "B Y":
			total += 5
		case "B Z":
			total += 9
		case "C X":
			total += 7
		case "C Y":
			total += 2
		case "C Z":
			total += 6
		}
	}

	return total, nil
}
