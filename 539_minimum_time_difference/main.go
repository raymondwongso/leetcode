package main

import (
	"fmt"
	"slices"
	"time"
)

func main() {
	timePoints := []string{"01:01", "02:01", "03:00"}
	fmt.Println(findMinDifference(timePoints))
}

// time complexity: O(n) where n is len(timePoints)
// space complexity: O(1)
func findMinDifference(timePoints []string) int {
	slices.Sort(timePoints)

	layout := "15:04"
	min := 1440
	x := 0

	for i := 1; i < len(timePoints); i++ {
		t1, _ := time.Parse(layout, timePoints[i-1])
		t2, _ := time.Parse(layout, timePoints[i])
		x = int(t2.Sub(t1).Minutes())

		if x < min {
			min = x
		}
	}

	t0, _ := time.Parse(layout, timePoints[0])
	tlast, _ := time.Parse(layout, timePoints[len(timePoints)-1])
	t0 = t0.Add(24 * time.Hour)
	x = int(t0.Sub(tlast).Minutes())
	if x < min {
		min = x
	}

	return min
}
