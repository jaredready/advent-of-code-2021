package one

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func One(inputPath string) int {
	file, err := os.Open(inputPath)
	if err != nil {
		fmt.Printf("Error happened opening file: %s\n", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	previousMeasurement, currentMeasurement := -1, -1
	depthIncrease := 0
	for scanner.Scan() {
		if previousMeasurement == -1 {
			previousMeasurement, _ = strconv.Atoi(scanner.Text())
			currentMeasurement = previousMeasurement
		} else {
			currentMeasurement, _ = strconv.Atoi(scanner.Text())
			if currentMeasurement > previousMeasurement {
				depthIncrease += 1
			}
			previousMeasurement = currentMeasurement
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return depthIncrease
}

func Two(inputPath string) string {
	return inputPath
}
