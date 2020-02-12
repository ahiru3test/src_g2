package main

import (
    "encoding/csv"
    "fmt"
	"os"
//	"strings"
    "./class"
)

//----------
//Class


//func (me CSVTable) init(sFilePath string) {
/*
	//init
	me.sFilePath = sFilePath

	//open
	file, err := os.Open(me.sFilePath)
    if err != nil {
        panic(err)
    }
    defer file.Close()

	//read
    reader := csv.NewReader(file)
	var line []string
	line, err = reader.Read()

	for i, s := range line {
		fmt.Printf("index: %d, name: %s, hlen:%d\n", i, s, len(me.sHeader))
//		me.sHeader.append(s)
//		append(me.sHeader, s)
//		me.sHeader := &s
//		me.sHeader[i] = &s
	}

//	n := len(line)
*/

/*
	for i:=0; i<len(line); i++ {
		o.mCsv[i] = line[i]
		fmt.Printf("line[%d]:%s\n", i, line[i])
	}
*/

/*
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
*/

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

//}

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




//----------
func main() {
	mCDict := NewCharaDictionary()
	mCDict.mDictionary[0] = NewChara()

	fmt.Printf("fname: %s\n", "aaa")

	//read file
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


//	mCDict[0] = "100"
//	fmt.Printf("fname: %s\n", mCDict[0])

//	m := make(map[string]int)
//	m := map[string]int{}
/*
csvt := NewCSVTable()
	csvt.init("./200130/aaa.csv")
	fmt.Printf("fname: %s\n", csvt.sFilePath)
*/	
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

//---

//CharaDictionary Class
type CharaDictionary struct{
	mDictionary		map[int]*Chara
}
//NewCharaDictionary CharaDictionary constructor
func NewCharaDictionary () *CharaDictionary {
	return new(CharaDictionary)
}

//---

