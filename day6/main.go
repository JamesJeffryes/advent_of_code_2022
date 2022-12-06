package main

import (
	"io/ioutil"
	"log"
	"time"
)

func main() {
	tStart := time.Now()

	buffer, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Part 1: %d \n", startIndex(buffer, 4))
	log.Printf("%s elapsed", time.Since(tStart))

	log.Printf("Part 2: %d \n", startIndex(buffer, 14))
	log.Printf("%s elapsed", time.Since(tStart))
}

func startIndex(buffer []byte, nDistinct int) int {
	for i := nDistinct - 1; i < len(buffer); i++ {
		seen := make(map[byte]struct{})
		for j := 0; j < nDistinct; j++ {
			if _, ok := seen[buffer[i-j]]; ok {
				break // early termination
			}
			seen[buffer[i-j]] = struct{}{}
		}
		if len(seen) == nDistinct {
			return i + 1
		}
	}
	return 0
}
