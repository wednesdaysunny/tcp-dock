package common

import (
	"github.com/golang/crypto/blowfish"
)

const (
	key = "SZKINGDOM"
)

var hexMap = map[string]byte{
	"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9,
	"A": 10, "a": 10, "B": 11, "b": 11, "C": 12, "c": 12, "D": 13, "d": 13, "E": 14,
	"e": 14, "F": 15, "f": 15,
}

func hexStr2Bytes(src string) []byte {
	m, n := 0, 0
	l := len(src) / 2
	rtn := make([]byte, 0)

	for i := 0; i < l; i++ {
		m = i*2 + 1
		n = m + 1
		cha0 := src[i*2 : m]
		cha1 := src[m:n]
		i0, ok := hexMap[cha0]
		i1, ok := hexMap[cha1]
		if ok {

		}
		b0 := byte(i0 << 4)
		rr := byte(b0 | byte(i1))
		rtn = append(rtn, rr)
	}
	return rtn
}

func byte2HexStr([]byte) string {
	return nil
}

func Decrypt(str string) string {
	ct := make([]byte, 0)
	cipher, _ := blowfish.NewCipher([]byte(key))
	hexByteArr := hexStr2Bytes(str)
	cipher.Decrypt(ct, hexByteArr)

	result := string(ct)
	return result
}

func Encrypt(str string) string {
	return nil
}
