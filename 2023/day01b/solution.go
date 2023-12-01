package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "strconv"
)

func parse(str string) int {
    i, err := strconv.Atoi(str)
    if err == nil {
        return i
    }
    
    value := map[string]int {
        "one": 1,
        "two": 2,
        "three": 3,
        "four": 4,
        "five": 5,
        "six": 6,
        "seven": 7,
        "eight": 8,
        "nine": 9,
    }

    return value[str]
}

func calcline(str string) int {
    regexstring := `(\d|one|two|three|four|five|six|seven|eight|nine)`;

    ref, _ := regexp.Compile(regexstring)
    first := ref.FindStringSubmatch(str)[1]

    rel, _ := regexp.Compile(`.*` + regexstring)
    last := rel.FindStringSubmatch(str)[1]

    // fmt.Println(str, first, last)
    
    i := parse(first) * 10 + parse(last)
    return i
}

func main() {
    file, _ := os.Open("./input.txt")
    reader := bufio.NewReader(file);

    sum := 0
    for {
        linebytes, _, err := reader.ReadLine()
        if err != nil {
            break
        }
        line := string(linebytes)
        sum += calcline(line)
    }

    fmt.Println(sum)
}
