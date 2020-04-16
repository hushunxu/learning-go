package main

import (
	//"math"
	"fmt"
)

const s string = "constant"

func main() {
	//var b bool
	//b := (1 !=0)
	//fmt.Println(b)
	//v2 := (1 == 2)
	//fmt.Println(v2)
	//_, lastName, nickName := GetName()
	//fmt.Println(lastName, nickName)
	//
	//fmt.Println("go" + "lang")
	//fmt.Println("1+1=", 1 + 1)
	//fmt.Println("7.0/3.0 =", 7.0/3.0)
	//fmt.Println(true && false)
	//fmt.Println(true || false)
	//fmt.Println(!true)
	//
	//var a = "init"
	//fmt.Println(a)
	//
	//var b,c int = 1,2
	//fmt.Println(b, c)
	//
	//var d = true
	//fmt.Println(d)
	//
	//var e int
	//fmt.Println(e)
	//
	//f := "short"
	//fmt.Println(f)
	//
	//fmt.Println(s)
	//
	//const n = 100
	//const d = 3e20 / n
	//fmt.Println(d)
	//fmt.Println(int64(n))
	//fmt.Println(math.Sin(n))

	i := 1
	for i<=3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break;
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

func GetName() (firstName, lastName, nickName string) {
	return "May", "Chan", "Chin21"
}