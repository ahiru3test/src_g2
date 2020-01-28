package main

import (
	"fmt"
)

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
