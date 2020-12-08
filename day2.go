package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day2_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	valid1 := 0
	valid2 := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		entry := scanner.Text()
		entryParts := strings.Split(entry, " ")
		minMaxParts := strings.Split(entryParts[0], "-")
		min, err := strconv.Atoi(minMaxParts[0])
		if err != nil {
			log.Fatal(err)
		}
		max, err := strconv.Atoi(minMaxParts[1])
		if err != nil {
			log.Fatal(err)
		}
		char := string(entryParts[1][0])
		pwd := entryParts[2]

		if count := strings.Count(pwd, char); count >= min && count <= max {
			valid1++
		}
		if (string(pwd[min-1]) == char) != (string(pwd[max-1]) == char) {
			valid2++
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d valid passwords\n", valid1)
	fmt.Printf("Part 2: %d valid passwords\n", valid2)
}