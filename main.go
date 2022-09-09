package main

import (
	"awesomeProject/Application"
	"awesomeProject/CPU"
	"awesomeProject/Data"
	"awesomeProject/Device"
	"awesomeProject/Disk"
	"awesomeProject/FireWall"
	"awesomeProject/Memory"
	"awesomeProject/Network"
	"awesomeProject/OS"
	"awesomeProject/Process"
	"awesomeProject/Service"
	"awesomeProject/utils"
	"encoding/json"
	"fmt"
	"github.com/shirou/gopsutil/host"
)

func main() {
	//获取采集信息
	// cpu
	var cpu Data.Cpu = CPU.GetOSInfo()
	//应用
	var application []Data.Application = Application.QueryWindowsExe()
	//操作系统
	var os Data.Os = OS.GetOSInfo()
	// 网络
	var network []Data.Network = Network.LocalAddresses()
	// 防火墙
	var FireWall Data.FireWall = FireWall.GetFireWallInfo()
	// 磁盘
	var Disk []Data.Disk = Disk.GetDiskInfo()
	// 内存
	var Memory Data.Memory = Memory.GetMemInfo()
	// 设备
	var Device host.InfoStat = Device.GetDeviceInfo()
	// 进程
	var Process []Data.Process = Process.GetProcessInfo()
	// 服务
	var Service []Data.Service = Service.GetServiceInfo()

	// 将采集信息整合成结构体
	windowsInfo := Data.Windows{
		Application: application,
		Cpu:         cpu,
		Os:          os,
		Network:     network,
		FireWall:    FireWall,
		Disk:        Disk,
		Memory:      Memory,
		Device:      Device,
		Process:     Process,
		Service:     Service,
	}

	// 生成json文件
	windows, _ := json.Marshal(windowsInfo)
	utils.WriteFile("windows.json", windows)
	fmt.Println("采集结束")
}
