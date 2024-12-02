// https://adventofcode.com/2024/day/2
// Day 2: Red-Nosed Reports
package main

import (
	"aoc/libaoc"
	"fmt"
	"strings"
)

type Report struct {
	crease string // In- or De- ?
	values []int  // Report data
	unsafe bool   // Is it gonna blow?
}

func main() {
	// input, err := libaoc.ReadLines("example.txt")
	input, err := libaoc.ReadLines("input.txt")
	if err != nil {
		panic("No input!")
	}
	reports := getReports(input)
	fmt.Printf("%d safe reports\n", countSafeReports(reports))
}

// getReports reads the input and makes Reports out of them.
func getReports(input []string) (reports []Report) {
	for _, line := range input {
		exploded := strings.Fields(line)
		var report Report
		for _, val := range exploded {
			report.values = append(report.values, libaoc.SilentAtoi(val))
		}
		switch { // Get the in- or decrease
		case report.values[0] > report.values[1]:
			report.crease = "de"
		case report.values[0] < report.values[1]:
			report.crease = "in"
		default: // Only reached if the same numbers
			report.unsafe = true
			reports = append(reports, report)
			continue
		}
		report.unsafe = checkReportvaluesUnsafe(report)
		reports = append(reports, report)
	}
	return
}

// checkReportvaluesUnsafe checks if a report is unsafe
func checkReportvaluesUnsafe(report Report) bool {
	for i := 0; i < len(report.values)-1; i++ {
		var high, low int

		switch report.crease { // Which is high, which is low?
		case "de":
			high = report.values[i]
			low = report.values[i+1]
		case "in":
			low = report.values[i]
			high = report.values[i+1]
		}
		switch { // Let's do the safety-maths!
		case high-low > 3: // To big a -crease
			return true
		case high-low <= 0: // Wrong way!
			return true
		default: // Nothing wrong here
		}
	}
	return false
}

// countSafeReports returns the number of safe reports
func countSafeReports(reports []Report) (num int) {
	for _, report := range reports {
		if report.unsafe == false {
			num++
		}
	}
	return num
}
