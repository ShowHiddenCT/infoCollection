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
	version, _ := host.KernelVersion()
	platform, family, version, _ := host.PlatformInformation()
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
