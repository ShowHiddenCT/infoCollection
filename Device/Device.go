package Device

import (
	"github.com/shirou/gopsutil/host"
)

func GetDeviceInfo() host.InfoStat {
	// 获取设备信息
	n, _ := host.Info()
	return *n
}

//func main() {
//	GetDeviceInfo()
//}
