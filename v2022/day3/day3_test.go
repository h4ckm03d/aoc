package day3

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
	res := Solve2(file)
	if res != 2817 {
		t.Errorf("should be 2817 got %d", res)
	}
}

func TestGetPriorities(t *testing.T) {
	res := GetPriorities("vJrwpWtwJgWrhcsFMMfFFhFp")
	t.Log(string(res), res, 'A', 'a')
	if res != 'p' {
		t.Errorf(`%v is not a valid priority`, res)
	}

	if res < 'a' {
		t.Log("->", int64(res-'A')+27)
	} else {
		t.Log("->", int64(res-'a')+1)
	}

	res = GetPriorities2([]string{"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg"})

	t.Log(string(res), res, getvalue(res))
	if res != 'r' {
		t.Errorf(`%v is not a valid priority`, res)
	}

	res = GetPriorities2([]string{"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw"})

	t.Log(string(res), res, getvalue(res), 'Z')
	if res != 'Z' && getvalue(res) != 52 {
		t.Errorf(`%v is not a valid priority %d`, res, getvalue(res))
	}
}
