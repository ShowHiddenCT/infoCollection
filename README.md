###  一、项目信息

#### 1.1 项目结构


<table>
    <tr>
        <td>Application.go</td>
        <td>采集应用软件信息</td>
    </tr>
    <tr>
        <td>CPU.go</td>
        <td>采集CPU信息</td>
    </tr>
     <tr>
        <td>Data.go</td>
        <td>保存所有数据的结构体</td>
    </tr>
     <tr>
        <td>Device.go</td>
        <td>采集设备信息</td>
    </tr>
     <tr>
        <td>Disk.go</td>
        <td>采集磁盘信息</td>
    </tr>
     <tr>
        <td>FireWall.go</td>
        <td>采集防火墙状态、规则</td>
    </tr>
     <tr>
        <td>Memory.go</td>
        <td>采集内存信息</td>
    </tr>
     <tr>
        <td>Network.go</td>
        <td>采集网络地址、网管等</td>
    </tr>
     <tr>
        <td>OS.go</td>
        <td>采集操作系统的信息</td>
    </tr>
     <tr>
        <td>Process.go</td>
        <td>采集进程信息</td>
    </tr>
     <tr>
        <td>Service.go</td>
        <td>采集系统运行的后台服务信息</td>
    </tr>
     <tr>
        <td>utils.go</td>
        <td>工具类包</td>
    </tr>
     <tr>
        <td>main.go</td>
        <td>主函数，调整数据结构，生成文件</td>
    </tr>
     <tr>
        <td>bin文件夹</td>
        <td>生成的可执行文件</td>
    </tr>
     <tr>
        <td>windows.json</td>
        <td>全部采集的数据，最终生成了json文件</td>
    </tr>
</table>
#### 1.2 数据结构
```go
// 嵌套结构体
type Windows struct {
  Application 	[]Application 	`json:"Application"`
  Cpu     		Cpu      		`json:"Cpu"`
  Os      		Os       		`json:"Os"`
  Network   	[]Network   	`json:"Network"`
  FireWall   	FireWall    	`json:"FireWall"`
  Disk     		[]Disk     		`json:"Disk"`
  Memory    	Memory     		`json:"Memory"`
  Device    	host.InfoStat 	`json:"Device"`
  Process   	[]Process   	`json:"Process"`
  Service   	[]Service   	`json:"Service"`
}
```


### 二、 各项信息结构

#### 2.1、Application采集

##### 结构体
|参数|说明|
|---|----|
|DisplayIcon|Icon位置|
|DisplayName|应用名|
|DisplayVersion|版本|
|InstallLocation|下载地址|
|Publisher|发行者|
|UninstallString|卸载命令| 	
```go
type Application struct {
    DisplayIcon     string `json:"displayIcon"`
    DisplayName     string `json:"displayName"`
    DisplayVersion  string `json:"displayVersion"`
    InstallLocation string `json:"installLocation"`
    Publisher       string `json:"publisher"`
    UninstallString string `json:"uninstallString"`
}

## 2.4、 CPU采集
##### 2.4.1 结构体
|参数|说明|
|---|----|
|Info|cpu的基本信息|
|LogicalCount|cpu逻辑核心数量|
|PhysicalCount|cpu物理核心数量|
|Usage|cpu利用率|
|Time|cpu时间信息|

```go
type Cpu struct {
    Info          []cpu.InfoStat  `json:"info"`
    LogicalCount  int             `json:"logicalCount"`
    PhysicalCount int             `json:"physicalCount"`
    Usage         []float64       `json:"usage"`
    Time          []cpu.TimesStat `json:"time"`
}
```

#### 2.2、 Device采集
##### 结构体
> Device结构体直接使用的是gopsutil包的host子包中编写好的**InfoStat**结构体

> InfoStat 结构体如下：

```go
type InfoStat struct {
    Hostname             string `json:"hostname"`
    Uptime               uint64 `json:"uptime"`
    BootTime             uint64 `json:"bootTime"`
    Procs                uint64 `json:"procs"`           // number of processes
    OS                   string `json:"os"`              // ex: freebsd, linux
    Platform             string `json:"platform"`        // ex: ubuntu, linuxmint
    PlatformFamily       string `json:"platformFamily"`  // ex: debian, rhel
    PlatformVersion      string `json:"platformVersion"` // version of the complete OS
    KernelVersion        string `json:"kernelVersion"`   // version of the OS kernel (if available)
    KernelArch           string `json:"kernelArch"`      // native cpu architecture queried at runtime, as returned by `uname -m` or empty string in case of error
    VirtualizationSystem string `json:"virtualizationSystem"`
    VirtualizationRole   string `json:"virtualizationRole"` // guest or host
    HostID               string `json:"hostid"`             // ex: uuid
}
```

#### 2.3、 Disk采集

##### 结构体

|参数|说明|
|---|----|
|Name|磁盘名|
|ReadCount|读次数|
|MergedReadCount|合并读取计数|
|WriteCount|写次数|
|MergedWriteCount|合并写计数|
|ReadBytes|读字节|
|WriteBytes|写字节|
|ReadTime|读时间|
|WriteTime|写时间|
|SerialNumber|序列码|
|Label|标签|

```go
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
```

#### 2.4、 FireWall采集
##### 结构体
> 结构体使用的是嵌套结构体 分为防火墙状态、防火墙规则

###### (1) 防火墙结构体

|参数|说明|
|---|----|
|FireWallState|防火墙状态|
|FireWallRules|防火墙规则|


```go
type FireWall struct {
    FireWallState FireWallState   `json:"FireWallState"`
    FireWallRules []FireWallRules `json:"FireWallRules"`
}
```
###### (2)防火墙状态结构体

|参数|说明|
|---|----|
|ConfigFile|配置文件|
|OperationMode|操作模式|
|ExceptionMode|例外模式|
|BroadcastMulticastMode|多播/广播相应模式|
|NotificationMode|通知模式|
|GroupPolicyVersion|组策略模式|
|RemoteAdminMode|远程管理模式|
```go
type FireWallState struct {
    ConfigFile             string `json:"ConfigFile"`
    OperationMode          string `json:"OperationMode"`
    ExceptionMode          string `json:"ExceptionMode"`
    BroadcastMulticastMode string `json:"BroadcastMulticastMode"`
    NotificationMode       string `json:"NotificationMode"`
    GroupPolicyVersion     string `json:"GroupPolicyVersion"`
    RemoteAdminMode        string `json:"RemoteAdminMode"`
}
```

###### (3)防火墙规则结构体
|参数|说明|
|---|----|
|RulerName|规则名称|
|Enabled|是否启用|
|Direction|方向|
|ConfigFile|配置文件|
|Grouping|分组|
|LocalIP|本地IP|
|LongRangeIP|远程IP|
|Agreement|协议|
|LocalPort|本地端口|
|LongRangePort|远程端口|
|EdgeTraversal|边缘遍历|
|Operation|操作|	 		

```go
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
```
#### 2.5、 Memory采集

##### 结构体

|参数|说明|
|---|----|
|Total|全部内存大小|
|Used|已使用的内存大小|
|Free|未使用的内存大小|
|UsedPercent|已使用百分比|
|PgIn|表征载入页数|
|PgOut|淘汰页数|
|PgFault|缺页错误数|
```go
type Memory struct {
    Total       uint64  `json:"total"`
    Used        uint64  `json:"used"`
    Free        uint64  `json:"free"`
    UsedPercent float64 `json:"usedPercent"`
    PgIn        uint64  `json:"pgin"`
    PgOut       uint64  `json:"pgout"`
    PgFault     uint64  `json:"pgfault"`
}
```
#### 2.6、 Network采集
##### 结构体
|参数|说明|
|---|----|
|Name|接口名|
|IpAddress|ip地址|
|IpMask|子网掩码|
|Gateway|默认网关|

```go
type Network struct {
    Name      string                  `json:"name"`
    IpAddress syscall.IpAddressString `json:"ipAddress"`
    IpMask    syscall.IpMaskString    `json:"ipMask"`
    Gateway   syscall.IpAddressString `json:"gateway"`
}
```

#### 2.7、 OS采集
##### 结构体
|参数|说明|
|---|----|
|OsName|操作系统名称|
|OsType|操作系统类型|
|OsVersion|操作系统版本|
```go
type Os struct {
    OsName    string `json:"osName"`
    OsType    string `json:"osType"`
    OsVersion string `json:"osVersion"`
}
```
#### 2.8、 Process采集
##### 结构体
|参数|说明|
|---|----|
|ProcessId|进程号|
|ProcessName|进程名|
|ProcessMenm|占用资源|
```go
type Process struct {
    ProcessId   string  `json:"ProcessId"`
    ProcessName string  `json:"ProcessName"`
    ProcessMem  float32 `json:"ProcessMem"`
}
```

#### 2.9、 Service采集
##### 结构体
|参数|说明|
|---|----|
|ServiceName|服务名称|
|DisplayName|展示名称|
|Type|类型|
|State|状态|
|Win32ExitCode|win32退出代码|
|ServiceExitCode|服务退出代码|
|CheckPoint|检测点|
|WaitHint||
```go
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
```