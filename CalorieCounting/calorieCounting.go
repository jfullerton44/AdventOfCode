package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, err := os.Open("day1in.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	totals := make([]int, 0)
	curr := 0
	for scanner.Scan() {
		text := scanner.Text()
		intVar, err := strconv.Atoi(text)
		if err != nil {
			totals = append(totals, curr)
			curr = 0
		}
		curr += intVar
	}

	sort.IntSlice(totals).Sort()

	len := len(totals)

	top := totals[len-1]
	sum := totals[len-1] + totals[len-2] + totals[len-3]
	fmt.Printf("Top: %d \nTop 3: %d", top, sum)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
