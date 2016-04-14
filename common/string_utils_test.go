package common

import (
	"fmt"
	"testing"
)

func TestRandomString(t *testing.T) {
	for i := 0; i < 10; i++ {
		res := RandomString(8)
		fmt.Println(res)
	}
}
