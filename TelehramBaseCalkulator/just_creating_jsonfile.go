package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

//structur getting numbers,  bases, algebratic char,
//and base that result should be converted to

type Numbers struct {
	Base1, Base2, BaseRet int
	Num1                  string
	Num2                  string
	Requestcreated        string
}

type User struct {
	UserID      int
	Char        string
	NumbersCalk []Numbers
}

//simple Json write method
func main() {
	data := User{
		UserID: 123,
		Char:   "*",
		NumbersCalk: []Numbers{
			Numbers{Num1: "100",
				Base1:          2,
				Num2:           "FF",
				Base2:          16,
				BaseRet:        16,
				Requestcreated: time.Now().String(),
			},
		},
	}
	file, _ := json.MarshalIndent(data, "", " ")
	_ = ioutil.WriteFile("numbers.json", file, 0644)
	fmt.Println(string(file))
}
