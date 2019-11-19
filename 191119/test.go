package main

import (
    "encoding/csv"
    "fmt"
    "os"
)

func main() {
    file, err := os.Open("./aaa.csv")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    reader := csv.NewReader(file)
    var line []string

    for {
        line, err = reader.Read()
        if err != nil {
            break
        }
        fmt.Println(line)
    }
}
