package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"
)

func main() {
	tStart := time.Now()
	input := getInput("input.txt")
	//log.Println(input)
	totalValue := 0
	for _, indexes := range input {
		if indexes[0] >= indexes[2] && indexes[1] <= indexes[3] || indexes[0] <= indexes[2] && indexes[1] >= indexes[3] {
			totalValue += 1
		}
	}
	log.Printf("Part 1: %d \n", totalValue)
	log.Printf("%s elapsed", time.Since(tStart))

	totalValue = 0
	for _, indexes := range input {
		if indexes[1] < indexes[2] || indexes[0] > indexes[3] {
			continue
		}
		totalValue += 1
	}
	log.Printf("Part 2: %d \n", totalValue)
	log.Printf("%s elapsed", time.Since(tStart))
}

func getInput(filename string) (input [][4]int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	exp := regexp.MustCompile("\\d+")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		indexes := [4]int{}
		iStrs := exp.FindAllString(scanner.Text(), -1)
		for i, str := range iStrs {
			ind, err := strconv.Atoi(str)
			if err != nil {
				log.Fatalln(err)
			}
			indexes[i] = ind
		}
		input = append(input, indexes)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return input
}
