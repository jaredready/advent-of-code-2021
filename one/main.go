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

func Two(inputPath string) int {
	file, err := os.Open(inputPath)
	if err != nil {
		fmt.Printf("Error happened opening file: %s\n", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	previousMeasurement, currentMeasurement := -1, -1
	depthIncrease := 0
	window := [3]int{-1, -1, -1}

	for scanner.Scan() {
		window[2] = window[1]
		window[1] = window[0]
		window[0], _ = strconv.Atoi(scanner.Text())

		if window[0] == -1 || window[1] == -1 || window[2] == -1 {
			continue
		}

		if previousMeasurement == -1 {
			previousMeasurement = window[0] + window[1] + window[2]
			currentMeasurement = previousMeasurement
		} else {
			currentMeasurement = window[0] + window[1] + window[2]
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
