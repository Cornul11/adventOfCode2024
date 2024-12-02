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

	var firstNumbers []int
	var secondNumbers []int
	for scanner.Scan() {
		line := scanner.Text()
		// split the line by space
		numbers := strings.Split(line, "   ")
		if len(numbers) != 2 {
			log.Fatal("Invalid input format")
		}
		firstNumberInt, err := strconv.Atoi(strings.TrimSpace(numbers[0]))
		if err != nil {
			log.Fatal(err)
		}
		secondNumberInt, err := strconv.Atoi(strings.TrimSpace(numbers[1]))
		if err != nil {
			log.Fatal(err)
		}
		firstNumbers = append(firstNumbers, firstNumberInt)
		secondNumbers = append(secondNumbers, secondNumberInt)
	}
	if len(firstNumbers) != len(secondNumbers) {
		log.Fatal("Invalid input format")
	}

	// part 1
	sum := similarityDifferenceSum(firstNumbers, secondNumbers)

	// part 2
	sum = similarityScoreSum(firstNumbers, secondNumbers)
	log.Println(sum)
}

func similarityScoreSum(firstNumbers []int, secondNumbers []int) int {
	score := 0

	for _, num := range firstNumbers {
		numOccurrences := countOccurrences(secondNumbers, num)
		score += num * numOccurrences
	}
	return score
}

func similarityDifferenceSum(firstNumbers []int, secondNumbers []int) int {
	var differences []int

	for len(firstNumbers) > 0 {
		firstNumber := min(firstNumbers)
		secondNumber := min(secondNumbers)
		difference := firstNumber - secondNumber
		if difference < 0 {
			difference = difference * -1
		}
		differences = append(differences, difference)
		firstNumbers = filterSlice(firstNumbers, firstNumber)
		secondNumbers = filterSlice(secondNumbers, secondNumber)
	}

	sum := 0
	for _, diff := range differences {
		sum += diff
	}
	return sum
}

// min returns the minimum value in a slice of integers
func min(numbers []int) int {
	if len(numbers) == 0 {
		log.Fatal("Cannot find minimum of an empty slice")
	}
	minValue := numbers[0]
	for _, num := range numbers {
		if num < minValue {
			minValue = num
		}
	}
	return minValue
}

// filterSlice removes the first occurrence of a value from a slice of integers
func filterSlice(slice []int, valueToRemove int) []int {
	filtered := []int{}
	removed := false
	for _, num := range slice {
		if num == valueToRemove && !removed {
			removed = true
			continue
		}
		filtered = append(filtered, num)
	}
	return filtered
}

func countOccurrences(slice []int, value int) int {
	count := 0
	for _, num := range slice {
		if num == value {
			count++
		}
	}
	return count
}
