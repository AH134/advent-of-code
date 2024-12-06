package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	return file
}

// PART ONE

func appendIDs(file *os.File, leftID *[]int, rightID *[]int) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		leftNum := strings.Fields(scanner.Text())[0]
		rightNum := strings.Fields(scanner.Text())[1]

		l, err := strconv.Atoi(leftNum)
		if err != nil {
			log.Fatal(err)
		}
		*leftID = append(*leftID, l)

		r, err := strconv.Atoi(rightNum)
		if err != nil {
			log.Fatal(err)
		}
		*rightID = append(*rightID, r)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func getTotalDistance(leftID *[]int, rightID *[]int) int {
	totalDistance := 0
	totalNumbers := len(*leftID)
	for i := 0; i < totalNumbers; i++ {
		distance := absDiff((*leftID)[i], (*rightID)[i])
		totalDistance += distance
	}

	return totalDistance
}

func absDiff(x, y int) int {
	if x < y {
		return y - x
	}

	return x - y
}

func partOne() {
	var rightID []int
	var leftID []int

	file := readFile("./input.txt")
	defer file.Close()

	appendIDs(file, &leftID, &rightID)

	sort.Slice(rightID, func(i, j int) bool {
		return rightID[i] < rightID[j]
	})

	sort.Slice(leftID, func(i, j int) bool {
		return leftID[i] < leftID[j]
	})

	totalDistance := getTotalDistance(&leftID, &rightID)
	fmt.Println("--Part One--")
	fmt.Printf("Total Distance: %d\n", totalDistance)
}

// PART TWO

func parseInput(file *os.File, leftID *[]int, rightIdMap *map[int]int) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		leftNum := strings.Fields(scanner.Text())[0]
		rightNum := strings.Fields(scanner.Text())[1]

		l, err := strconv.Atoi(leftNum)
		if err != nil {
			log.Fatal(err)
		}
		*leftID = append(*leftID, l)

		r, err := strconv.Atoi(rightNum)
		if err != nil {
			log.Fatal(err)
		}

		(*rightIdMap)[r]++

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func getSimilarityScore(leftID *[]int, rightIdMap *map[int]int) int {
	score := 0

	for i := range *leftID {
		if val, exists := (*rightIdMap)[((*leftID)[i])]; exists {
			score += ((*leftID)[i] * val)
		}
	}

	return score
}

func partTwo() {
	var leftID []int
	rightIdMap := make(map[int]int)

	file := readFile("./input.txt")
	defer file.Close()

	parseInput(file, &leftID, &rightIdMap)
	score := getSimilarityScore(&leftID, &rightIdMap)

	fmt.Println("--Part Two--")
	fmt.Println("Similarity Score:", score)

}

func main() {
	partOne()
	partTwo()
}
