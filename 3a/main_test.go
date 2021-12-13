package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestString2Int1(t *testing.T) {
	input := "1"
	expected := 1

	assert.Equal(t, expected, bin2int(input))
}

func TestString2Int2(t *testing.T) {
	input := "110"
	expected := 6

	assert.Equal(t, expected, bin2int(input))
}

func TestString2Int3(t *testing.T) {
	input := "01001"
	expected := 9

	assert.Equal(t, expected, bin2int(input))
}

func TestString2Int4(t *testing.T) {
	input := "00000"
	expected := 0

	assert.Equal(t, expected, bin2int(input))
}
