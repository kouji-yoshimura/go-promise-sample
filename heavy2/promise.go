package heavy2

type Result struct {
	Value interface{}
	Error error
}

func AsyncTest1() <-chan Result {
	r := make(chan Result)
	go func() {
		defer close(r)
		res, err := Test1()
		r <- Result{
			Value: res,
			Error: err,
		}
	}()
	return r
}

func AsyncTest2(str string) <-chan Result {
	r := make(chan Result)
	go func(s string) {
		defer close(r)
		res, err := Test2(s)
		r <- Result{
			Value: res,
			Error: err,
		}
	}(str)
	return r
}
