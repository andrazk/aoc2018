package helpers

import (
	"bufio"
	"os"
)

// ReadLines returns each line in file as string
func ReadLines(path string) (out []string, err error) {
	file, err := os.Open(path)
	if err != nil {
		return out, err
	}
	defer file.Close()

	out = make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		out = append(out, scanner.Text())
	}

	return out, scanner.Err()
}
