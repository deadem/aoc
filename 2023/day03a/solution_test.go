package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestCalcline(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(0, process([]string{ "467..114.." }))
    assert.Equal(4361, process([]string{
        "467..114..",
        "...*......",
        "..35..633.",
        "......#...",
        "617*......",
        ".....+.58.",
        "..592.....",
        "......755.",
        "...$.*....",
        ".664.598..",
    }))
}
