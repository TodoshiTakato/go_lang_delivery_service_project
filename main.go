package main

import (
	f "fmt"
)

func main() {
	//f.Println("Hello, Go!")
	//
	//s := "I like Go!"
	//f.Println(s + "\n" + s + "\n" + s)
	//
	//f.Println(strings.Repeat(s+"\n", 3) + s)
	//
	//f.Println("Hello Go"[0])         // вывод: 72
	//f.Println(string("Hello Go"[0])) // вывод: H
	//
	//var hello string
	//
	//hello = "Hello Go!"
	//
	//var a int = 2019
	//
	//f.Println(hello)
	//f.Println(a)

	//var symbol0 uint8 = 'c'
	//var symbol1 uint16 = 'c'
	//var symbol2 uint32 = 'c'
	//var symbol3 uint64 = 'c'
	//var symbol4 int8 = 'c'
	//var symbol5 int16 = 'c'
	//var symbol6 int32 = 'c'
	//var symbol7 int64 = 'c'
	//f.Println(symbol0, symbol1, symbol2, symbol3, symbol4, symbol5, symbol6, symbol7)
	//f.Println(
	//	strconv.Itoa(int(symbol0)), strconv.Itoa(int(symbol1)), strconv.Itoa(int(symbol2)),
	//	strconv.Itoa(int(symbol3)), strconv.Itoa(int(symbol4)), strconv.Itoa(int(symbol5)),
	//	strconv.Itoa(int(symbol6)), strconv.Itoa(int(symbol7)))
	//f.Println(
	//	string(symbol0), string(symbol1), string(symbol2), string(symbol3),
	//	string(symbol4), string(symbol5), string(symbol6), string(symbol7))

	//var (
	//	name string = "Nurlan"
	//	age  int    = 25
	//)
	//
	//f.Println(name)
	//f.Println(age)

	var a uint16
	f.Scan(&a)
	f.Print(a * a)
}
