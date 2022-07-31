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
	"github.com/shirou/gopsutil/host"
)

func main() {
	var cpu Data.Cpu = CPU.GetOSInfo()
	var application []Data.Application = Application.QueryWindowsExe()
	var os Data.Os = OS.GetOSInfo()
	var network []Data.Network = Network.LocalAddresses()
	var FireWall Data.FireWall = FireWall.GetFireWallInfo()
	var Disk []Data.Disk = Disk.GetDiskInfo()
	var Memory Data.Memory = Memory.GetMemInfo()
	var Device host.InfoStat = Device.GetDeviceInfo()
	var Process []Data.Process = Process.GetProcessInfo()
	var Service []Data.Service = Service.GetServiceInfo()
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
	windows, _ := json.Marshal(windowsInfo)
	utils.WriteFile("windows.json", windows)
}
