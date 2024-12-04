package day3

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Part1() {
	f, err := os.ReadFile("./day3/input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	content := string(f)

	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	matches := re.FindAllStringSubmatch(content, -1)

	sum := 0
	for _, match := range matches {
		// match[0] is the entire match, match[1] is X, match[2] is Y
		x, err := strconv.Atoi(match[1])
		if err != nil {
			log.Fatal(err.Error())
		}

		y, err := strconv.Atoi(match[2])
		if err != nil {
			log.Fatal(err.Error())
		}

		sum += (x * y)
	}

	log.Print(sum)
}

func Part2() {
	f, err := os.ReadFile("./day3/input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	content := string(f)

	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	matches := re.FindAllStringSubmatch(content, -1)

	sum := 0
	for _, match := range matches {

		startIdx := strings.Index(content, match[0])

		slicedContent := content[0:startIdx]

		doIdx := strings.LastIndex(slicedContent, "do()")

		dontIdx := strings.LastIndex(slicedContent, "don't()")

		if dontIdx > doIdx {
			continue
		}

		x, err := strconv.Atoi(match[1])
		if err != nil {
			log.Fatal(err.Error())
		}

		y, err := strconv.Atoi(match[2])
		if err != nil {
			log.Fatal(err.Error())
		}

		sum += (x * y)

	}

	log.Print(sum)
}
