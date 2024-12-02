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

		// Increasing or decreasing should be sampled better now - We'll get the most decreases or increases as general direction
		var de, in int
		for i := 0; i < len(report.values)-1; i++ { // count the -creases
			switch {
			case report.values[i] > report.values[i+1]:
				de++
			case report.values[i] < report.values[i+1]:
				in++
			}
		}

		// get the general direction
		switch {
		case in > de:
			report.crease = "in"
		case de > in:
			report.crease = "de"
		default: // This one has no real direction? All the same values? We'll just pick a direction, will get sorted later.
			report.crease = "in"
		}
		report.unsafe = checkReportvaluesUnsafeWithDampener(report)
		reports = append(reports, report)
	}
	// fmt.Printf("%v\n", reports)
	return reports
}

// checkReportvaluesUnsafeWithDampener generates dampened versions and looks for a safe version among those.
// Kicker is: if we just find one safe solution we can return safe (false).
func checkReportvaluesUnsafeWithDampener(report Report) bool {
	if !checkReportvaluesUnsafe(report) { // The original can also be safe
		return false
	}
	for i := 0; i < len(report.values); i++ {
		dampenedReport := Report{
			crease: report.crease,
		}
		for n, val := range report.values {
			if i != n {
				dampenedReport.values = append(dampenedReport.values, val)
			}
		}
		if !checkReportvaluesUnsafe(dampenedReport) {
			return false
		}
	}

	return true // We didn't run into any safe report
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
		if !report.unsafe {
			num++
		}
	}
	return num
}
