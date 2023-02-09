package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Println(solution2(file))
}

func solution1(input io.Reader) int64 {
	scanner := bufio.NewScanner(input)
	max := int64(-99)
	temp := int64(0)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		txt := scanner.Text()
		if len(txt) == 0 {
			if temp > max {
				max = temp
			}
			temp = 0
		}
		val, _ := strconv.Atoi(txt)
		temp += int64(val)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return max
}

func solution2(input io.Reader) int64 {
	scanner := bufio.NewScanner(input)
	max := int64(-99)
	temp := int64(0)
	rank := make([]int64, 0)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		txt := scanner.Text()
		if len(txt) == 0 {
			if temp > max {
				max = temp
				rank = append([]int64{max}, rank...)
			}
			temp = 0
		}
		val, _ := strconv.Atoi(txt)
		temp += int64(val)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(rank)

	return rank[0] + rank[1] + rank[2]
}
