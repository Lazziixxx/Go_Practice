package main

import "fmt"

var a int

func main () {
	a = 1
	go func() {
		a = 2
	}()

	a =3
	fmt.Printf("%d\n", a)
}

//go run -race raceDetect.go
/*
WARNING: DATA RACE
Write at 0x0000008135c8 by goroutine 7:
  main.main.func1()
      D:/01 Ability/06 Go/TestProject/sync/raceDetect.go:10 +0x44

Previous write at 0x0000008135c8 by main goroutine:
  main.main()
      D:/01 Ability/06 Go/TestProject/sync/raceDetect.go:13 +0x78

竞争检测
 */
