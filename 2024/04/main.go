package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Puzzle struct {
	width  int
	height int
	matrix [][]string
}

func getFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Error opening", filename)
	}

	return file
}

func filterFile(file *os.File) Puzzle {
	var puzzle Puzzle
	arr := make([][]string, 0)

	row := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		newRow := make([]string, 0)
		newRow = append(newRow, strings.Split(scanner.Text(), "")...)
		arr = append(arr, newRow)
		row++
	}

	puzzle.matrix = arr
	puzzle.height = len(arr)
	puzzle.width = len(arr[0])

	return puzzle
}

func getTotalXmas(puzzle *Puzzle) int {
	count := 0
	for row := 0; row < (*puzzle).height; row++ {
		for col := 0; col < (*puzzle).width; col++ {
			checkXmas(puzzle, row, col, &count)
		}
	}

	return count
}

func checkXmas(puzzle *Puzzle, row, col int, count *int) {
	if (*puzzle).matrix[row][col] != "X" {
		return
	}

	checkXmasToEast(puzzle, row, col, count)
	checkXmasToNorth(puzzle, row, col, count)
	checkXmasToWest(puzzle, row, col, count)
	checkXmasToSouth(puzzle, row, col, count)
	checkXmasToNorthEast(puzzle, row, col, count)
	checkXmasToNorthWest(puzzle, row, col, count)
	checkXmasToSouthEast(puzzle, row, col, count)
	checkXmasToSouthWest(puzzle, row, col, count)

}

func gotMatch(puzzle *Puzzle, row, col int, symbol string) bool {

	if col < 0 {
		return false
	}
	if row < 0 {
		return false
	}

	if row >= (*puzzle).height {
		return false
	}
	if col >= (*puzzle).width {
		return false
	}

	return (*puzzle).matrix[row][col] == symbol
}
func checkXmasToEast(puzzle *Puzzle, row, col int, counter *int) {
	if !gotMatch(puzzle, row, col+1, "M") {
		return
	}
	if !gotMatch(puzzle, row, col+2, "A") {
		return
	}
	if !gotMatch(puzzle, row, col+3, "S") {
		return
	}
	*counter += 1
}

func checkXmasToWest(puzzle *Puzzle, row, col int, counter *int) {
	if !gotMatch(puzzle, row, col-1, "M") {
		return
	}
	if !gotMatch(puzzle, row, col-2, "A") {
		return
	}
	if !gotMatch(puzzle, row, col-3, "S") {
		return
	}
	*counter += 1
}

func checkXmasToNorth(puzzle *Puzzle, row, col int, counter *int) {
	if !gotMatch(puzzle, row-1, col, "M") {
		return
	}
	if !gotMatch(puzzle, row-2, col, "A") {
		return
	}
	if !gotMatch(puzzle, row-3, col, "S") {
		return
	}
	*counter += 1
}

func checkXmasToSouth(puzzle *Puzzle, row, col int, counter *int) {
	if !gotMatch(puzzle, row+1, col, "M") {
		return
	}
	if !gotMatch(puzzle, row+2, col, "A") {
		return
	}
	if !gotMatch(puzzle, row+3, col, "S") {
		return
	}
	*counter += 1
}

func checkXmasToNorthEast(puzzle *Puzzle, row, col int, counter *int) {
	if !gotMatch(puzzle, row-1, col+1, "M") {
		return
	}
	if !gotMatch(puzzle, row-2, col+2, "A") {
		return
	}
	if !gotMatch(puzzle, row-3, col+3, "S") {
		return
	}
	*counter += 1
}

func checkXmasToNorthWest(puzzle *Puzzle, row, col int, counter *int) {
	if !gotMatch(puzzle, row-1, col-1, "M") {
		return
	}
	if !gotMatch(puzzle, row-2, col-2, "A") {
		return
	}
	if !gotMatch(puzzle, row-3, col-3, "S") {
		return
	}
	*counter += 1
}

func checkXmasToSouthEast(puzzle *Puzzle, row, col int, counter *int) {
	if !gotMatch(puzzle, row+1, col+1, "M") {
		return
	}
	if !gotMatch(puzzle, row+2, col+2, "A") {
		return
	}
	if !gotMatch(puzzle, row+3, col+3, "S") {
		return
	}
	*counter += 1
}

func checkXmasToSouthWest(puzzle *Puzzle, row, col int, counter *int) {
	if !gotMatch(puzzle, row+1, col-1, "M") {
		return
	}
	if !gotMatch(puzzle, row+2, col-2, "A") {
		return
	}
	if !gotMatch(puzzle, row+3, col-3, "S") {
		return
	}
	*counter += 1
}

func main() {
	file := getFile("input.txt")
	filteredList := filterFile(file)
	// fmt.Println(filteredList.matrix)
	total := getTotalXmas(&filteredList)
	fmt.Println("total:", total)
}
