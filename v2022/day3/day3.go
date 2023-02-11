package day3

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"
)

func Solve1(input io.Reader) int64 {
	sum := int64(0)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		txt := scanner.Text()
		priority := GetPriorities(txt)

		sum += getvalue(priority)
		// Lowercase item types a through z have priorities 1 through 26.
		// Uppercase item types A through Z have priorities 27 through 52.
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return sum
}

func getvalue(priority byte) int64 {
	if priority < 'a' {
		// fmt.Println("->", int64(priority-'a')+1)
		return int64(priority-'A') + 27
	}

	return int64(priority-'a') + 1
}

// GetPriorities is to get the same char in first half string with the second string
func GetPriorities(input string) byte {
	exists := map[byte]bool{}

	for i := 0; i < len(input)/2; i++ {
		if !exists[input[i]] {
			exists[input[i]] = true
		}
	}

	for i := len(input) / 2; i < len(input); i++ {
		if exists[input[i]] {
			return input[i]
		}
	}
	return 0
}

func GetPriorities2(inputs []string) byte {
	// exists := map[byte]int{}
	for i := range inputs[0] {
		if strings.ContainsAny(inputs[1], string(inputs[0][i])) && strings.ContainsAny(inputs[2], string(inputs[0][i])) {
			return inputs[0][i]
		}
	}

	return 0
}

func Solve2(input io.Reader) int64 {
	sum := int64(0)
	scanner := bufio.NewScanner(input)
	groups := make([]string, 3)
	counter := 0
	for scanner.Scan() {
		txt := scanner.Text()
		groups[counter] = txt
		counter++
		if counter == 3 {
			fmt.Println(groups)
			priority := GetPriorities2(groups)
			sum += getvalue(priority)
			counter = 0
			fmt.Println(string(priority), priority, getvalue(priority), sum)
		}
		// Lowercase item types a through z have priorities 1 through 26.
		// Uppercase item types A through Z have priorities 27 through 52.
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return sum
}
