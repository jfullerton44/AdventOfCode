package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func part2() int {
	f, err := os.Open("in.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	total := 0
	num := 0
	map1 := make(map[string]bool)
	map2 := make(map[string]bool)
	for scanner.Scan() {
		if num == 0 {
			map1 = make(map[string]bool)
			map2 = make(map[string]bool)
		}
		text := scanner.Text()
		for i := 0; i < len(text); i++ {
			char := string(text[i])
			if num == 0 {
				map1[char] = true
			} else if num == 1 {
				if map1[char] {
					map2[char] = true
				}
			} else {
				if map2[char] {
					val := char[0]
					if val > 96 {
						val = val - 96
					}
					if val > 64 {
						val = val - 38
					}
					total += int(val)
					break
				}
			}
		}

		num++
		if num == 3 {
			num = 0
		}
	}
	return total
}

func part1() int {
	f, err := os.Open("in.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	total := 0
	for scanner.Scan() {
		text := scanner.Text()
		set := make(map[string]bool)
		for i := 0; i < len(text)/2; i++ {
			char := string(text[i])
			set[char] = true
		}
		for i := len(text) / 2; i < len(text); i++ {
			char := string(text[i])
			if set[char] {
				val := char[0]
				if val > 96 {
					val = val - 96
				}
				if val > 64 {
					val = val - 38
				}
				fmt.Printf("%s %d\n", char, int(val))
				total += int(val)
				break
			}

		}
	}
	return total
}
