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
	count := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		textSplit := strings.Index(text, ",")
		first := text[:textSplit]
		second := text[textSplit+1:]

		firstSplit := strings.Index(first, "-")
		first1, _ := strconv.Atoi(first[:firstSplit])
		first2, _ := strconv.Atoi(first[firstSplit+1:])

		secondSplit := strings.Index(second, "-")
		second1, _ := strconv.Atoi(second[:secondSplit])
		second2, _ := strconv.Atoi(second[secondSplit+1:])

		if first1 >= second1 && first1 <= second2 {
			count++
		} else if first2 >= second1 && first2 <= second2 {
			count++
		} else if first1 <= second1 && first2 >= second1 {
			count++
		} else if first1 <= second2 && first2 >= second2 {
			count++
		}
	}

	return count, nil
}

func round1() (int, error) {
	f, err := os.Open("in.txt")

	if err != nil {
		return 0, err
	}

	defer f.Close()
	count := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		textSplit := strings.Index(text, ",")
		first := text[:textSplit]
		second := text[textSplit+1:]

		firstSplit := strings.Index(first, "-")
		first1, _ := strconv.Atoi(first[:firstSplit])
		first2, _ := strconv.Atoi(first[firstSplit+1:])

		secondSplit := strings.Index(second, "-")
		second1, _ := strconv.Atoi(second[:secondSplit])
		second2, _ := strconv.Atoi(second[secondSplit+1:])

		if first1 >= second1 && first2 <= second2 {
			count++
		} else if first1 <= second1 && first2 >= second2 {
			count++
		}
	}

	return count, nil
}
