package main

import (
    "fmt"
    "./data"
)

func main() {
    var x data.Mydata
    x.Num = 10
    x.Str = "something"
    fmt.Printf("x.Num=%d, x.Str=%s\n", x.Num, x.Str)
}
