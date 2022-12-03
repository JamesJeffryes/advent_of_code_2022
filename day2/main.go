package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	tStart := time.Now()
	rounds := getInput("input.txt")
	//log.Println(rounds)
	var totalValue int

	for _, throws := range rounds {
		pts := runRound(throws[0], throws[1])
		//log.Println(pts)
		totalValue += pts
	}
	log.Printf("Part 1: %d \n", totalValue)
	log.Printf("%s elapsed", time.Since(tStart))

	totalValue = 0
	for _, throws := range rounds {
		opts := getLoseDrawWin(throws[0])
		pts := runRound(throws[0], opts[int(throws[1]-'A')])
		//log.Println(pts)
		totalValue += pts
	}
	log.Printf("Part 2: %d \n", totalValue)
	log.Printf("%s elapsed", time.Since(tStart))
}

func getInput(filename string) (input [][2]rune) {
	replacer := strings.NewReplacer("X", "A", "Y", "B", "Z", "C")
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		runes := []rune(replacer.Replace(scanner.Text()))
		input = append(input, [2]rune{runes[0], runes[2]})
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return input
}

func getLoseDrawWin(they rune) [3]rune {
	switch they {
	case 'A':
		return [3]rune{'C', 'A', 'B'}
	case 'B':
		return [3]rune{'A', 'B', 'C'}
	case 'C':
		return [3]rune{'B', 'C', 'A'}
	}
	log.Fatal("Input error")
	return [3]rune{}
}

func runRound(they, you rune) int {
	for i, val := range getLoseDrawWin(they) {
		if you == val {
			return 3*i + int(you-'A'+1)
		}
	}
	log.Fatal("Input error")
	return 0
}
