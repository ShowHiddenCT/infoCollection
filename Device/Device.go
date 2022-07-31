package Device

import (
	"github.com/shirou/gopsutil/host"
)

func GetDeviceInfo() host.InfoStat {
	n, _ := host.Info()
	return *n
}

//func main() {
//	collet()
//}
