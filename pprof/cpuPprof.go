package main

import (
	"math/rand"
	"time"
)

func generate(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}

func bubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums)-i; j++ {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
}

func main() {
  f, _ := os.OpenFile("cpu.pprof", os.O_CREATE|os.O_RDWR, 0644)
  defer f.Close()
  pprof.StartCPUProfile(f)
  defer pprof.StopCPUProfile()
	n := 10
	for i := 0; i < 5; i++ {
		nums := generate(n)
		bubbleSort(nums)
		n *= 10
	}
}

/*
go build cpuProfile.go
go tool pprof -http=:9999 cpu.pprof 网页中查看图形分析数据， 需要安装Graphviz并设置Goland的环境变量
也可以通过go tool pprof cpu.pprof查看
go tool pprof cpu.pprof
Type: cpu
Time: Mar 8, 2021 at 7:37pm (CST)
Duration: 11.78s, Total samples = 11.57s (98.18%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 11.54s, 99.74% of 11.57s total
Dropped 1 node (cum <= 0.06s)
      flat  flat%   sum%        cum   cum%
    11.54s 99.74% 99.74%     11.57s   100%  main.bubbleSort
         0     0% 99.74%     11.57s   100%  main.main
         0     0% 99.74%     11.57s   100%  runtime.main
------------------------------------------------------------------
*/
