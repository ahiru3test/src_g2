package main

import (
	"fmt"
)

//GMain is struct
type GMain struct {
	i int
	n int
	s string
}
//NewGMain is GMain's constructor
func NewGMain (i int, n int) *GMain {
    p := new(GMain)
    p.i = i
    p.n = n
    return p
}

func main() {
	gm := NewGMain(111, 222)
    fmt.Printf("i: %d n: %d \n", gm.i, gm.n)
}
