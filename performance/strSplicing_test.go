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
func sprintfConcat(n int, str string) string {
  s := ""
  for i := 0; i < n; i++ {
    s = fmt.Sprintf("%s%s", s, str)  
  }
  
  return s
}

/* 使用strings.Builder拼接 */
func builderConcat(n int, str string) string {
  var builder strings.Builder
  for i := 0; i < n; i++ {
    builder.WriteString(str)
  }
  
  return builder.String()
}

/* 使用bytes.Buffer拼接 */
func bufferConcat(n int, str string) string {
  buf := new(bytes.Buffer)
  for i := 0; i < n; i++ {
    buf.WriteString(str)
  }
  
  return buf.String()
}

/* 使用[]byte拼接 */
func byteConcat(n int, str string) string {
  buf := make([]byte, 0)
  for i := 0; i < n; i++ {
    buf = append(buf, str...)
  }
  
  return string(buf)
}

/* 使用[]byte拼接 预先申请好容量 */
func preByteConcat(n int, str string) string {
  buf := make([]byte, 0, n*len(str))
  for i := 0; i < n; i++ {
    buf = append(buf, str...)
  }
  
  return string(buf)
}

func preBuilderConcat(n int, str string) string{
	var builder strings.Builder
	builder.Grow(n*len(str))
	for i := 0; i < n; i++ {
		builder.WriteString(str)
	}
	
	return builder.String()
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
func BenchmarkPreBuilderConcat(b *testing.B) { benchmark(b, preBuilderConcat) }

/*
内存分析
go test -bench="Concat$" -benchmem .
goos: windows
goarch: amd64
BenchmarkPlusConcat-12                26          43829646 ns/op        530997337 B/op     10033 allocs/op
BenchmarkSprintfConcat-12             15          83917653 ns/op        834439570 B/op     37531 allocs/op
BenchmarkBuilderConcat-12          13309             92914 ns/op          522224 B/op         23 allocs/op
BenchmarkBufferConcat-12           12430             95019 ns/op          423537 B/op         13 allocs/op
BenchmarkByteConcat-12             13184             90213 ns/op          628721 B/op         24 allocs/op
BenchmarkPreByteConcat-12          25010             49793 ns/op          212992 B/op          2 allocs/op
+和fmt.Sprintf的效率最低 消耗内存最大。

耗时分析
go test -bench .
goos: windows
goarch: amd64
BenchmarkPlusConcat-12                25          43580028 ns/op
BenchmarkSprintfConcat-12             13          79326492 ns/op
BenchmarkBuilderConcat-12          12974             89454 ns/op
BenchmarkBufferConcat-12           10000            102043 ns/op
BenchmarkByteConcat-12             10000            100989 ns/op
BenchmarkPreByteConcat-12          24915             53024 ns/op

PreByteConcat方式预先分配了内存 所以性能最好
综合考虑 一般推荐使用strings.Builder
string.Builder也提供了预先分配内存的方式
通过Builder.Grow(len(str))预先分配内存后性能变好
BenchmarkPreBuilderConcat-12               23454             51127 ns/op
BenchmarkPreBuilderConcat-12               23865             50592 ns/op          106496 B/op          1 allocs/op
*/

/* 分析 */
/* 
strings.Builder和+n比较
差距巨大 主要是两者内存分配方式不一样 
"+" 拼接两个字符，按照新字符串的大小开辟一段新的空间 总共分配了10 + 10 *2 + 10 *3......+10*10000 = 500M
"strings.Builder" 内存申请有对应策略
*/

func TestBuilderConcat(t *testing.T) {
	var str = randomString(10)
	var builder strings.Builder
	cap := 0
	for i := 0; i < 10000; i++ {
		if builder.Cap() != cap {
			fmt.Print(builder.Cap(), " ")
			cap = builder.Cap()
		}
			builder.WriteString(str)
	}
}

/*
测试结果
go test -run="TestBuilderConcat" . -v
=== RUN   TestBuilderConcat
16 32 64 128 256 512 1024 2048 2688 3456 4864 6144 8192 10240 13568 18432 24576 32768 40960 57344 73728 98304 122880 --- PASS: TestBuilderConcat (0.00s)
0-2048字节：内存分配按照两倍去申请。 比如拼接第二个字节的时候，直接申请了32字节，拼接第三个的时候就无需申请内存了。
2048往后，以640字节大小申请？？？ --- 这是怎么决策的
*/

/*
strings.Builder和bytes.Buffer比较
性能略高 主要因为bytes.Buffer最后将[]Byte转换为string，重新申请了一块内存，而string.Builder是直接强转成字符串返回；如下：
func (b *Builder) String() string {
	return *(*string)(unsafe.Pointer(&b.buf))
}
*/
