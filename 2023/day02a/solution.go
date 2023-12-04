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
        "red": 12,
        "green": 13,
        "blue": 14,
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

            if limits[color] < count {
                return 0
            }

            fmt.Println(count, color)
        }
    }

    return gameId
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
