package main

import (
	"context"
	"strconv"
	"strings"
)

// Part2 solves the second part of the day's puzzle
func Part2(ctx context.Context, input string) (interface{}, error) {
	depthStrings := strings.Fields(input)
	var depths = make([]int, len(depthStrings))
	for i, d := range depthStrings {
		depth, err := strconv.Atoi(d)
		if err != nil {
			return nil, err
		}
		depths[i] = depth
	}

	var windows = make([]int, len(depths)-2)
	for i := 0; i < len(windows); i++ {
		windows[i] = depths[i] + depths[i+1] + depths[i+2]
	}

	count := 0
	for i := 0; i < len(windows)-1; i++ {
		if windows[i] < windows[i+1] {
			count++
		}
	}
	return count, nil
}
