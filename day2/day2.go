package day2

import (
	"aoc2024/helpers"
	"log"
	"strconv"
	"strings"
)

func Part1() {
	slice, err := helpers.ReadFile("day2/input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	count := 0
	for _, txt := range slice {

		txtSlice := strings.Split(txt, " ")

		nums := []int{}
		for _, ts := range txtSlice {
			num, err := strconv.Atoi(ts)
			if err != nil {
				log.Fatal(err.Error())
			}
			nums = append(nums, num)
		}

		isIncreasing := nums[1] > nums[0]
		isDecreasing := nums[1] < nums[0]

		allGood := true
		for i, num := range nums[:len(nums)-1] {
			diff := nums[i+1] - num
			if diff >= 1 && diff <= 3 && isIncreasing {
				allGood = true
			} else if diff <= -1 && diff >= -3 && isDecreasing {
				allGood = true
			} else {
				allGood = false
				break
			}
		}
		if allGood {
			count++
		}
	}
	log.Print(count)
}

func Part2() {
	slice, err := helpers.ReadFile("./day2/input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	count := 0
	for _, txt := range slice {

		txtSlice := strings.Split(txt, " ")

		nums := []int{}
		for _, ts := range txtSlice {
			num, err := strconv.Atoi(ts)
			if err != nil {
				log.Fatal(err.Error())
			}
			nums = append(nums, num)
		}

		rmIdx := 0
		for rmIdx < len(nums) {

			shortNums := []int{}

			for i, num := range nums {
				if i != rmIdx {
					shortNums = append(shortNums, num)
				}
			}

			isIncreasing := shortNums[1] > shortNums[0]
			isDecreasing := shortNums[1] < shortNums[0]

			allGood := true
			for i, num := range shortNums[:len(shortNums)-1] {
				diff := shortNums[i+1] - num
				if diff >= 1 && diff <= 3 && isIncreasing {
					allGood = true
				} else if diff <= -1 && diff >= -3 && isDecreasing {
					allGood = true
				} else {
					allGood = false
					break
				}
			}
			if allGood {
				count++
				break
			}

			rmIdx++
		}

	}
	log.Print(count)
}
