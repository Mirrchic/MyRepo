package main

import (
	"calk/calculator"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//This file it is just showing the way
//this programs works, without connectoin to Telegramm api
//Types to handle json strings
type Numbers struct {
	Num1           string `json:"Num1"`
	Base1          int    `json:"Base1"`
	Num2           string `json:"Num2"`
	Base2          int    `json:"Base2"`
	BaseRet        int    `json:"BaseRet"`
	Requestcreated string `json:"Requestcreated"`
}

type User struct {
	UserID      int       `json:"UserID"`
	Char        string    `json:"Char"`
	NumbersCalk []Numbers `json:"NumbersCalk"`
}

func main() {
	a := User{}
	readJson, _ := ioutil.ReadFile("numbers.json")
	_ = json.Unmarshal(readJson, &a)
	result := calculator.GetResult(a.NumbersCalk[0].Num1, a.NumbersCalk[0].Num2, a.NumbersCalk[0].Base1, a.NumbersCalk[0].Base2, a.NumbersCalk[0].BaseRet,
		a.Char)
	fmt.Println("Operation were at:", a.NumbersCalk[0].Requestcreated, "\n", "sum:", result, "\n", "Base:", a.NumbersCalk[0].BaseRet)
}
