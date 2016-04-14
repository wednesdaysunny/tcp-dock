package tcplogic

import (
	"log"
	"testing"
)

func TestTcplogic(t *testing.T) {
	RunDataPack()
}

func RunDataPack() {
	var pack RequestPackage
	pack.Init()
	pack.GetPackHeadDataString(1)
	data_string := pack.GetPackageDataString()
	log.Printf("%04d", len(data_string))
}
