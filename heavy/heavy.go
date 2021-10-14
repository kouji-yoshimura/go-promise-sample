package heavy

import (
	"fmt"
	"math/rand"
	"time"
)

func Test1() int32 {
	fmt.Println("async test1 start. waiting for 1s...")
	time.Sleep(time.Second * time.Duration(1))
	fmt.Println("test1 executed")

	return int32(rand.Int31n(500))
}

func Test2(s string) string {
	fmt.Println("async test2 start. waiting for 3s...")
	time.Sleep(time.Second * time.Duration(3))
	fmt.Println("test2 executed")

	return "test2 received " + s
}
