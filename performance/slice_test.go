/*
数组操作集
https://ueokande.github.io/go-slice-tricks/
*/

/*切片的性能陷阱*/
package main

import (
	"math/rand"
	"time"
	"testing"
	"runtime"
)
func lastNumBySlice(n []int) []int{
	return n[len(n) - 2:] //返回了原数组的切片
}

func lastNumByCopy(n []int) []int {
	num := make([]int, 0)
	copy(num, n[len(n) - 2:]) //拷贝出最后两个元素并返回新的切片
	return num
}

func generateWithCap(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}

func printMem(t *testing.T) {
	t.Helper()
	var rtm runtime.MemStats
	runtime.ReadMemStats(&rtm)
	t.Logf("%.2f MB", float64(rtm.Alloc)/1024./1024.)
}

func testLastChars(t *testing.T, f func([]int) []int) {
	t.Helper()
	ans := make([][]int, 0)
	for i := 0; i < 100; i++ {
		nums := generateWithCap(128 * 1024)
		ans = append(ans, f(nums))
		//runtime.GC() //显示的调用GC,对于Slice方式来说无效果 对于Copy方式效果明显
	} 
	printMem(t)
	_ = ans
}

func TestLastCharsBySlice(t *testing.T) { testLastChars(t, lastNumBySlice) }
func TestLastCharsByCopy(t *testing.T) { testLastChars(t, lastNumByCopy) }

/*
测试结果
 go test -run=^TestLastChars -v
=== RUN   TestLastCharsBySlice
    slice_test.go:53: 157.44 MB  //底层申请的数组由于一直被slice引用 无法被GC释放
--- PASS: TestLastCharsBySlice (0.32s)
=== RUN   TestLastCharsByCopy
    slice_test.go:54: 2.29 MB //GC释放了generateWithCap申请的内存
--- PASS: TestLastCharsByCopy (0.27s)

手动的在每次append后增加GC查看效果
go test -run=^TestLastChars -v
=== RUN   TestLastCharsBySlice
    slice_test.go:54: 115.84 MB
--- PASS: TestLastCharsBySlice (0.32s)
=== RUN   TestLastCharsByCopy
    slice_test.go:55: 0.22 MB  //使用Copy内存占用更少了
--- PASS: TestLastCharsByCopy (0.27s)
*/
