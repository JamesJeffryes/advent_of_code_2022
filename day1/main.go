package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"time"
)

func main() {
	tStart := time.Now()
	input := getInput("input.txt")
	//log.Println(input)
	totalValue := 0
	calSlice := make([]int, len(input))
	for i, cals := range input {
		for _, c := range cals {
			calSlice[i] += c
		}
		if calSlice[i] > totalValue {
			totalValue = calSlice[i]
		}
	}
	log.Printf("Part 1: %d \n", totalValue)
	log.Printf("%s elapsed", time.Since(tStart))

	totalValue = 0
	sort.Slice(calSlice, func(i, j int) bool {
		return calSlice[i] > calSlice[j]
	})
	for _, c := range calSlice[:3] {
		totalValue += c
	}
	log.Printf("Part 2: %d \n", totalValue)
	log.Printf("%s elapsed", time.Since(tStart))
}

func getInput(filename string) (input [][]int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	ints := make([]int, 0)
	for scanner.Scan() {
		str := scanner.Text()
		if str == "" {
			input = append(input, ints)
			ints = make([]int, 0)
			continue
		}
		val, err := strconv.Atoi(str)
		if err != nil {
			log.Fatalln(err)
		}
		ints = append(ints, val)
	}
	input = append(input, ints)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return input
}
