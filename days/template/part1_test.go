package main_test

import (
	"context"
	"fmt"
	"log"
	"testing"

	main "github.com/mbamber/aoc/days/template"
	"github.com/mbamber/aoc/inputs"
	"github.com/stretchr/testify/assert"
)

func TestPart1Answer(t *testing.T) {
	t.Skip("test template, skipping...")
	expected := 0

	input := inputs.Load(-1)
	out, err := main.Part1(context.Background(), input)
	assert.NoError(t, err)
	assert.Equal(t, expected, out)
}

func ExamplePart1() {
	input := inputs.LoadExample(1, 1)
	out, err := main.Part1(context.Background(), input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out)
	// Output:
	//
}
