/* error处理的最佳实践 */
package main

func main() {
  /* 基本用法 */
  /* fmt.Errorf 返回一个被包装的error */
  err1 := errors.New("error1")
  err2 := fmt.Errorf("error2:[%w]", err1)
  err3 := fmt.Errorf("error3:[%w]", err2)
  fmt.Println(err3)
   
}

https://studygolang.com/articles/23454
