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
	//b.N对每个用例都不一样 如果用例能在1s内完成，b.N的值会以1，2，3，5，20，30....增加
	for i := 0; i < b.N; i++ {
		for _, value := range tests {
			if value.out != compareDeflect(value.ins1, value.ins2) {
				b.Error("test failed")
			}
		}
	}
}

/*
go test -bench . 
运行当前目录下的benchmark用例

go test -bench='xxx$' .
运行以xxx结尾的benchmark用例

修改绑核：
go test -bench='xxx$' -cpu=2,4 .
增加-cpu选型 即GOMAXPROCS CPU核数

提升测试准确度:
go test -bench='xxx$' -benchtime=5s .
benchmark默认测试时间是1s 可以使用-benchtime=5s指定测试时间，提升准确度（实际测试时间比5s会长，编译执行销毁需要时间）
-benchtime=100x  也可以指定为测试100次
go test -bench='xxx$' -benchtime=5s -count=3 .
设置进行benchmark的轮数 进行三伦benchmark
*/
