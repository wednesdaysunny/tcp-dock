package tcplogic

import (
	"bytes"
	"fmt"
)

const (
	VERSION_DEFAULT = "KDGATEWAY1.2"
	OPERATOR_ROLE   = "1"
	SOH             = "|"
)

type RequestPackage struct {
	// head
	HeadLength string // CAHR(4),前面以字符‘0’填充 请求包头长度指“请求包头”的字节长度
	DataLength string // CHAR(4),前面以字符‘0’填充 指“请求数据”的长度

	CRC             string // CHAR(8)
	Version         string // 当前协议版本编号(固定串“KDGATEWAY1.2”)
	CustCode        string // VARCHAR(10), 登陆后送，为操作者代码
	OperatorSite    string // VARCHAR(64)
	Branch          string // VARCHAR(6), 登陆后送
	OperatorChannel string // CHAR(1), 见数据字典
	SessionNo       string // VARCHAR(10), 登陆后送

	OperatorRole    string // VARCHAR(20)用户角色 OP_ROLE
	RequestSerialNo string // VARCHAR(20)请求包序列号
	Remark          string // VARCHAR(20)

	// data
	RequestData []string //	请求功能号(VARCHAR(5)) + 请求数据域
}

func (request RequestPackage) Init() {
	request.HeadLength = "0000"
	request.DataLength = "0000"

	request.CRC = ""
	request.Version = VERSION_DEFAULT
	request.CustCode = ""
	request.OperatorSite = ""
	request.Branch = ""
	request.OperatorChannel = ""
	request.SessionNo = ""

	request.OperatorRole = OPERATOR_ROLE
	request.RequestSerialNo = ""
	request.Remark = ""

	request.RequestData = make([]string, 0)
}

func (request RequestPackage) AddRequestData(data_string string) {
	request.RequestData = append(request.RequestData, data_string)
}

func (request RequestPackage) GetPackageDataString() string {
	var buffer bytes.Buffer
	for _, content := range request.RequestData {
		buffer.WriteString(content)
		buffer.WriteString(SOH)
	}
	return buffer.String()
}

func (request RequestPackage) GetPackHeadDataString(data_length int) string {
	request.DataLength = fmt.Sprintf("%04d", data_length)

	var buffer bytes.Buffer
	buffer.WriteString(request.DataLength)
	buffer.WriteString(SOH)
	buffer.WriteString(request.CRC)
	buffer.WriteString(SOH)
	buffer.WriteString(request.Version)
	buffer.WriteString(SOH)
	buffer.WriteString(request.CustCode)
	buffer.WriteString(SOH)
	buffer.WriteString(request.OperatorSite)
	buffer.WriteString(SOH)
	buffer.WriteString(request.Branch)
	buffer.WriteString(SOH)
	buffer.WriteString(request.OperatorChannel)
	buffer.WriteString(SOH)
	buffer.WriteString(request.SessionNo)
	buffer.WriteString(SOH)
	buffer.WriteString(request.OperatorRole)
	buffer.WriteString(SOH)
	buffer.WriteString(request.RequestSerialNo)
	buffer.WriteString(SOH)
	buffer.WriteString(request.Remark)
	buffer.WriteString(SOH)

	head_content := buffer.String()
	head_length := len(head_content) + len(request.HeadLength) + len(SOH)
	request.HeadLength = fmt.Sprintf("%04d", head_length)
	return request.HeadLength + SOH + head_content
}

func (p RequestPackage) GetSendContent() string {
	return p.GetPackHeadDataString() + p.GetPackageDataString()
}

type ResponsePackage struct {
	HeadLength string // CHAR(4),前面以字符‘0’填充
	DataLength string // VARCHAR(n),前面以字符‘0’填充

	CRC     string // CHAR(8)
	Version string // 当前协议版本编号(固定串“KDGATEWAY1.2”)

	Code        string // VARCHAR(10), “0”表示正常
	Msg         string // VARCHAR(200),返回非0则表示交易处理出现某种交易错误或系统错误(见数据字典)
	NextPackage string // CHAR(1),‘0’－无，‘1’－表示有后续包(取后续包发99请求)

	FieldNum string // VARCHAR(10)
	RowNum   string // VARCHAR(10)

	//	RequestFunction     string // VARCHAR(20)
	//	RequestPackageIndex string // VARCHAR(20)

	//	ReservedField string // VARCHAR(20)
}

func (rp *ResponsePackage) Unpack(content string) {
}
