package main

import (
    "fmt"
    "github.com/Zyian/aocget"
    "os"
    "sort"
    "strconv"
    "strings"
)

func main() {
    aoc := aocget.NewClient(os.Getenv("SESSION_TOKEN"))
    i, err := aoc.DownloadInputString(2022, 1)
    if err != nil {
        fmt.Printf("unable to download input: %v\n", err)
        os.Exit(1)
    }
    elvesCalories := []int{}
    for _, e := range strings.Split(i, "\n\n") {
        calorieTotal := 0
        for _, f := range strings.Split(e, "\n") {
            c, _ := strconv.ParseInt(f, 10, 64)
            calorieTotal += int(c)
        }
        elvesCalories = append(elvesCalories, calorieTotal)
    }

    sort.Ints(elvesCalories)
    fmt.Printf("The top calories held are: %d\n", elvesCalories[len(elvesCalories)-1])

    topThree := 0
    for _, c := range elvesCalories[len(elvesCalories)-3:] {
        topThree += c
    }
    fmt.Printf("The sum of the top 3 is: %d\n", topThree)
}