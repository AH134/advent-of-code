package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseInput(filename string, reportsArr *[][]int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	row := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		reportRow := strings.Fields(scanner.Text())
		var newRow []int

		// conver numbers into integers
		for i := 0; i < len(reportRow); i++ {
			num, err := strconv.Atoi(reportRow[i])
			if err != nil {
				log.Fatal(err)
			}
			newRow = append(newRow, num)
		}

		(*reportsArr) = append((*reportsArr), newRow)
		row++
	}

}

func absDiff(x, y int) int {
	if x < y {
		return y - x
	}

	return x - y
}

// part one
func getSafeReportsOne(reportsArr *[][]int) int {
	var isIncreasing bool
	totalSafeReports := 0

	for _, row := range *reportsArr {
		isSafe := true
		isIncreasing = row[0] < row[1]
		for j := range len(row) - 1 {
			if row[j] < row[j+1] != isIncreasing {
				isSafe = false
				break
			}

			diff := absDiff(row[j], row[j+1])
			// fmt.Println("difference:", diff)
			if diff < 1 || diff > 3 {
				isSafe = false
				break
			}
		}

		if isSafe {
			totalSafeReports++
		}
	}

	return totalSafeReports

}

func main() {
	// every row is a single report
	// reports contain multiple numbers
	reports := make([][]int, 0)
	parseInput("./input.txt", &reports)

	totalSafeReports := getSafeReportsOne(&reports)
	fmt.Println("Total Safe Reports:", totalSafeReports)
}
