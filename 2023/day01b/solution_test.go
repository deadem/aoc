package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestCalcline(t *testing.T) {
    assert := assert.New(t)

    assert.Equal(18, calcline("vjhkjhjkuiyio1osidjlkfhsfhskfhskf7zasjdkh8asdasdabkmiu"))
    assert.Equal(79, calcline("789"))
    assert.Equal(79, calcline("seven111nine"))
    assert.Equal(79, calcline("sevenine"))
    assert.Equal(83, calcline("eighthree"))
    assert.Equal(18, calcline("oneight"))
    assert.Equal(82, calcline("eightwo"))
}
