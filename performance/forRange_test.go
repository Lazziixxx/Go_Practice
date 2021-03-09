package main

import "fmt"
import "testing"

/* range的回顾 */
func rangeArrayTest() {
	words := []string{"w", "x", "y", "z"}
	for i, s := range words {
		words = append(words, "test") // 对于words，只会在循环开始前计算一次，循环体中修改切片的长度也不会改变本循环的次数
		fmt.Println(i, s)
	}
	/*
	for i := range words {
		fmt.Println(i, words[i])
	} 
	*/
}

/* map */
func rangeMapTest() {
	m := map[string]int{
		"one":1,
		"two":2,
		"three":3,
	}
	
	for k, v := range m {
		delete(m, "two")
		m["four"] = 4 
		m["five"] = 5
		//根据描述 对map进行range
		//迭代过程中，删除还未迭代到的键值对时，该键值一定不会被迭代到-----经测试，删除了two，range编译map的时候还会打出来，why？
		//在for循环里进行key的删除，如果第一次读到的就是two，那还是会打出来。。。。
		//创建新的键值对，可能被迭代到，也可能不会被迭代---------经测试 符合
		fmt.Printf("%v, %v\n", k, v)
	}
}

/* channel */
func rangeChannelTest() {
	ch := make(chan string)
	go func() {
		ch <- "i"
		ch <- "love"
		ch <- "u"
		close(ch)
	}()
	for n := range ch {
		fmt.Printf("%s ", n)
	}
}

func rangeTestFunc (t *testing.T, f func()){
	f()
}

func TestRangeArray(t *testing.T) { rangeTestFunc(t, rangeArrayTest) }
func TestRangeMap(t *testing.T) { rangeTestFunc(t, rangeMapTest) }
func TestRangeChannel(t *testing.T) { rangeTestFunc(t, rangeChannelTest) }

/* 比较for i:= 0; i <n; i++ 和range的性能区别 */
/* 简单的Int数组，用例较为简单，结论：性能无差别 */

type item struct {
	idx int
	val [4096]byte
}
/* 只遍历下标 */
func BenchmarkForStruct(b *testing.B) {
	var items [1024]item
	for i := 0; i < b.N; i++ {
		length := len(items)
		var tmp int
		for j := 0; j < length; j++ {
			tmp = items[j].idx
		}
		_ = tmp
	}
}

/* 只遍历下标 */
func BenchmarkRangeIdxStruct(b *testing.B) {
	var items [1024]item
	for i := 0; i < b.N; i++ {
		var tmp int
		for j := range items {
			tmp = items[j].idx
		}
		_ = tmp
	}	
}

func BenchmarkRangeStruct(b *testing.B) {
	var items [1024]item
	for i := 0; i < b.N; i++ {
		var tmp int
		for _, itemTmp := range items {
			tmp = itemTmp.idx
		}
		_ = tmp
	}		
}

/*
go test -bench=Struct$ .
BenchmarkForStruct-12            4522668               260 ns/op
BenchmarkRangeIdxStruct-12       4675369               251 ns/op
BenchmarkRangeStruct-12             4450            237567 ns/op
如果只是访问下标，for和range无性能差异
但如果用range访问了struct数组的每个元素 会产生很大的开销
原因：range的本质是拷贝，对于每个struct都拷贝返回
简单验证：可以通过for和range对一个数组每个元素进行加操作，最后dump出数组，发现range中对数组的操作没有生效
*/

func generateItems(n int) []*item {
	items := make([]*item, 0, n)
	for i := 0; i < n; i++ {
		items = append(items, &item{idx: i})
	}
	return items
}

func BenchmarkForPointer(b *testing.B) {
	items := generateItems(1024)
	for i := 0; i < b.N; i++ {
		length := len(items)
		var tmp int
		for k := 0; k < length; k++ {
			tmp = items[k].idx
		}
		_ = tmp
	}
}

func BenchmarkRangePointer(b *testing.B) {
	items := generateItems(1024)
	for i := 0; i < b.N; i++ {
		var tmp int
		for _, item := range items {
			tmp = item.idx
		}
		_ = tmp
	}
}

/*
测试结果
BenchmarkForPointer-12                    923090              1211 ns/op
BenchmarkRangePointer-12                  999000              1228 ns/op
性能无差别（使用指针还可以直接修改对应结构体的值）
所以，如果只访问数组下标，用for和range无差异
如果需要访问数组的值，最好用for，不用range，除非数组元素是指针；否则会影响性能
*/
