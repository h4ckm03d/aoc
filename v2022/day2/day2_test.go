package day2

import (
	"log"
	"os"
	"testing"
)

func TestDay2_1(t *testing.T) {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	t.Log(Solve1(file))
}

func TestDay2_2(t *testing.T) {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	t.Log(Solve2(file))
}
