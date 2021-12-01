package main_test

import (
	"context"
	"fmt"
	"log"
	"testing"

	main "github.com/mbamber/aoc/days/day_1"
	"github.com/mbamber/aoc/inputs"
	"github.com/stretchr/testify/assert"
)

func TestPart2Answer(t *testing.T) {
	expected := 1781

	input := inputs.Load(1)
	out, err := main.Part2(context.Background(), input)
	assert.NoError(t, err)
	assert.Equal(t, expected, out)
}

func ExamplePart2() {
	input := inputs.LoadExample(1, 1)
	out, err := main.Part2(context.Background(), input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out)
	// Output:
	// 5
}
