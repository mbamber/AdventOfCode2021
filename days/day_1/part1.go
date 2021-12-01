package main

import (
	"context"
	"strconv"
	"strings"
)

// Part1 solves the first part of the day's puzzle
func Part1(ctx context.Context, input string) (interface{}, error) {
	depthStrings := strings.Fields(input)
	var depths = make([]int, len(depthStrings))
	for i, d := range depthStrings {
		depth, err := strconv.Atoi(d)
		if err != nil {
			return nil, err
		}
		depths[i] = depth
	}

	count := 0
	for i := 0; i < len(depths)-1; i++ {
		if depths[i] < depths[i+1] {
			count++
		}
	}
	return count, nil
}
