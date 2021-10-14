package heavy

func AsyncTest1() <-chan int32 {
	r := make(chan int32)
	go func() {
		defer close(r)
		r <- Test1()
	}()
	return r
}

func AsyncTest2(str string) <-chan string {
	r := make(chan string)
	go func(s string) {
		defer close(r)
		r <- Test2(s)
	}(str)
	return r
}

