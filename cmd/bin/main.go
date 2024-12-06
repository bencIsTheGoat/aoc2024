package main

import (
	"aoc2024/day4"
	"log"
)

func main() {
	output, err := day4.Part1()
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Print(output)
}
