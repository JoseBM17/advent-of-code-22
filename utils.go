package advent_of_code_22

import (
	"bufio"
	"os"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
func ReadLines(filename string) []string {
	file, err := os.Open(filename)
	Check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	lines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	Check(scanner.Err())
	return lines
}
