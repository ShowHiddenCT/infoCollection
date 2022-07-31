package CPU

import (
	"awesomeProject/Data"
	"github.com/shirou/gopsutil/cpu"
	"time"
)

/*
	采集CPU信息
	Info  cpu的基本信息
	LogicalCount  cpu逻辑核心数量
	PhysicalCount cpu物理核心数量
	Usage	cpu利用率
	Time	cpu时间信息
*/
func GetOSInfo() Data.Cpu {

	cpuInfos, _ := cpu.Info()

	/* cpu数量
	CPU 的核数有两种，一种是物理核数，一种是逻辑核数。物理核数就是主板上实际有多少个 CPU，一个物理 CPU 上可以有多个核心，这些核心被称为逻辑核。
	*/
	//逻辑核心数量
	cpuLogicalCount, _ := cpu.Counts(true)
	//物理核心数量
	cpuPhysicalCount, _ := cpu.Counts(false)

	/*cpu利用率
	percpu为false时，获取总的 CPU 使用率，percpu为true时，分别获取每个 CPU 的使用率，返回一个[]float64类型的值。
	*/
	cpuUsage, _ := cpu.Percent(time.Second, true)

	/*cpu有关时间信息
	总 CPU 和 每个单独的 CPU 时间占用情况。传入percpu=false返回总的，传入percpu=true返回单个的。每个 CPU 时间占用情况是一个TimeStat结构
	*/
	cpuTime, _ := cpu.Times(true)

	// 整合成结构体，返回main函数
	cpu := Data.Cpu{
		Info:          cpuInfos,
		LogicalCount:  cpuLogicalCount,
		PhysicalCount: cpuPhysicalCount,
		Usage:         cpuUsage,
		Time:          cpuTime,
	}
	return cpu
}
