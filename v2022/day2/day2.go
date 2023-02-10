package day2

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"
)

type Shape int

const (
	Rock Shape = iota + 1
	Paper
	Scissor
)

var Mapp = map[string]Shape{
	"X": Rock,
	"Y": Paper,
	"Z": Scissor,
	"A": Rock,
	"B": Paper,
	"C": Scissor,
}

var Task2 = map[string]string{
	"X": "lose",
	"Y": "draw",
	"Z": "win",
}

func (s Shape) Duel(v Shape) int {
	if (s == Rock && v == Scissor) || (s == Paper && v == Rock) || (s == Scissor && v == Paper) {
		return 6
	}

	if s == v {
		return 3
	}

	return 0
}

func (s Shape) LoseWith() (v Shape) {
	switch s {
	case Rock:
		v = Paper
	case Scissor:
		v = Rock
	case Paper:
		v = Scissor
	}

	return
}

func (s Shape) WinWith() (v Shape) {
	switch s {
	case Rock:
		v = Scissor
	case Scissor:
		v = Paper
	case Paper:
		v = Rock
	}

	return
}

func Solve1(input io.Reader) int64 {
	sum := int64(0)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		txt := scanner.Text()
		ch := strings.Split(txt, " ")

		sum += int64(Mapp[ch[1]].Duel(Mapp[ch[0]]) + int(Mapp[ch[1]]))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return sum
}

func Solve2(input io.Reader) int64 {
	sum := int64(0)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		txt := scanner.Text()
		ch := strings.Split(txt, " ")
		var selected Shape
		switch Task2[ch[1]] {
		case "lose":
			selected = Mapp[ch[0]].WinWith()
		case "win":
			selected = Mapp[ch[0]].LoseWith()
		case "draw":
			selected = Mapp[ch[0]]
		}
		fmt.Println(ch, Mapp[ch[0]].Duel(selected), int(Mapp[ch[0]]), selected)
		sum += int64(selected.Duel(Mapp[ch[0]]) + int(selected))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return sum
}
