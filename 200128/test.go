package main

import (
    "encoding/csv"
    "fmt"
	"os"
	"strings"
)

//----------
//Class

//CSVTable Class
type CSVTable struct{
	sFilePath	string
//	eErr		error
	sHeader		[]String
	mCsv		map[int]string
	mCsv2		map[int][string]string
}

//NewCSVTable CSVTable constructor
func NewCSVTable (sFilePath string) *CSVTable {
	//init
	o := new(CSVTable)
	o.sFilePath = sFilePath
//	o.eErr = nil

	//open
	file, err := os.Open(o.sFilePath)
    if err != nil {
        panic(err)
    }
    defer file.Close()

	//read
    reader := csv.NewReader(file)
	var line []string
	var s string
	line, err = reader.Read()

	for _, str
	s = line[0]
		fmt.Printf("[%s]", s)
	if err != nil {
		slice := strings.Split(s, ",")
		for _, str := range slice {
			fmt.Printf("[%s]", str)
		}
	} else {
		//error
	}

	//read firstline
/*
    for {
        line, err = reader.Read()
        if err != nil {
            break
        }
        fmt.Println(line)
    }
*/

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
	fmt.Printf("fname: %s\n", csvt.sFilePath)
	
	//
/*	
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
*/
}
