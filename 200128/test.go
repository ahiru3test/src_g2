package main

import (
    "encoding/csv"
    "fmt"
	"os"
)

//----------
//Class

//CSVTable Class
type CSVTable struct{
	sFile	string
	eErr	error
	mCsv	map[string]int
}

//NewCSVTable CSVTable constructor
func NewCSVTable (file string) *CSVTable {
	o := new(CSVTable)
	o.sFile = file
	o.eErr = nil
	//

	return o
}

/*
func (t CSVTable) init(file string) {
	sFile := file
}
*/

/*
func NewPerson (name string, age int) *Person {
    if age < 0 {
        return nil
    }
    p := new(Person)
    p.Name = name
    p.Age = age
    return p
}
*/

//Chara Class
type Chara struct{
}
func (t Chara) aaa() {
}

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


//----------
func main() {
//	m := make(map[string]int)
//	m := map[string]int{}
	csvt := NewCSVTable("./aaa.csv")
	fmt.Printf("fname: %s, error:%s\n", csvt.sFile, csvt.eErr)
	
	//
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
