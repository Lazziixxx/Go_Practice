package main

import (
  "github.com/pkg/profile"
  "math/rand"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

/*产生长度为n的string*/
func randomString(n int) string {
  a := make([]byte, n)
  for i := range a {
    a[i] = letterBytes[rand.Intn(len(letterBytes))]
  }
  
  return string(a)
}

func generateString(n int) string {
  s := ""
  for i := 0; i < n; i++ {
    s += randomString(n)
  }
  
  return s
}

func main() {
  //defer profile.Start(profile.MemProfile, profile.MemProfileRate(1)).Stop()
  //defer profile.Start().Stop()
	stopper := profile.Start(profile.CPUProfile, profile.ProfilePath(".")) 
  //分析CPU使用情况 传参数profile.CPUProfile， profile.ProfilePath(".")指定文件生成在当前目录
  //如果要分析内存使用情况，第一个参数profile.MemProfile
	defer stopper.Stop() //在测试程序结束时停止性能分析
  generateString(100)
}
