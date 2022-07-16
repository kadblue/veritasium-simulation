package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	success := 0
	times := 100000
	size := 100

	for i := 0; i < times; i++ {
		if Simulate(size) {
			success++
		}
	}
	fmt.Println(success)
}

func CreateBoxes(size int) []int {
	rand.Seed(time.Now().UnixNano())

	numbers := make([]int, size)

	for i := 0; i < size; i++ {
		numbers[i] = i
	}

	rand.Shuffle(len(numbers), func(i, j int) { numbers[i], numbers[j] = numbers[j], numbers[i] })

	return numbers
}

func PrisonerAttempt(numbers []int, prisonerNumber int, attempts int) bool {
	currentNumber := numbers[prisonerNumber]
	for i := 0; i < attempts-1; i++ {
		if currentNumber == prisonerNumber {
			return true
		}
		currentNumber = numbers[currentNumber]
	}
	return false
}

func Simulate(size int) bool {
	numbers := CreateBoxes(size)
	for i := 0; i < size; i++ {
		if !PrisonerAttempt(numbers, i, size/2) {
			return false
		}
	}
	return true
}
