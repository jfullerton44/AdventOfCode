package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

func round1() (int, error) {
	monkeys, err := createMonkeys()
	if err != nil {
		return 0, err
	}
	for i := 1; i <= 20; i++ {
		for _, monkey := range monkeys {
			for _, item := range monkey.items {
				item = monkey.operation(item) / 3
				if monkey.test(item) {
					monkeys[monkey.passMonkey].items = append(monkeys[monkey.passMonkey].items, item)
				} else {
					monkeys[monkey.failMonkey].items = append(monkeys[monkey.failMonkey].items, item)
				}
				monkey.inspectedItems++
			}
			monkey.items = make([]int, 0)
		}
		// fmt.Printf("Round %d\n", i)
		// for k, monkey := range monkeys {
		// 	fmt.Printf("Monkey %d:", k)
		// 	fmt.Println(monkey.items)
		// }
	}
	itemsTouched := make([]int, len(monkeys))

	for i, monkey := range monkeys {
		itemsTouched[i] = monkey.inspectedItems
	}
	sort.Ints(itemsTouched)
	return itemsTouched[len(itemsTouched)-1] * itemsTouched[len(itemsTouched)-2], nil
}

func round2() (int, error) {
	monkeys, err := createMonkeys()
	if err != nil {
		return 0, err
	}
	divisor := 1
	for _, monkey := range monkeys {
		divisor *= monkey.testDivisor
	}
	for i := 1; i <= 10000; i++ {
		for _, monkey := range monkeys {
			for _, item := range monkey.items {
				item = monkey.operation(item) % divisor
				if monkey.test(item) {
					monkeys[monkey.passMonkey].items = append(monkeys[monkey.passMonkey].items, item)
				} else {
					monkeys[monkey.failMonkey].items = append(monkeys[monkey.failMonkey].items, item)
				}
				monkey.inspectedItems++
			}
			monkey.items = make([]int, 0)
		}
		// fmt.Printf("Round %d\n", i)
		// for k, monkey := range monkeys {
		// 	fmt.Printf("Monkey %d:", k)
		// 	fmt.Println(monkey.items)
		// }
	}
	itemsTouched := make([]int, len(monkeys))

	for i, monkey := range monkeys {
		itemsTouched[i] = monkey.inspectedItems
	}
	sort.Ints(itemsTouched)
	return itemsTouched[len(itemsTouched)-1] * itemsTouched[len(itemsTouched)-2], nil
}

func createMonkeys() ([]*Monkey, error) {
	monkeys := make([]*Monkey, 0)
	f, err := os.Open("in.txt")

	if err != nil {
		return nil, err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		monkey := &Monkey{}
		// Read in Starting Items
		scanner.Scan()
		text := scanner.Text()
		list := strings.Split(text[18:], ", ")
		intList := make([]int, len(list))
		for i, str := range list {
			intVal, err := strconv.Atoi(str)
			if err != nil {
				return nil, err
			}
			intList[i] = intVal
		}
		monkey.items = intList

		// Read in Operation
		scanner.Scan()
		text = scanner.Text()
		operation := strings.Split(text[23:], " ")
		if operation[0] == "+" {
			op := func(in int) int {
				val, err := strconv.Atoi(operation[1])
				if err != nil {
					return (in + in)
				}
				return (in + val)
			}
			monkey.operation = op
		} else if operation[0] == "*" {
			op := func(in int) int {
				val, err := strconv.Atoi(operation[1])
				if err != nil {
					return in * in
				}
				return in * val
			}
			monkey.operation = op
		}

		//Read in Test
		scanner.Scan()
		text = scanner.Text()
		testVal := text[21:]

		intVal, err := strconv.Atoi(testVal)
		if err != nil {
			return nil, err
		}
		test := func(in int) bool {
			return in%intVal == 0
		}
		monkey.test = test
		monkey.testDivisor = intVal
		monkeys = append(monkeys, monkey)

		// Read in pass monkey
		scanner.Scan()
		text = scanner.Text()
		passVal := text[29:]
		intMonkey, err := strconv.Atoi(passVal)
		if err != nil {
			return nil, err
		}
		monkey.passMonkey = intMonkey

		// Read in fail monkey
		scanner.Scan()
		text = scanner.Text()
		failVal := text[30:]
		intMonkey, err = strconv.Atoi(failVal)
		if err != nil {
			return nil, err
		}
		monkey.failMonkey = intMonkey
		scanner.Scan()
	}
	return monkeys, nil
}

type Monkey struct {
	items          []int
	operation      Operation
	test           Test
	passMonkey     int
	failMonkey     int
	inspectedItems int
	testDivisor    int
}

type Operation func(int) int

type Test func(int) bool
