package helpers

import (
	"bufio"
	"os"
)

func ReadFile(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)

	lines := []string{}

	for scanner.Scan() {
		txt := scanner.Text()

		lines = append(lines, txt)
	}

	return lines, nil

}
