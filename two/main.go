package two

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func One(inputPath string) int {
	file, err := os.Open(inputPath)
	if err != nil {
		fmt.Printf("Error happened opening file: %s\n", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	horizontalPosition, depth := 0, 0

	for scanner.Scan() {
		directive := strings.Split(scanner.Text(), " ")
		command := directive[0]
		units, _ := strconv.Atoi(directive[1])
		if command == "forward" {
			horizontalPosition += units
		}
		if command == "down" {
			depth += units
		}
		if command == "up" {
			depth -= units
		}
	}
	return horizontalPosition * depth
}

func Two(inputPath string) int {
	return 0
}
