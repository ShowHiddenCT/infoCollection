package Memory

import (
	"awesomeProject/Data"
	"github.com/shirou/gopsutil/mem"
)

/*
	采集内存信息
	Total 		全部内存大小
	Used 		已使用的内存大小
	Free 		未使用的内存大小
	UsedPercent 已使用百分比
	PgIn 		表征载入页数
	PgOut 		淘汰页数
	PgFault 	缺页错误数
*/
func GetMemInfo() Data.Memory {
	swapMemory, _ := mem.SwapMemory()
	mem := Data.Memory{}
	mem.Total = swapMemory.Total
	mem.Used = swapMemory.Used
	mem.Free = swapMemory.Free
	mem.UsedPercent = swapMemory.UsedPercent
	mem.PgIn = swapMemory.PgIn
	mem.PgOut = swapMemory.PgOut
	mem.PgFault = swapMemory.PgFault
	return mem
}
