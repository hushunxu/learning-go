package main

import "fmt"

func main() {
	s := make([]string, 3)
	//fmt.Println(s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	//fmt.Println(s)
	//fmt.Println(len(s))
	//
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println(s)

	c := make([]string , len(s))
	copy(c, s)
	//fmt.Println("cpy", c)


	//slice 支持通过 slice[low:high] 语法进行“切片”操作。 例如，下面的操作可以得到一个包含元素 s[2]、s[3] 和 s[4] 的 slice。
	//2到4
	l := s[2:5]
	fmt.Println("slice 2:5", l)

	//这个 slice 包含从 s[0] 到 s[5]（不包含 5）的元素。
	l = s[:5]
	fmt.Println("slice :5", l)

	//2到5。这个 slice 包含从 s[2]（包含 2）之后的元素。
	l = s[2:]
	fmt.Println("slice 2:", l)

	t := []string {"g", "h", "i"}
	fmt.Println(t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)
}