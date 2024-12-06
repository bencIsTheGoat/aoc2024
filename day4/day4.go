package day4

import (
	"aoc2024/helpers"
)

// parse to 2d slices
// search diagonal
func Part1() (int, error) {
	txt, err := helpers.ReadFile("day4/input.txt")
	if err != nil {
		return 0, err
	}

	// Have 2d array
	matrix := [][]string{}
	for _, line := range txt {
		row := []string{}
		for _, char := range line {
			row = append(row, string(char))
		}
		matrix = append(matrix, row)
	}

	// log.Println(WordSearch(txt, "XMAS"))

	count := 0

	// Now do searching
	for y, row := range matrix {
		for x := range row {

			// Up -_
			word := ""
			dist := 0
			for dist < 4 && y-dist >= 0 {
				word += matrix[y-dist][x]
				dist++
			}

			if word == "XMAS" {
				count++
			}

			// Right Up -+
			word = ""
			dist = 0
			for dist < 4 && y-dist >= 0 && x+dist < len(matrix[0]) {
				word += matrix[y-dist][x+dist]
				dist++
			}

			if word == "XMAS" {
				count++
			}

			// Right _+
			word = ""
			dist = 0
			for dist < 4 && x+dist < len(matrix[0]) {
				word += matrix[y][x+dist]
				dist++
			}

			if word == "XMAS" {
				count++
			}

			// Right Down ++
			word = ""
			dist = 0
			for dist < 4 && x+dist < len(matrix[0]) && y+dist < len(matrix) {
				word += matrix[y+dist][x+dist]
				dist++
			}

			if word == "XMAS" {
				count++
			}

			// Down +_
			word = ""
			dist = 0
			for dist < 4 && y+dist < len(matrix) {
				word += matrix[y+dist][x]
				dist++
			}

			if word == "XMAS" {
				count++
			}

			// Down Left
			word = ""
			dist = 0
			for dist < 4 && y+dist < len(matrix) && x-dist >= 0 {
				word += matrix[y+dist][x-dist]
				dist++
			}

			if word == "XMAS" {
				count++
			}

			// Left
			word = ""
			dist = 0
			for dist < 4 && x-dist >= 0 {
				word += matrix[y][x-dist]
				dist++
			}

			if word == "XMAS" {
				count++
			}

			// Left Up
			word = ""
			dist = 0
			for dist < 4 && x-dist >= 0 && y-dist >= 0 {
				word += matrix[y-dist][x-dist]
				dist++
			}

			if word == "XMAS" {
				count++
			}
		}
	}
	return count, nil
}
