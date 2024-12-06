package main

import (
	"os"
	"sort"
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

func Distance(left []int, right []int) int {
	sort.Ints(left)
	sort.Ints(right)

	distance := 0
	for index, value := range left {
		res := right[index] - value
		distance += max(res, -res)
	}
	return distance
}
