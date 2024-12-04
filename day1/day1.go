package day1

import (
	"aoc2024/helpers"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
)

func Part1() {

	slice, err := helpers.ReadFile("./day1/input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	left := []int{}
	right := []int{}

	for _, txt := range slice {

		line := strings.Split(txt, "   ")

		leftI, err := strconv.Atoi(line[0])
		if err != nil {
			log.Fatal(err.Error())
		}
		left = append(left, leftI)

		rightI, err := strconv.Atoi(line[1])
		if err != nil {
			log.Fatal(err.Error())
		}
		right = append(right, rightI)
	}

	sort.Ints(left)
	sort.Ints(right)

	sum := 0

	for i, n := range left {
		sum += int(math.Abs(float64(n) - float64(right[i])))
	}

	log.Print(sum)
	return
}

func Part2() {
	slice, err := helpers.ReadFile("./day1/input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	left := []int{}
	right := []int{}

	for _, txt := range slice {

		line := strings.Split(txt, "   ")

		leftI, err := strconv.Atoi(line[0])
		if err != nil {
			log.Fatal(err.Error())
		}
		left = append(left, leftI)

		rightI, err := strconv.Atoi(line[1])
		if err != nil {
			log.Fatal(err.Error())
		}
		right = append(right, rightI)
	}

	sum := 0
	for _, ln := range left {
		count := 0
		for _, rn := range right {
			if rn == ln {
				count++
			}
		}
		sum += (count * ln)
	}

	log.Print(sum)
}
