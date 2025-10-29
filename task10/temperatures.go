// Temperature grouping program.
// Groups temperature values by ranges of 10 degrees.
package main

import (
	"fmt"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Celsius float32

// InputTemperatures lets the user input temperature values and validates input data.
func InputTemperatures(ranges map[int][]Celsius, rangesKeys *[]int) {
	for input := ""; ; {
		fmt.Println("Enter the temperature value or type Stop to finish.")

		if _, err := fmt.Scan(&input); err != nil {
			log.Println(err)
			fmt.Println("Something went wrong")
			return
		}

		if strings.ToLower(strings.TrimSpace(input)) == "stop" {
			if len(ranges) == 0 {
				fmt.Println("No values were received.")
			} else {
				fmt.Println("\nHere are the results:")
			}
			return
		}

		tempVal, err := strconv.ParseFloat(input, 32)
		if err != nil {
			fmt.Print("Error: invalid input.\n\n")
			continue
		}

		if tempVal < -273.15 {
			fmt.Print("Error: temperature below absolute zero.\n\n")
			continue
		}

		fmt.Print("Temperature added successfully.\n\n")

		key := int(math.Floor(tempVal / 10))
		if _, ok := ranges[key]; !ok {
			*rangesKeys = append(*rangesKeys, key)
		}

		ranges[key] = append(ranges[key], Celsius(tempVal))
	}
}

// OutputTemperatures prints the grouped temperature values in ascending order.
func OutputTemperatures(ranges map[int][]Celsius, rangesKeys *[]int) {
	for _, key := range *rangesKeys {
		fmt.Println(ranges[key])
	}
}

func main() {
	// Map storing temperature values grouped in increments of 10.
	ranges := make(map[int][]Celsius)

	// Slice storing the integer keys (temperature ranges).
	rangesKeys := make([]int, 0)

	InputTemperatures(ranges, &rangesKeys)

	sort.Ints(rangesKeys)

	OutputTemperatures(ranges, &rangesKeys)
}
