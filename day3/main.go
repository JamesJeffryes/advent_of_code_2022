package main

import (
	"bufio"
	"log"
	"os"
	"time"
)

func main() {
	tStart := time.Now()
	rucksacks := getInput("input.txt")
	var totalValue int
	for _, sack := range rucksacks {
		item := findCommonItemStr(sack)
		//log.Println(item, string(item))
		value := getValue(item)
		//log.Println(value)
		totalValue += value
	}
	log.Printf("Part 1: %d \n", totalValue)
	log.Printf("%s elapsed", time.Since(tStart))

	totalValue = 0
	for i := 0; i < len(rucksacks); i += 3 {
		item := findCommonItemSlice(rucksacks[i : i+3])
		//log.Println(item, string(item))
		value := getValue(item)
		//log.Println(value)
		totalValue += value
	}
	log.Printf("Part 2: %d \n", totalValue)
	log.Printf("%s elapsed", time.Since(tStart))
}

func getInput(filename string) (input []string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return input
}

func findCommonItemStr(items string) rune {
	splitInd := len(items) / 2
	compartmentOne := make(map[rune]struct{}, splitInd)
	for _, s := range items[:splitInd] {
		compartmentOne[s] = struct{}{}
	}
	for _, s := range items[splitInd:] {
		if _, ok := compartmentOne[s]; ok {
			return s
		}
	}
	return '0'
}

func findCommonItemSlice(sacks []string) rune {
	oldCommon := make(map[rune]struct{}, len(sacks[0]))
	for _, s := range sacks[0] {
		oldCommon[s] = struct{}{}
	}
	for _, items := range sacks[1:] {
		newCommon := make(map[rune]struct{})
		for _, s := range items {
			if _, ok := oldCommon[s]; ok {
				newCommon[s] = struct{}{}
			}
		}
		oldCommon = newCommon
	}
	for s := range oldCommon {
		return s
	}
	return '0'
}

func getValue(item rune) int {

	if item >= 'a' && item <= 'z' {
		return int(item-'a') + 1
	} else if item >= 'A' && item <= 'Z' {
		return int(item-'A') + 27
	} else {
		log.Fatalf("item %v has undefined value", item)
	}
	return 0
}
