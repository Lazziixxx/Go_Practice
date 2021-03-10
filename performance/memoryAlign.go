/* Go中的内存对齐 */
package main

type Arg struct {
  num1 int // windows 64位上8字节
  num2 int
}

type Flag struct {
  num1 int16 // windows 64位上2字节
  num2 int32 // windows 64位上4字节
}

func main() {
  fmt.Println(unsafe.Sizeof(Arg{}))  // 16
  fmt.Println(unsafe.Sizeof(Flag{})) // 8
  
  //unsafe.Alignof(xxxx) 返回对齐系数
  fmt.Println(unsafe.Alignof(Args{}))
  fmt.Println(unsafe.Alignof(Flag{}))
  
  type demo1 struct {
    a int8
    b int16
    c int32
  }

  type demo2 struct {
    a int8
    c int32
    b int16
  }
  /* 由于内存对齐的存在，结构体的顺序影响了结构体的大小 跟C一样 */
	fmt.Println(unsafe.Sizeof(demo1{})) // 8
	fmt.Println(unsafe.Sizeof(demo2{})) // 12
  
  /* strcut{}在结构体最后 编译器会做内存对齐 */
  type demo3 struct {
    c int32
    a struct{}
  }

  type demo4 struct {
    a struct{}
    c int32
  }

  type demo5 struct {
    b int64
    c int32
    a struct{}
  }
  
  fmt.Println(unsafe.Sizeof(demo3{})) // 8 结尾有struct{} 按照前一个字段的大小分配padding内存
	fmt.Println(unsafe.Sizeof(demo4{})) // 4  
  fmt.Println(unsafe.Sizeof(demo5{})) // 16  8 + 4 + 4
}
