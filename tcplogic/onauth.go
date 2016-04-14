package tcplogic

//import (
//	"tcp-dock/pool"
//)

func randomString(length int) string {
	return ""
}

func OnAuth() {
	var request RequestPackage
	request.Init()
	request.AddRequestData("100")
	sign_request_serial = randomString(10)
	request.RequestSerialNo = sign_request_serial
}
