package Data

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"syscall"
)

/*
	结构体嵌套
	Application 	应用信息
	Cpu 			CPU信息
	Os 				操作系统信息
	Network 		网络信息
	FireWall 		防火墙状态和规则
	Disk 			磁盘
	Memory 			内存
	Device 			设备信息
	Process 		进程信息
	Service 		服务
*/

// 最终生成的结构体
type Windows struct {
	Application []Application `json:"Application"`
	Cpu         Cpu           `json:"Cpu"`
	Os          Os            `json:"Os"`
	Network     []Network     `json:"Network"`
	FireWall    FireWall      `json:"FireWall"`
	Disk        []Disk        `json:"Disk"`
	Memory      Memory        `json:"Memory"`
	Device      host.InfoStat `json:"Device"`
	Process     []Process     `json:"Process"`
	Service     []Service     `json:"Service"`
}

/*
	采集应用信息
	DisplayIcon 		Icon位置
	DisplayName 		应用名
	DisplayVersion 		版本
	InstallLocation 	下载地址
	Publisher 			发行者
	UninstallString 	卸载命令
*/
type Application struct {
	DisplayIcon     string `json:"displayIcon"`
	DisplayName     string `json:"displayName"`
	DisplayVersion  string `json:"displayVersion"`
	InstallLocation string `json:"installLocation"`
	Publisher       string `json:"publisher"`
	UninstallString string `json:"uninstallString"`
}

/*
	采集CPU信息
	Info  			cpu的基本信息
	LogicalCount  	cpu逻辑核心数量
	PhysicalCount 	cpu物理核心数量
	Usage			cpu利用率
	Time			cpu时间信息
*/
type Cpu struct {
	Info          []cpu.InfoStat  `json:"info"`
	LogicalCount  int             `json:"logicalCount"`
	PhysicalCount int             `json:"physicalCount"`
	Usage         []float64       `json:"usage"`
	Time          []cpu.TimesStat `json:"time"`
}

/*
	采集操作系统信息
	OsName 		操作系统名称
	OsType 		操作系统类型
	OsVersion 	操作系统版本
*/
type Os struct {
	OsName    string `json:"osName"`
	OsType    string `json:"osType"`
	OsVersion string `json:"osVersion"`
}

/*
	采集网络信息
	Name  		接口名
	IpAddress 	ip地址
	IpMask 		子网掩码
	Gateway 	默认网关
*/
type Network struct {
	Name      string                  `json:"name"`
	IpAddress syscall.IpAddressString `json:"ipAddress"`
	IpMask    syscall.IpMaskString    `json:"ipMask"`
	Gateway   syscall.IpAddressString `json:"gateway"`
}

/*
	采集防火墙信息
	FireWallState 防火墙状态
	FireWallRules 防火墙规则
*/
type FireWall struct {
	FireWallState FireWallState   `json:"FireWallState"`
	FireWallRules []FireWallRules `json:"FireWallRules"`
}

/*
	防火墙状态
	ConfigFile 				配置文件
	OperationMode 			操作模式
	ExceptionMode 			例外模式
	BroadcastMulticastMode 	多播/广播相应模式
	NotificationMode 		通知模式
	GroupPolicyVersion 		组策略模式
	RemoteAdminMode 		远程管理模式
*/
type FireWallState struct {
	ConfigFile             string `json:"ConfigFile"`
	OperationMode          string `json:"OperationMode"`
	ExceptionMode          string `json:"ExceptionMode"`
	BroadcastMulticastMode string `json:"BroadcastMulticastMode"`
	NotificationMode       string `json:"NotificationMode"`
	GroupPolicyVersion     string `json:"GroupPolicyVersion"`
	RemoteAdminMode        string `json:"RemoteAdminMode"`
}

/*
	防火墙规则
	RulerName 		规则名称
	Enabled	 		是否启用
	Direction 		方向
	ConfigFile 		配置文件
	Grouping 		分组
	LocalIP 		本地IP
	LongRangeIP 	远程IP
	Agreement 		协议
	LocalPort 		本地端口
	LongRangePort 	远程端口
	EdgeTraversal 	边缘遍历
	Operation 		操作
*/
type FireWallRules struct {
	RulerName     string `json:"RulerName"`
	Enabled       string `json:"Enabled"`
	Direction     string `json:"Direction"`
	ConfigFile    string `json:"ConfigFile"`
	Grouping      string `json:"Grouping"`
	LocalIP       string `json:"LocalIP"`
	LongRangeIP   string `json:"LongRangeIP"`
	Agreement     string `json:"Agreement"`
	LocalPort     string `json:"LocalPort"`
	LongRangePort string `json:"LongRangePort"`
	EdgeTraversal string `json:"EdgeTraversal"`
	Operation     string `json:"Operation"`
}

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
	SerialNumber  		序列码
	Label
*/
type Disk struct {
	Name             string `json:"Name"`
	ReadCount        uint64 `json:"ReadCount"`
	MergedReadCount  uint64 `json:"MergedReadCount"`
	WriteCount       uint64 `json:"WriteCount"`
	MergedWriteCount uint64 `json:"MergedWriteCount"`
	ReadBytes        uint64 `json:"ReadBytes"`
	WriteBytes       uint64 `json:"WriteBytes"`
	ReadTime         uint64 `json:"ReadTime"`
	WriteTime        uint64 `json:"WriteTime"`
	IopsInProgress   uint64 `json:"IopsInProgress"`
	IoTime           uint64 `json:"IoTime"`
	WeightedIO       uint64 `json:"WeightedIO"`
	SerialNumber     string `json:"SerialNumber"`
	Label            string `json:"Label"`
}

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
type Memory struct {
	Total       uint64  `json:"total"`
	Used        uint64  `json:"used"`
	Free        uint64  `json:"free"`
	UsedPercent float64 `json:"usedPercent"`
	PgIn        uint64  `json:"pgin"`
	PgOut       uint64  `json:"pgout"`
	PgFault     uint64  `json:"pgfault"`
}

/*
	采集进程信息
	ProcessId 		进程号
	ProcessName 	进程名
	ProcessMenm 	占用资源
*/
type Process struct {
	ProcessId   string  `json:"ProcessId"`
	ProcessName string  `json:"ProcessName"`
	ProcessMem  float32 `json:"ProcessMem"`
}

/*
	采集服务信息
	ServiceName 		服务名称
	DisplayName 		展示名称
	Type 				类型
	State 				状态
	Win32ExitCode 		win32退出代码
	ServiceExitCode 	服务退出代码
	CheckPoint 			检测点
	WaitHint
*/
type Service struct {
	ServiceName     string `json:"ServiceName"`
	DisplayName     string `json:"DisplayName"`
	Type            string `json:"Type"`
	State           string `json:"State"`
	Win32ExitCode   string `json:"Win32ExitCode"`
	ServiceExitCode string `json:"ServiceExitCode"`
	CheckPoint      string `json:"CheckPoint"`
	WaitHint        string `json:"WaitHint"`
}
