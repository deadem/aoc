package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestCalcline(t *testing.T) {
    assert := assert.New(t)

    assert.Equal(1, calcline("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"))
    assert.Equal(2, calcline("Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue"))
    assert.Equal(0, calcline("Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red"))
}
