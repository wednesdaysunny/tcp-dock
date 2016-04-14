package common

import (
	"fmt"
	"testing"
)

func TestDecrypt(t *testing.T) {
	res := Decrypt("BC42F61D9C531CD3")
	fmt.Println(res)
}

func TestEncrypt(t *testing.T) {
	res := Encrypt("58748506")
	fmt.Printf(res)
}
