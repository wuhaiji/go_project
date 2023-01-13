package main

import (
	"fmt"
)

type bin int

func (sf bin) String() string {
	return fmt.Sprintf("%b", sf)
}

type Mover interface {
	Move()
}

type Dog struct {
	Name string
}

func (d Dog) Move() {
	fmt.Printf("%v快跑\n", d.Name)
}

func main() {
	//num1 := [...]int{1, 2, 1}
	//num2 := [...]int{1, 2, 3}
	//eq := num1 == num2
	//appendBool := strconv.FormatBool(eq)
	//fmt.Println("" + appendBool)
	//fmt.Println(bin(42))
	//
	//a := [5]int{1, 2, 3, 4, 5}
	//s := a[:] // s := a[low:high]
	//s = append(s, 6)
	//s = append(s, 7)
	//s = append(s, 8)
	//fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))
	//s2 := s[3:10] // 索引的上限是cap(s)而不是len(s)
	//fmt.Printf("s2:%v len(s2):%v cap(s2):%v\n", s2, len(s2), cap(s2))
	//
	//var s1 []int = nil
	//fmt.Printf("s1:%s\n", len(s1) == 0)
	//fmt.Println("map===================================================================================")
	//scoreMap := map[string]int{}
	//scoreMap["zs"] = 90
	//scoreMap["ls"] = 11
	//scoreMap["sd"] = 11
	//fmt.Printf("scoreMap %v\n", scoreMap)
	//if v, ok := scoreMap["ww"]; ok {
	//	fmt.Printf("v %v\n", v)
	//} else {
	//	fmt.Printf("查无此人\n")
	//}
	//if v, ok := scoreMap["ls"]; ok {
	//	fmt.Printf("v %v\n", v)
	//} else {
	//	fmt.Printf("查无此人\n")
	//}
	//for k, v := range scoreMap {
	//	fmt.Printf("%v:%v\n", k, v)
	//}
	//delete(scoreMap, "ls")
	//fmt.Println("----")
	//for k, v := range scoreMap {
	//	fmt.Printf("%v:%v\n", k, v)
	//}
	//
	//var sliceMap = make(map[string][]string, 3)
	//fmt.Println(sliceMap)
	//fmt.Println("after init")
	//key := "中国"
	//value, ok := sliceMap[key]
	//if !ok {
	//	value = make([]string, 0, 2)
	//}
	//value = append(value, "北京", "上海")
	//sliceMap[key] = value
	//fmt.Println(sliceMap)

	var dog Mover = &Dog{"旺财"}
	v, ok := dog.(*Dog)
	if ok {
		v.Name = "富贵"
		fmt.Printf("类型断言成功 %v", v)

	} else {
		fmt.Printf("类型断言失败")
	}
}
func justifyType(x interface{}) {

	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string, value is %v \n", v)
	case int:
		fmt.Printf("x is a int, value is %v \n", v)
	case bool:
		fmt.Printf("x is a bool, value is %v \n", v)
	default:
		fmt.Printf("not support type!\n")
	}
}
