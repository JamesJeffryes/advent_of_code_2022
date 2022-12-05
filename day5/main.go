package main

import (
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	tStart := time.Now()
	stacks, moves := getInput("input.txt")
	moveCratesOne(stacks, moves)
	log.Printf("Part 1: %s \n", getTopCrate(stacks))
	log.Printf("%s elapsed", time.Since(tStart))

	stacks, moves = getInput("input.txt")
	moveCratesAll(stacks, moves)
	log.Printf("Part 2: %s \n", getTopCrate(stacks))
	log.Printf("%s elapsed", time.Since(tStart))
}

func getTopCrate(outStacks [][]string) string {
	topCrate := strings.Builder{}
	for _, stack := range outStacks {
		topCrate.WriteString(stack[len(stack)-1])
	}
	return topCrate.String()
}

func getInput(filename string) ([][]string, [][3]int) {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	spContents := strings.SplitN(string(contents), "\n\n", 2)

	cargoLines := strings.Split(spContents[0], "\n")
	digitRE := regexp.MustCompile("\\d+")
	nStacks := len(digitRE.FindAllString(cargoLines[len(cargoLines)-1], -1))
	stacks := make([][]string, nStacks)
	for i := 0; i < len(cargoLines)-1; i++ {
		for j := 0; j < nStacks; j++ {
			container := string(cargoLines[len(cargoLines)-i-2][j*4+1])
			if container != " " {
				stacks[j] = append(stacks[j], container)
			}
		}

	}

	moves := make([][3]int, 0)
	for _, move := range strings.Split(spContents[1], "\n") {
		m := [3]int{}
		for i, d := range digitRE.FindAllString(move, 3) {
			val, err := strconv.Atoi(d)
			if err != nil {
				log.Fatal(err)
			}
			m[i] = val
		}
		moves = append(moves, m)

	}
	return stacks, moves
}

func moveCratesOne(stacks [][]string, moves [][3]int) {
	//NB: this mutates stacks becuase I am lazy
	for _, m := range moves {
		for i := 0; i < m[0]; i++ {
			from := m[1] - 1
			to := m[2] - 1
			stacks[to] = append(stacks[to], stacks[from][len(stacks[from])-1])
			stacks[from] = stacks[from][:len(stacks[from])-1]
		}
	}
}

func moveCratesAll(stacks [][]string, moves [][3]int) {
	//NB: this mutates stacks becuase I am lazy
	for _, m := range moves {
		from := m[1] - 1
		to := m[2] - 1
		stacks[to] = append(stacks[to], stacks[from][len(stacks[from])-m[0]:]...)
		stacks[from] = stacks[from][:len(stacks[from])-m[0]]
	}
}
