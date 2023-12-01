package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "strconv"
)

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
        fmt.Println("Read line:", line)

        re, _ := regexp.Compile(`\d`)
        matched := re.FindAllString(line, -1)
        
        i, _ := strconv.Atoi(matched[0] + matched[len(matched) - 1])
        fmt.Println(i)
        sum += i

        fmt.Println(matched[0], matched[len(matched) - 1])
    }

    fmt.Println(sum)
}
