package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strings"
	"time"
)

const LAYOUT = "15:04"

func parseTimeAndDiff(start_time, end_time string) float64 {
	t1, _ := time.Parse(LAYOUT, end_time)
	t2, _ := time.Parse(LAYOUT, start_time)

	return t1.Sub(t2).Minutes()
}

func sortMapKeysAlphabetically(taskNames map[string]int) []string {
	keys := make([]string, 0, len(taskNames))
	for k := range taskNames {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func main() {
	// First check if filename was given as a command line argument. If not, exit the program.
	if len(os.Args) < 2 {
		log.Fatalf("You need to provide a file to be parsed.")
		return
	}

	// Take log file to parse as a command line argument.
	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("Cannot read file:", err)
		return
	}

	scanner := bufio.NewScanner(bytes.NewReader(file))

	taskNames := make(map[string]int)
	totalDuration := 0

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) >= 1 {
			fields := strings.Split(line, " ")
			taskDurationFields := strings.Split(fields[0], "-")
			taskName := strings.Join(fields[1:], " ")
			start_time := taskDurationFields[0]
			end_time := taskDurationFields[1]

			durationInMinutes := int(parseTimeAndDiff(start_time, end_time))

			totalDuration += durationInMinutes

			if val, ok := taskNames[taskName]; ok {
				// Add value to existing taskName
				taskNames[taskName] = val + durationInMinutes
			} else {
				// Add new taskName
				taskNames[taskName] = durationInMinutes
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return
	}

	for _, key := range sortMapKeysAlphabetically(taskNames) {
		val := taskNames[key]
		totalDurationInPercentages := math.Floor(float64(val) / float64(totalDuration) * 100.0)
		fmt.Printf("%s %d minutes %2.f%%\n", key, val, totalDurationInPercentages)

	}

}
