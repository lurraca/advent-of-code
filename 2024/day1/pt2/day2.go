package main

import (
	"os"
	"strconv"
	"strings"
)

const filename = "2024/day1/input.txt"

func main() {
	input := loadInput(filename)

	left, right := generateListAsSlices(input)

	println(Distance(left, right))
}

func loadInput(filename string) string {
	b, err := os.ReadFile("2024/day1/input.txt")
	if err != nil {
		panic(err)
	}
	input := string(b)
	return input
}

func generateListAsSlices(input string) ([]int, []int) {
	var left []int
	var right []int

	for _, line := range strings.Split(input, "\n") {
		leftString, _ := strconv.Atoi(strings.Split(line, "   ")[0])
		left = append(left, leftString)
		rightString, _ := strconv.Atoi(strings.Split(line, "   ")[1])
		right = append(right, rightString)
	}
	return left, right
}

func count(el int, set []int) int {
	count := 0
	for i, _ := range set {
		if el == set[i] {
			count++
		}
	}

	return count
}

func Distance(left []int, right []int) int {
	distance := 0
	set := make(map[int]int)

	for _, value := range left {
		_, exist := set[value]
		if !exist {
			distance += count(value, right) * value
		}
	}

	return distance
}
