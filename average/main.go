package main

import (
    "fmt"
    "github.com/nhaplycafedang/headfirstgo/datafile"
    "log"
)

func main() {
    numbers, err := datafile.GetFloats("data.txt")
    if err != nill {
        log.Fatal(err)
    }
    var sum float64 = 0
    for _, number := range numbers {
        sum += number
    }
    sampleCount := float64(len(numbers))
    fmt.Printf("Average: %0.2f\n", sum/sampleCount)
}
