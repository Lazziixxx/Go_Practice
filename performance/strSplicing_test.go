/* Go中字符串不可变 拼接字符串实际上是创建了新的字符对象 如果存在大量字符拼接 会严重的影响性能 */
package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

/* 以下是几种实现N个字符串拼接的方式 */

/* 使用+实现 */
func plusConcat(n int, str string) string {
  s := ""
  for i := 0; i < n; i++ {
    s += str
  }
  
  return s
}

/* 使用fmt.Sprintf拼接 */
func sprintfContact(n int, str string) string {
  s := ""
  for i := 0; i < n; i++ {
    s = fmt.Sprintf("%s%s", s, str)  
  }
  
  return s
}

/* 使用strings.Builder拼接 */
func builderContact(n int, str string) string {
  var builder strings.Builder
  for i := 0; i < n; i++ {
    builder.WriteString(str)
  }
  
  return builder.String()
}

/* 使用bytes.Buffer拼接 */
func bufferContact(n int, str string) string {
  buf := new(bytes.Buffer)
  for i := 0; i < n; i++ {
    buf.WriteString(str)
  }
  
  return buf.String()
}

/* 使用[]byte拼接 */
func byteContact(n int, str string) string {
  buf := make([]byte, 0, n*len(str))
  for i := 0; i < n; i++ {
    buf = append(buf, str...)
  }
  
  return string(buf)
}

func benchmark(b *testing.B, f func(int, string) string) {
  var str = randomString(10) //生成一个长度为10的字符串
  for i := 0; i < b.N; i++ {
    f(10000, str)
  }
}

func BenchmarkPlusConcat(b *testing.B)    { benchmark(b, plusConcat) }
func BenchmarkSprintfConcat(b *testing.B) { benchmark(b, sprintfConcat) }
func BenchmarkBuilderConcat(b *testing.B) { benchmark(b, builderConcat) }
func BenchmarkBufferConcat(b *testing.B)  { benchmark(b, bufferConcat) }
func BenchmarkByteConcat(b *testing.B)    { benchmark(b, byteConcat) }
func BenchmarkPreByteConcat(b *testing.B) { benchmark(b, preByteConcat) }

