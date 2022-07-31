package Disk

import (
	"awesomeProject/Data"
	"github.com/shirou/gopsutil/disk"
)

/*
	采集磁盘信息
	Name 				磁盘名
	ReadCount 			读次数
	MergedReadCount 	合并读取计数
	WriteCount 			写次数
	MergedWriteCount 	合并写计数
	ReadBytes 			读字节
	WriteBytes 			写字节
	ReadTime 			读时间
	WriteTime 			写时间
	IopsInProgress
	IoTime
	SerialNumber       	序列码
	Label
*/
func GetDiskInfo() []Data.Disk {
	mapStat, _ := disk.IOCounters()
	var mems = []Data.Disk{}
	var mem = Data.Disk{}
	for _, stat := range mapStat {
		mem.Name = stat.Name
		mem.ReadCount = stat.ReadCount
		mem.MergedWriteCount = stat.MergedWriteCount
		mem.WriteCount = stat.WriteCount
		mem.MergedReadCount = stat.MergedReadCount
		mem.ReadBytes = stat.ReadBytes
		mem.WriteBytes = stat.WriteBytes
		mem.ReadTime = stat.ReadTime
		mem.WriteTime = stat.WriteTime
		mem.IopsInProgress = stat.IopsInProgress
		mem.IoTime = stat.IoTime
		mem.WeightedIO = stat.WeightedIO
		mem.SerialNumber = stat.SerialNumber
		mem.Label = stat.Label

		mems = append(mems, mem)
	}
	return mems
}
