package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func printfByType(x interface{}) {
	//当用空接口作为函数接收参数时
	/*
	静态类型是interface {}
	但是动态类型需要使用类型断言去判断
	 */
	switch x.(type) {
	case nil:
		fmt.Println("err")
	case int:
		fmt.Println(x) // int(x)错误
	case bool:
		if x == true {
			// 如果使用if x 报错non-bool x (type interface {}) used as if condition
			//此时x是空接口类型 非bool类型 不能当作true/false使用
			fmt.Println("TRUE")
		} else {
			fmt.Println("FALSE")
		}
	default:
		fmt.Println("Unsupport Type") // int(x)错误
	}
	return
}

func main() {
	var w io.Writer
	w = os.Stdout
	//x.(T) x表示接口类型 T表示类型 断言检查操作对象的动态类型和断言的类型是否匹配
	f := w.(* os.File) //如果成功 返回x的动态值 os.Stdout
	fmt.Println(f)
	//c := w.(* bytes.Buffer)//如果失败 panic
	//panic: interface conversion: io.Writer is *os.File, not *bytes.Buffer

	//断言接口类型 必须实现了接口的所有方法的类型才success
	rw := w.(io.ReadWriter)
	fmt.Println(rw)
	w = new(ByteCounter)
	//rw = w.(io.ReadWriter)
	//基于ByteCounter类型没有实现Read接口 断言失败
	//panic: interface conversion: *main.ByteCounter is not io.ReadWriter: missing method Read
	rw= nil
	//w = rw.(io.Writer)
	//对于nil的断言操作 始终会发生panic
	var w1 io.Writer = os.Stdout
	fmt.Println(w1)
	w1, ok := w1.(*os.File)
	if ok {
		fmt.Println(w1, ok) //如果断言的返回结果有两个
	}
	//常用使用场景
	/*
		if f, ok := w.(*os.File); ok {
			// ...use f...
		}
	 */
	t1, ok :=w1.(*bytes.Buffer)
	fmt.Println(t1, ok) // t1类型是*os.File 去断言*bytes.Buffer类型 失败 ，但此时不会panic
	var byteTest *bytes.Buffer
	var w2 io.Writer = byteTest
	w2, ok =w2.(*bytes.Buffer)// t2类型是*bytes.Buffer类型 成功 但是接口值是nil
	fmt.Println(w2, ok) //

	var a int = 1
	printfByType(a)

	var b *int = nil
	printfByType(b)

	var c interface{}
	printfByType(c)

	var d bool = true
	printfByType(d)

	//e := []int{1,2,3,4}
	//var g interface{} = e
	//h := g[1:3]  //当用空接口承载数组or切片的时候 无法再对空接口进行取切片
	//fmt.Println(h)
}

//空接口的使用
/*
任何值可以赋值给空接口 但是空接口不可以强转回具体类型
 */
