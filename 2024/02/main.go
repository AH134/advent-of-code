package main

import (
	"bufio"
	"flag"
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

func safe(reportsList *[][]int) int {
	count := 0
	for _, report := range *reportsList {
		diffs := make([]int, 0)
		for i := 0; i < len(report)-1; i++ {
			diffs = append(diffs, report[i]-report[i+1])
		}

		if allIncreasing(&diffs) || allDecreasing(&diffs) {
			count++
		}
	}

	return count
}

func safePartTwo(reportsList *[][]int) int {
	count := 0
	for _, report := range *reportsList {
		// loop through length of list with one element removed
		// removed and see if it is safe
		for i := 0; i < len(report); i++ {
			reportCopy := make([]int, len(report))
			copy(reportCopy, report)
			reportCopy = append(reportCopy[:i], reportCopy[i+1:]...)

			diffs := make([]int, 0)
			for j := 0; j < len(reportCopy)-1; j++ {
				diffs = append(diffs, reportCopy[j]-reportCopy[j+1])
			}
			if allIncreasing(&diffs) || allDecreasing(&diffs) {
				count++
				break
			}
		}

	}

	return count
}

func allIncreasing(diffList *[]int) bool {
	for _, el := range *diffList {
		if el < 1 || el > 3 {
			return false
		}
	}

	return true
}

func allDecreasing(diffList *[]int) bool {
	for _, el := range *diffList {
		if el > -1 || el < -3 {
			return false
		}
	}

	return true
}

func main() {
	filePtr := flag.String("file", "foo", "a string")
	partPtr := flag.Int("part", 0, "a string")
	flag.Parse()

	if *filePtr == "foo" {
		log.Fatal("Must input a file")
	}

	reports := make([][]int, 0)
	parseInput(*filePtr, &reports)

	switch *partPtr {
	case 1:
		fmt.Println("Safe:", safe(&reports))
	case 2:
		fmt.Println("Safe:", safePartTwo(&reports))
	case 0:
		log.Fatal("Input part")
	}
}
