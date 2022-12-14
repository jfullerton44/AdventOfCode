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

func round1() (int, error) {
	f, err := os.Open("in.txt")

	if err != nil {
		return 0, err
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	pair := 1
	total := 0
	for scanner.Scan() {
		text1 := scanner.Text()
		scanner.Scan()
		text2 := scanner.Text()
		// text1 = strings.Replace(text1, "[[[]]]", fmt.Sprint(math.MinInt+2), -1)
		// text2 = strings.Replace(text2, "[[[]]]", fmt.Sprint(math.MinInt+2), -1)
		// text1 = strings.Replace(text1, "[[]]", fmt.Sprint(math.MinInt+1), -1)
		// text2 = strings.Replace(text2, "[[]]", fmt.Sprint(math.MinInt+1), -1)
		// text1 = strings.Replace(text1, "[]", fmt.Sprint(math.MinInt), -1)
		// text2 = strings.Replace(text2, "[]", fmt.Sprint(math.MinInt), -1)
		text1 = strings.Replace(text1, "[", ",-100", -1)
		text1 = strings.Replace(text1, "]", ",-50", -1)
		text2 = strings.Replace(text2, "[", ",-100", -1)
		text2 = strings.Replace(text2, "]", ",-50", -1)
		text1 = strings.Replace(text1, ",,", ",", -1)
		text2 = strings.Replace(text2, ",,", ",", -1)
		str1 := strings.Split(text1, ",")
		if str1[0] == "" {
			str1 = str1[1:]
		}
		list1 := make([]int, len(str1))

		for i, str := range str1 {
			list1[i], err = strconv.Atoi(str)
			if err != nil {
				fmt.Println(text1)
				fmt.Println(scanner.Text())
				fmt.Println(str1)
				return 0, err
			}
		}
		str2 := strings.Split(text2, ",")
		if str2[0] == "" {
			str2 = str2[1:]
		}
		list2 := make([]int, len(str2))
		for i, str := range str2 {
			list2[i], err = strconv.Atoi(str)
			if err != nil {
				fmt.Println(str2)
				return 0, err
			}
		}
		fmt.Println(list1)
		fmt.Println(list2)
		if isOrdered(list1, list2) {
			total += pair
			fmt.Println(pair)
		}
		fmt.Println()

		pair++
		scanner.Scan()
	}

	return total, nil
}

func round2() (int, error) {

	return 0, nil
}

func isOrdered(a []int, b []int) bool {
	aOpen := false
	bOpen := false
	j := 0
	for i := 0; i < len(a); i++ {
		elem := a[i]
		elemb := b[j]
		if elem == -100 {
			aOpen = true
			i++
			elem = a[i]
		}
		if elemb == -100 {
			bOpen = true
			j++
			elemb = b[j]
		}

		if i >= len(b) {
			return false
		}
		if elem > elemb {
			return false
		}
		if elem < elemb {
			return true
		}
	}
	return true
}
