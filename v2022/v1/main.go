package main

import (
	"bufio"
	"fmt"
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

	scanner := bufio.NewScanner(file)
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
	fmt.Println(max)
}
