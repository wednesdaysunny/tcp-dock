package common

type Task interface {
}

type TimeTask struct {
	ProcessTime int
}

type TcpTask struct {
}

type GlobalTask struct {
	TcpChan       chan TcpTask
	TimeChan      chan TimeTask
	ReconnectChan chan int
}

func (task *GlobalTask) InitTask(tcp_count, time_count, reconnect_count int) {
	task.TcpChan = make(chan TcpTask, tcp_count)
	task.TimeChan = make(chan TimeTask, time_count)
	task.ReconnectChan = make(chan int, reconnect_count)
}
