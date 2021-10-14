package heavy2

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/pkg/errors"
)

func Test1() (int32, error) {
	dur := time.Duration(rand.Int31n(10))
	fmt.Println("async test1 start. waiting for", dur, "...")
	time.Sleep(time.Second * dur)
	fmt.Println("test1 executed")

	res := int32(rand.Int31n(100))
	if res <= 10 {
	  return 0, errors.New("Test1 Error")
	}

	return res, nil
}

func Test2(s string) (string, error) {
	fmt.Println("async test2 start. waiting for 3s...")
	time.Sleep(time.Second * time.Duration(3))
	fmt.Println("test2 executed")

	if len(s) == 0 {
	  return "", errors.New("Test2 Error")
	}

	return "test2 received " + s, nil
}
