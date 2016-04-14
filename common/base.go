package common

import (
	"fmt"
	"hash/crc32"
	"log"
	"os"

	"github.com/axgle/mahonia"
)

func EncodeUTF2GBK(content string) string {
	enc := mahonia.NewEncoder("gbk")
	return enc.ConvertString(content)
}

func DecodeGBK2UTF(content string) string {
	enc := mahonia.NewDecoder("gbk")
	return enc.ConvertString(content)
}
func GetCrcContent(content string) string {
	table := crc32.MakeTable(crc32.IEEE)
	check_sum := crc32.Checksum([]byte(content), table)
	crc_content := fmt.Sprintf("%x", check_sum)
	for len(crc_content) < 8 {
		crc_content += "0"
	}
	return crc_content
}

func init() {
	log.Println("Start Test Common..........")
	Test()
	log.Println("Test Common Success..........")
}

func Test() {
	raw_content := "0042|0045|45042736|KDGATEWAY1.2||||C||1|||301|Z|301300019090|C4FA61910C690E1C|thinkive|"
	if "73711fd0" != GetCrcContent(raw_content) {
		os.Exit(2)
		log.Println("Error TestCrc32")
	}
	content := "source_测试"
	if content != DecodeGBK2UTF(EncodeUTF2GBK(content)) {
		os.Exit(2)
		log.Println("Error DecodeGBK2UTF EncodeUTF2GBK")
	}
}
