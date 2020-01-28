package main

//Import
import (
    "encoding/csv"
    "fmt"
	"os"
)

//Class

//Taiyaki Class
type Taiyaki struct{}

//Atama is Taiyaki's Atama
func (t Taiyaki) Atama() {
    fmt.Println("たい焼きの頭の方にはあんこがいっぱい入っている")
}

//Shippo is Taiyaki's Shippo
func (t Taiyaki) Shippo() {
    fmt.Println("たい焼きの尻尾にはあんこがほとんど入っていない")
    fmt.Println("しかしカリカリしていて美味しい")
}


//Main
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
