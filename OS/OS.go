package OS

//操作系统
import (
	"awesomeProject/Data"
	"github.com/shirou/gopsutil/host"
	//"reflect"
)

/*
	采集操作系统信息
	OsName 操作系统名称
	OsType 操作系统类型
	OsVersion 操作系统版本
*/

func GetOSInfo() Data.Os {
	// 获取版本
	version, _ := host.KernelVersion()
	//获取平台信息
	platform, family, version, _ := host.PlatformInformation()
	//整合信息
	os := Data.Os{
		OsName:    platform,
		OsType:    family,
		OsVersion: version,
	}
	return os
}

//func main() {
//	OsType, family, OsVersion := getOSInfo2()
//	fmt.Printf("%s\n", OsType)
//	fmt.Printf("%s\n", family)
//	fmt.Printf("%s\n", OsVersion)
//}
