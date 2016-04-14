package main

//服务器端
import (
	"fmt"
	"log"
	"net" //支持通讯的包
	"time"
)

//开始服务器
func startServer() {
	//连接主机、端口，采用ｔｃｐ方式通信，监听７７７７端口
	listener, err := net.Listen("tcp", "localhost:7777")
	checkError1(err)
	fmt.Println("建立成功!")
	for {
		//等待客户端接入
		conn, err := listener.Accept()
		checkError1(err)
		//开一个goroutines处理客户端消息，这是golang的特色，实现并发就只go一下就好
		go doServerStuff(conn)
	}
}

//处理客户端消息
func doServerStuff(conn net.Conn) {
	nameInfo := make([]byte, 512) //生成一个缓存数组
	_, err := conn.Read(nameInfo)
	fmt.Println("-----29", string(nameInfo))
	checkError1(err)

	for {
		time.Sleep(time.Second * 50)
		buf := make([]byte, 512)
		_, err := conn.Read(buf) //读取客户机发的消息
		flag := checkError1(err)
		if flag == 0 {
			break
		}
		fmt.Println(string(buf)) //打印出来
		conn.Write([]byte("hellohello"))
	}
}

//检查错误
func checkError1(err error) int {
	if err != nil {
		if err.Error() == "EOF" {
			//fmt.Println("用户退出了")
			return 0
		}
		log.Fatal("an error!", err.Error())
		return -1
	}
	return 1
}

func main() {
	//开启服务
	startServer()
}
