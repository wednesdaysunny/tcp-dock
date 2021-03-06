package common

import (
	"encoding/hex"
	"github.com/golang/crypto/blowfish"
	"strings"
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
	startCnt := 0
	l := len(src) / 2
	rtn := make([]byte, 0)

	for i := 0; i < l; i++ {
		startCnt = i*2 + 1
		cha0 := Substr(src, i*2, 1)
		cha1 := Substr(src, startCnt, 1)
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

func byte2HexStr(byteArr []byte) string {
	return hex.EncodeToString(byteArr)
}

func Decrypt(str string) string {
	ct := make([]byte, 8)
	cipher, _ := blowfish.NewCipher([]byte(key))
	hexByteArr := hexStr2Bytes(str)
	cipher.Decrypt(ct, hexByteArr)

	result := string(ct)
	return result
}

func Encrypt(str string) string {
	cipher, _ := blowfish.NewCipher([]byte(key))
	dest := make([]byte, 8)
	cipher.Encrypt(dest, []byte(str))

	result := byte2HexStr(dest)
	result = strings.ToUpper(result)
	return result
}
