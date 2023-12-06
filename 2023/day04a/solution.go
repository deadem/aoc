package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "strconv"
    "slices"
)

func calcline(str string) int {
    regexstring := `Card\s+(?:\d+):\s*((?:\d+\s*)+)\|\s*((?:\d+\s*)+)`;
    ref, _ := regexp.Compile(regexstring)
    match := ref.FindStringSubmatch(str)

    numbersregex := `\d+`
    refnumber, _ := regexp.Compile(numbersregex)
    left := []int{}
    for _, numbers := range refnumber.FindAllString(match[1], -1) {
        number, _ := strconv.Atoi(string(numbers))
        left = append(left, number)
    }
    right := []int{}
    for _, numbers := range refnumber.FindAllString(match[2], -1) {
        number, _ := strconv.Atoi(string(numbers))
        right = append(right, number)
    }

    count := 0
    for _, number := range left {
        if slices.IndexFunc(right, func(c int) bool { return c == number }) >= 0 {
            if count == 0 {
                count = 1
            } else {
                count = count * 2
            }
        }
    }

    // fmt.Println(left, "---\n---", right, "--", count)

    return count
}

func process(lines []string) int {
    // fmt.Println(lines)

    sum := 0
    for y := 0; y < len(lines); y++ {
        sum = sum + calcline(lines[y])
    }

    return sum
}

func main() {
    file, _ := os.Open("./input.txt")
    reader := bufio.NewReader(file);

    lines := []string{}
    for {
        linebytes, _, err := reader.ReadLine()
        if err != nil {
            break
        }
        line := string(linebytes)
        lines = append(lines, line)
    }
    sum := process(lines)
    fmt.Println(sum)
}
