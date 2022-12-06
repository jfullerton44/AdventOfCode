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
	queue := make([]rune, 0)
	for scanner.Scan() {
		text := scanner.Text()
		for i, character := range text {
			if i < 100 {
				fmt.Println(i, queue)
			}
			queue = append(queue, character)
			if i > 13 {
				queue = queue[1:]
				if check(queue) {
					return i + 1, nil
				}
			}

		}
	}

	return 0, fmt.Errorf("not found")
}

func round1() (int, error) {
	f, err := os.Open("in.txt")

	if err != nil {
		return 0, err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	queue := make([]rune, 0)
	for scanner.Scan() {
		text := scanner.Text()
		for i, character := range text {
			if i < 100 {
				fmt.Println(i, queue)
			}
			queue = append(queue, character)
			if i > 3 {
				queue = queue[1:]
				if check(queue) {
					return i + 1, nil
				}
			}

		}
	}

	return 0, fmt.Errorf("not found")
}

func check(s []rune) bool {
	for i, v := range s {
		for j := i + 1; j < len(s); j++ {
			if s[j] == v {
				return false
			}
		}

	}

	return true
}
