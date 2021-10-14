package main

import (
	"fmt"
	"sync"

	"github.com/kouji-yoshimura/go-wg-expand/heavy"
	"github.com/kouji-yoshimura/go-wg-expand/heavy2"
)


func main() {
	fmt.Println("main start")

	// 1. channel パターン
	fmt.Println("=== channel pattern ===")
	ch1 := heavy.AsyncTest1()
	ch2 := heavy.AsyncTest2("hoge")
	a, b := <-ch1, <-ch2
	fmt.Println("Test1 result: ", a)
	fmt.Println("Test2 result: ", b)

	// 2. channel エラーありパターン
	fmt.Println("\n=== channel with error pattern ===")
	ch3 := heavy2.AsyncTest1()
	ch4 := heavy2.AsyncTest2("")
	a2, b2 := <-ch3, <-ch4
	fmt.Println("Test1 result: ", a2.Value)
	fmt.Println("Test1 error: ", a2.Error)
	fmt.Println("Test2 result: ", b2.Value)
	fmt.Println("Test2 Error: ", b2.Error)

	// 3. channel エラーあり・複数パターン
	fmt.Println("\n=== channel with error multiple pattern ===")
	count := 10
	chList := make([]<-chan heavy2.Result, count)
	for i := range chList {
		ch := heavy2.AsyncTest1()
		chList[i] = ch
	}
	bulkResList := make([]heavy2.Result, count)
	for i, bulkRes := range chList {
		bulkResList[i] = <-bulkRes
	}
	fmt.Println(bulkResList)

	// 4. WaitGroup パターン
	fmt.Println("\n=== WaitGroup pattern ===")
	wg := sync.WaitGroup{}
	wg.Add(2)
	var c int32
	go func() {
		c = heavy.Test1()
		wg.Done()
	}()

	var d string
	go func() {
		d = heavy.Test2("fuga")
		wg.Done()
	}()
	fmt.Println("waiting...")
	wg.Wait()
	fmt.Println("done")

	fmt.Println("Test1 result: ", c)
	fmt.Println("Test2 result: ", d)
}
