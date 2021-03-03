package main

import (
	"fmt"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

type Inset struct {
	value int
}

func (c *Inset) String()  string{
	c.value = 1
	return "ok"
}

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)

	c = 0
	var name = "bob"

	/* 基于ByteCounter的指针类型实现了Write方法  因此&c可以视作io.Writer接口类型传给fmt.Fprintf */
	fmt.Fprintf(&c, "hello,%s", name)
	fmt.Println(c)

	var s Inset
	fmt.Println(s.String()) // ok
	fmt.Println(s.value) // 1
	//fmt.Println(Inset{}.String()) //相比35行的 基于Inset变量调用string方法 这里非法，因此不可以在不能寻址的Inset值上调用这个方法
	var k fmt.Stringer = &s
	//var k fmt.Stringer = s //error  仅仅在*Inset上实现了String方法
	fmt.Println(k.String()) // ok
}

/*
一个类型如果实现了某个接口类型的所有方法  那么这个类型就实现了这个接口
 */

/*
经常会简要的把一个具体的类型描述成一个特定的接口类型。举个例子，*bytes.Buffer是io.Writer；*os.Files是io.ReadWriter
 */
