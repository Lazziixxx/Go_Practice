package main

import (
	"fmt"
	"sort"
)

type myString []string

func (p myString) Len() int           { return len(p) }
func (p myString) Less(i, j int) bool { return p[i] > p[j] }
func (p myString) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	stringTest := []string{"a", "d", "c"}

	sort.Sort(myString(stringTest)) //基于自己的类型定义三种方法 从大到小排序
	fmt.Println(stringTest)

	sort.Strings(stringTest)
	fmt.Println(stringTest)

	intTest := []int {1,2,3}
	sort.Ints(intTest)
	fmt.Println(intTest)

}
