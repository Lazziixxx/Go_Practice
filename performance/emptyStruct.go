/*
空结构体不占用内存 通常作为占位符使用
*/

/*********************************************/
/* 用法1：用Map实现Set集合 Map的Value是空结构体 */
type Set map[string]struc{}

func (s Set) Has(Key string) bool {
  _, ok := s[key]
  return ok
}

func (s Set) Add(Key string) {
  s[Key] = struct{}{} 
}

func (s Set) Delete(key string) {
  delet(s, key)
}
/*********************************************/

/*********************************************/
/* 用法2：不发送数据的信道 Channel */
/* channel不发送任何数据  只用来通知子协程执行任务 */
func worker(ch chan struct{}) {
  <- ch
  fmt.Println("do something")
  close(ch)
}

func main() {
  ch := make(chan struct{})
  go worker(ch)
  ch <- struct{}{}
}
/*********************************************/


/*********************************************/
/* 用法3：仅包含方法的结构体 */
type door struct{}

func (d door) Open() {
  ....
}

func (d door) Close() {
  ....
}
/*********************************************/

