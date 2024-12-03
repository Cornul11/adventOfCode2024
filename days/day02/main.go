package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)

	var reportData [][]int
	for scanner.Scan() {
		line := scanner.Text()
		// split the line by space
		numbers := strings.Split(line, " ")
		if len(numbers) <= 0 {
			log.Fatal("Invalid input format")
		}
		var numberList []int
		for _, num := range numbers {
			number, err := strconv.Atoi(strings.TrimSpace(num))
			if err != nil {
				log.Fatal(err)
			}
			numberList = append(numberList, number)
		}
		reportData = append(reportData, numberList)
	}

	// part 1
	numSafeReports := checkSafetyReport(reportData, checkRowSafe)

	// part 2
	numSafeReports = checkSafetyReport(reportData, checkRowSafeWithDampener)
	log.Println(numSafeReports)
}

func checkRowSafeWithDampener(row []int) bool {
	if isRowSafe(row) {
		return true
	}

	// try to remove one number from the row and check if it is safe
	for i := 0; i < len(row); i++ {
		modifiedRow := make([]int, 0, len(row)-1)
		for j := 0; j < len(row); j++ {
			if i != j {
				modifiedRow = append(modifiedRow, row[j])
			}
		}
		if isRowSafe(modifiedRow) {
			return true
		}
	}

	// if no removal makes the row safe, it is unsafe
	return false
}

func checkRowSafe(row []int) bool {
	return isRowSafe(row)
}

func isRowSafe(row []int) bool {
	if len(row) < 2 {
		return true // a single number is always safe
	}
	prev := row[0]
	increasing := true
	decreasing := true
	for _, num := range row[1:] {
		difference := num - prev
		if difference < 0 {
			increasing = false
		} else if difference > 0 {
			decreasing = false
		}
		if abs(difference) < 1 || abs(difference) > 3 {
			return false
		}
		prev = num
	}
	// row is safe if it is strictly increasing or strictly decreasing
	return increasing || decreasing
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func checkSafetyReport(reportData [][]int, rowChecker func([]int) bool) int {
	numSafeReports := 0

	for _, row := range reportData {
		if rowChecker(row) {
			numSafeReports++
		}
	}
	return numSafeReports
}
