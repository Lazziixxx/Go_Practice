package main

import (
	"fmt"
	"reflect"
	"testing"
)

type testslice struct {
	ins1 []string
	ins2 []string
	out bool
}

var tests = []testslice {
	{ins1:[]string{"a","b","c"}, ins2:[]string{"a","b","c"}, out:true},
	{ins1:[]string{"a","b","c"}, ins2:[]string{"a","b","c","d"}, out:false},
	{ins1:[]string{"a","b","c"}, ins2:[]string{"a","b","d"}, out:false},
}

func main() {
	a := []int{99:-1}//最后一个元素是-1 其他都是0
	fmt.Println(a[len(a) - 1])

	//q := [3]int{1,2,3}
	//q = [4]int{1,2,3,4}
	//编译错误 [4]int和[3]int是两种不同的类型

	/*数组清零方法*/
	b := [32]byte{1}
	zero(&b)

	//slice变长 数组定长
	c := []int{}
	c = append(c, 1)
	fmt.Println(c)

	s := []int{1,2,3,4,5}
	s1 := s[1:3] /*[) 右索引位置不取*/
	fmt.Println(s1)

	/* 关于循环比较两个slice的方法 */
	/* go的benchmark使用示例 */

	return
}


func zero(ptr *[32]byte) {
	*ptr = [32]byte{}
}

func compareDeflect(a,b []string) bool {
	return reflect.DeepEqual(a, b)
}

func compareCommon(a,b []string) bool {
	if len(a) != len(b) {
		return false
	}

	if a == nil || b == nil {
		return false
	}

	for i, val := range a {
		if val != b[i] {
			return false
		}
	}

	return true
}

func BenchmarkCompareCommon(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, value := range tests {
			if value.out != compareCommon(value.ins1, value.ins2) {
				b.Error("test failed")
			}
		}
	}
}

func BenchmarkCompareReflect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, value := range tests {
			if value.out != compareDeflect(value.ins1, value.ins2) {
				b.Error("test failed")
			}
		}
	}
}