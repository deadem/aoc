package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "strconv"
    "strings"
)

func calcline(str string) int {
    regexstring := `Game (\d+): (.+)`;

    limits := map[string]int {
        "red": 0,
        "green": 0,
        "blue": 0,
    }

    ref, _ := regexp.Compile(regexstring)
    match := ref.FindStringSubmatch(str)
    gameId, _ := strconv.Atoi(match[1])
    takes := match[2]

    fmt.Println(gameId, takes)

    for _, take := range strings.Split(takes, ";") {
        for _, colors := range strings.Split(take, ",") {
            refColor, _ := regexp.Compile(`(\d+)\s+(\S+)`)
            match := refColor.FindStringSubmatch(colors)

            count, _ := strconv.Atoi(match[1])
            color := match[2]

            limits[color] = max(limits[color], count)
            fmt.Println(count, color)
        }
    }

    return limits["red"] * limits["green"] * limits["blue"]
}

func isSymbol(char byte) bool {
    return char != '.' && (char < '0' || char > '9')
}

func isNumber(char byte) bool {
    return char >= '0' && char <= '9';
}

func isPartnumber(lines []string, x int, y int) bool {
    char := lines[y][x]
    if isNumber(char) {
        havesymbol := (y > 0 && (isSymbol(lines[y - 1][x]) || (x > 0 && isSymbol(lines[y - 1][x - 1])) || isSymbol(lines[y - 1][x + 1]))) ||
        (x > 0 && isSymbol(lines[y][x - 1])) || isSymbol(lines[y][x + 1]) ||
        isSymbol(lines[y + 1][x]) || (x > 0 && isSymbol(lines[y + 1][x - 1])) || isSymbol(lines[y + 1][x + 1])

        if havesymbol {
            return true
        }

        return isPartnumber(lines, x + 1, y)
    }

    return false
}

func process(lines []string) int {
    fmt.Println(lines)

    for y := 0; y < len(lines); y++ {
        for x := 0; x < len(lines[y]); x++ {
            char := lines[y][x]

            if isNumber(char) {
                partnumber := isPartnumber(lines, x, y)
                for x = x + 1; x < len(lines[y]) && isNumber(lines[y][x]); x++ {
                    if !partnumber {
                        line := lines[y]
                        lines[y] = line[0 : x - 1] + "." + line[x + 1 : len(line)]
                    }
                }
            }
        }
    }

    fmt.Println(lines)
    return 0
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
