package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func getFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Error opening", filename)
	}

	return file
}

func filterFile(file *os.File) []string {
	arr := make([]string, 0)

	// part one regex
	r, _ := regexp.Compile(`mul\(\d+,\d+\)`)
	// part two regex
	// r, _ := regexp.Compile(`don't|do|mul\(\d+,\d+\)`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		filteredItems := r.FindAllString(scanner.Text(), -1)
		arr = append(arr, filteredItems...)
	}

	return arr
}

func multiplyInstructions(mulList *[]string) int {
	total := 0
	isEnabled := true

	for i := 0; i < len(*mulList); i++ {
		if (*mulList)[i] == "don't" {
			isEnabled = false
			continue
		} else if (*mulList)[i] == "do" {
			isEnabled = true
			continue
		}

		if isEnabled {
			numbers := strings.Split((*mulList)[i], ",")
			first, err := strconv.Atoi(numbers[0][4:])
			if err != nil {
				log.Fatal("Not a number", first)
			}
			second, err := strconv.Atoi(numbers[1][:len(numbers[1])-1])
			if err != nil {
				log.Fatal("Not a number", second)
			}

			product := first * second
			total += product
		}
	}

	return total
}

func main() {
	file := getFile("input.txt")
	multiplicationArr := filterFile(file)
	total := multiplyInstructions(&multiplicationArr)
	fmt.Println("Total:", total)
}
