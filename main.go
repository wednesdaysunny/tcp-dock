package main

import (
	//	"fmt"
	"log"
	"net"
	//	"testing"
	"tcp-dock/common"
	"tcp-dock/pool"
	"time"
)

var (
	RUN_FLAG = true
)

type GlobalData struct {
	TcpChan       chan common.TcpTask
	TimeChan      chan common.TimeTask
	ReconnectChan chan int
}

func CheckReconnect(pool_connect pool.Pool, reconn_chan chan int) {
	need_reconnect := 0
	for RUN_FLAG {
		select {
		case count := <-reconn_chan:
			log.Println("Start Reconnect:", count)
			for i := 0; i < count; i += 1 {
				if err := pool_connect.Reconnect(); err != nil {
					need_reconnect += 1
				}
			}
		case <-time.After(5 * time.Second):
			log.Println("CheckReconnect", need_reconnect)
			if need_reconnect > 0 {
				reconn_chan <- need_reconnect
				need_reconnect = 0
			}

		}
	}
}

func OnProcessTcp(p pool.Pool, gb_data *GlobalData) error {
	//	log.Println("Poollen:", p.Len())
	pc, err := p.Get()
	if err == nil {
		defer pc.Close()
		icount, err1 := pc.Write([]byte("1111111111"))
		if err1 != nil && icount == 0 {
			pc.MarkUnusable()
			gb_data.ReconnectChan <- 1
			log.Println("OnProcessTcp Lost Connect")
		}
	}
	return err
}

func ProcessTcp(process_id int, p pool.Pool, gb_data *GlobalData) {
	log.Println(process_id, "Start ProcessTcp", p, gb_data)
	time.Sleep(time.Second)
	for RUN_FLAG {
		select {
		case task := <-gb_data.TcpChan:
			log.Println(process_id, "TcpTask Process", task)
			if err := OnProcessTcp(p, gb_data); err != nil {
				break
			}
		}
	}
	log.Println(process_id, "Stop ProcessTcp", p, gb_data)
}

type CustomizeSetting struct {
	ServerHost string
	Protocol   string
	InitCount  int
	MaxCount   int
}

func (custom_setting *CustomizeSetting) GetFactory() pool.Factory {
	factory := func() (net.Conn, error) {
		conn, err := net.Dial(custom_setting.Protocol, custom_setting.ServerHost)
		return conn, err
	}
	return factory
}

func (setting *CustomizeSetting) GetChannelPool() pool.Pool {
	factory := setting.GetFactory()
	pool_connect, start_err := pool.NewChannelPool(setting.InitCount, setting.MaxCount, factory)
	for start_err != nil {
		log.Println("NewChannelPool err", start_err)
		time.Sleep(time.Second * 5)
		pool_connect, start_err = pool.NewChannelPool(setting.InitCount, setting.MaxCount, factory)
	}
	return pool_connect
}

func main() {
	ShowVersion()

	SERVER_HOST := "127.0.0.1:7777"
	INIT_CONNECT := 5
	MAX_CONNECT := 10
	custom_setting := new(CustomizeSetting)
	custom_setting.Protocol = "tcp"
	custom_setting.ServerHost = SERVER_HOST
	custom_setting.InitCount = INIT_CONNECT
	custom_setting.MaxCount = MAX_CONNECT
	pool_connect := custom_setting.GetChannelPool()
	defer pool_connect.Close()

	gb_data := new(GlobalData)
	gb_data.TcpChan = make(chan common.TcpTask, 20000)
	gb_data.TimeChan = make(chan common.TimeTask, 20000)
	gb_data.ReconnectChan = make(chan int, MAX_CONNECT)

	go CheckReconnect(pool_connect, gb_data.ReconnectChan)
	for i := 0; i < MAX_CONNECT; i += 1 {
		go ProcessTcp(i, pool_connect, gb_data)
	}

	for RUN_FLAG {
		select {
		case task := <-gb_data.TimeChan:
			log.Println("TimeTask Process", task)
		}
	}
}
