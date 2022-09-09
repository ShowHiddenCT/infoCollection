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

#### 1.2采集信息

1.  应用: 软件名、发行者、版本、下载地址
2.  cpu: cpu的基本信息、cpu的核心数量、cpu的利用率、cpu的时间信息
3.  设备(Device): 开机时间、内核版本号、平台信息等等
4.  磁盘(Disk): 磁盘名、读写数、读写时间、序列号等等
5.  防火墙(FireWall): 防火墙状态、防火墙规则
6.  内存(Memory): 全部内存大小、已使用的内存大小、未使用的内存大小、已使用百分比
7.  网络(Network): 接口名、IP地址、子网掩码、默认网关
8.  操作系统(OS): 操作系统名称、类型、版本
9.  进程(Process): 进程名、进程号、占用资源
10. 后台服务(Service): 服务名称、展示名称、类型、状态、win32退出代码、服务退出代码、监测点等 

#### 1.3、数据格式
使用的是**嵌套结构体**

```go
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
```

#### 1.4、使用说明

（1）使用GoLand打开项目，运行main函数，生成windows.json文件。

（2）bin文件夹下有生成的windows.exe，运行windows.exe会在同一文件夹下生成windows.json文件。

### 二、 各项信息结构

### 2.1、Application采集
**函数名**：QueryWindowsExe()

#####  结构体

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
```
##### 实现

使用 golang.org/x/sys/windows/registry 包里的OpenKey函数获取注册表的信息，读取所有子项获得软件的各项信息
##### 返回示例

```json
"Application": [{
    "displayIcon": "F:\\software\\VisualStudio\\Professional\\Common7\\IDE\\devenv.exe",
    "displayName": "Visual Studio Professional 2019",
    "displayVersion": "16.9.31205.134",
    "installLocation": "F:\\software\\VisualStudio\\Professional",
    "publisher": "Microsoft Corporation",
    "uninstallString": "\"C:\\Program Files (x86)\\Microsoft Visual Studio\\Installer\\setup.exe\" uninstall --installPath \"F:\\software\\VisualStudio\\Professional\""
}, {
    "displayIcon": "C:\\Program Files (x86)\\Common Files\\Adobe\\AdobeGCClient\\AdobeGCClient.exe",
    "displayName": "Adobe Genuine Service",
    "displayVersion": "7.7.0.35",
    "installLocation": "",
    "publisher": "Adobe Inc.",
    "uninstallString": "\"C:\\Program Files (x86)\\Common Files\\Adobe\\AdobeGCClient\\AdobeCleanUpUtility.exe\""
}]
```
### 2.2、 CPU采集
**函数名**：GetOSInfo()

##### 结构体

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
##### 实现
使用 gopsutil包的cpu子包，该子包提供了获取cpu数据的接口， cpu.Info()获取cpu信息，cpu.Counts()获取cpu核心数量（逻辑和物理），cpu.Percent()获取cpu的使用率， cpu.Times()获得cpu的时间信息。
##### 返回示例

```json
"Cpu": {
    "info": [{
        "cpu": 0,
        "vendorId": "GenuineIntel",
        "family": "205",
        "model": "",
        "stepping": 0,
        "physicalId": "BFEBFBFF000906EA",
        "coreId": "",
        "cores": 8,
        "modelName": "Intel(R) Core(TM) i5-9300H CPU @ 2.40GHz",
        "mhz": 2400,
        "cacheSize": 0,
        "flags": [],
        "microcode": ""
    }],
    "logicalCount": 8,
    "physicalCount": 4,
    "usage": [19.402985074626866, 7.6923076923076925, 10.76923076923077, 18.461538461538463, 23.076923076923077, 10.76923076923077, 12.307692307692308, 10.76923076923077],
    "time": [{
        "cpu": "cpu0",
        "user": 64546.234375,
        "system": 46686.375,
        "idle": 795270.84375,
        "nice": 0,
        "iowait": 0,
        "irq": 16269.328125,
        "softirq": 0,
        "steal": 0,
        "guest": 0,
        "guestNice": 0
    }, {
        "cpu": "cpu1",
        "user": 44972.015625,
        "system": 20418.546875,
        "idle": 841112.5,
        "nice": 0,
        "iowait": 0,
        "irq": 2223.921875,
        "softirq": 0,
        "steal": 0,
        "guest": 0,
        "guestNice": 0
    }, {
        "cpu": "cpu2",
        "user": 101319.59375,
        "system": 25708.625,
        "idle": 779474.84375,
        "nice": 0,
        "iowait": 0,
        "irq": 906.125,
        "softirq": 0,
        "steal": 0,
        "guest": 0,
        "guestNice": 0
    }]
}
```
### 2.3、 Device采集
**函数名**：GetDeviceInfo()

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
##### 实现
使用gopsutil包的host子包，该子包可以获取主机相关信息，如开机时间、内核版本号、平台信息等等，使用host.Info()方法。
##### 返回示例

```json
"Device": {
    "hostname": "DESKTOP-RSCALT0",
    "uptime": 1053672,
    "bootTime": 1657735989,
    "procs": 336,
    "os": "windows",
    "platform": "Microsoft Windows 10 Enterprise",
    "platformFamily": "Standalone Workstation",
    "platformVersion": "10.0.19042 Build 19042",
    "kernelVersion": "10.0.19042 Build 19042",
    "kernelArch": "x86_64",
    "virtualizationSystem": "",
    "virtualizationRole": "",
    "hostid": "49bb5ada-5c12-4ab6-8370-983ec231cf74"
}
```

### 2.4、 Disk采集
**函数名**：GetDiskInfo()

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
##### 实现
使用gopsutil包的disk子包，子包disk用于获取磁盘信息。disk可获取 IO 统计、分区和使用率信息，使用disk.IOCounters()方法获取数据。
##### 返回示例

```json
"Disk": [{
    "Name": "C:",
    "ReadCount": 2686887,
    "MergedReadCount": 0,
    "WriteCount": 7488986,
    "MergedWriteCount": 0,
    "ReadBytes": 112127159808,
    "WriteBytes": 129376860160,
    "ReadTime": 1265,
    "WriteTime": 908,
    "IopsInProgress": 0,
    "IoTime": 0,
    "WeightedIO": 0,
    "SerialNumber": "",
    "Label": ""
}, {
    "Name": "F:",
    "ReadCount": 474974,
    "MergedReadCount": 0,
    "WriteCount": 113694,
    "MergedWriteCount": 0,
    "ReadBytes": 26499145216,
    "WriteBytes": 2372001792,
    "ReadTime": 131,
    "WriteTime": 42,
    "IopsInProgress": 0,
    "IoTime": 0,
    "WeightedIO": 0,
    "SerialNumber": "",
    "Label": ""
}, {
    "Name": "G:",
    "ReadCount": 404031,
    "MergedReadCount": 0,
    "WriteCount": 260683,
    "MergedWriteCount": 0,
    "ReadBytes": 43266053120,
    "WriteBytes": 118305923072,
    "ReadTime": 100,
    "WriteTime": 174,
    "IopsInProgress": 0,
    "IoTime": 0,
    "WeightedIO": 0,
    "SerialNumber": "",
    "Label": ""
}]
```

### 2.5、 FireWall采集
**函数名**：GetFireWallInfo()

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
##### 实现
使用os包的exec子包,使用其中的exec.Command()方法代替cmd运行"netsh firewall show state"命令导出防火墙的状态(这里使用utils工具包里的convert解决中文乱码的问题),再使用Split和Replace等方法调整数据格式,获得详细的防火墙状态;获得防火墙规则同理,不过运行的是"netsh advfirewall firewall show rule name=all"命令
##### 返回示例

```json
"FireWall": {
    "FireWallState": {
        "ConfigFile": "标准",
        "OperationMode": "启用",
        "ExceptionMode": "启用",
        "BroadcastMulticastMode": "启用",
        "NotificationMode": "启用",
        "GroupPolicyVersion": "WindowsDefender防火墙",
        "RemoteAdminMode": "禁用"
    },
    "FireWallRules": [{
        "RulerName": "MicrosoftEdge(mDNS-In)",
        "Enabled": "是",
        "Direction": "入",
        "ConfigFile": "域,专用,公用",
        "Grouping": "MicrosoftEdge",
        "LocalIP": "任何",
        "LongRangeIP": "任何",
        "Agreement": "UDP",
        "LocalPort": "　5353",
        "LongRangePort": "　　任何",
        "EdgeTraversal": "否",
        "Operation": "允许"
    }, {
        "RulerName": "WallpaperEngine：壁纸引擎",
        "Enabled": "是",
        "Direction": "入",
        "ConfigFile": "专用",
        "Grouping": "",
        "LocalIP": "任何",
        "LongRangeIP": "任何",
        "Agreement": "UDP",
        "LocalPort": "　任何",
        "LongRangePort": "　　任何",
        "EdgeTraversal": "否",
        "Operation": "允许"
    }]
```

### 2.6、 Memory采集
**函数名**：GetMemInfo()

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

##### 实现
使用gopsutil包的mem子包,mem包有获取内存信息的函数,使用mem.SwapMemory()函数获取内存信息。
##### 返回示例

```json
"Memory": {
    "total": 26777976832,
    "used": 21954007040,
    "free": 4823969792,
    "usedPercent": 81.98530896391209,
    "pgin": 0,
    "pgout": 0,
    "pgfault": 0
}
```

### 2.7、 Network采集
**函数名**：LocalAddresses()

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
##### 实现
使用gopsutil包的net子包,net.Interfaces()函数会返回系统网络接口的列表,获取数据后再调整格式。
##### 返回示例

```json
"Network": [{
    "name": "以太网",
    "ipAddress": {
        "String": [48, 46, 48, 46, 48, 46, 48, 0, 0, 0, 0, 0, 0, 0, 0, 0]
    },
    "ipMask": {
        "String": [48, 46, 48, 46, 48, 46, 48, 0, 0, 0, 0, 0, 0, 0, 0, 0]
    },
    "gateway": {
        "String": [48, 46, 48, 46, 48, 46, 48, 0, 0, 0, 0, 0, 0, 0, 0, 0]
    }
    }, {
    "name": "本地连接* 1",
    "ipAddress": {
        "String": [48, 46, 48, 46, 48, 46, 48, 0, 0, 0, 0, 0, 0, 0, 0, 0]
    },
    "ipMask": {
        "String": [48, 46, 48, 46, 48, 46, 48, 0, 0, 0, 0, 0, 0, 0, 0, 0]
    },
    "gateway": {
        "String": [48, 46, 48, 46, 48, 46, 48, 0, 0, 0, 0, 0, 0, 0, 0, 0]
    }
    }, {
    "name": "VMware Network Adapter VMnet1",
    "ipAddress": {
        "String": [49, 57, 50, 46, 49, 54, 56, 46, 49, 56, 54, 46, 49, 0, 0, 0]
    },
    "ipMask": {
        "String": [50, 53, 53, 46, 50, 53, 53, 46, 50, 53, 53, 46, 48, 0, 0, 0]
    },
    "gateway": {
        "String": [48, 46, 48, 46, 48, 46, 48, 0, 0, 0, 0, 0, 0, 0, 0, 0]
    }
    }, {
    "name": "WLAN",
    "ipAddress": {
        "String": [49, 57, 50, 46, 49, 54, 56, 46, 48, 46, 49, 48, 55, 0, 0, 0]
    },
    "ipMask": {
        "String": [50, 53, 53, 46, 50, 53, 53, 46, 50, 53, 53, 46, 48, 0, 0, 0]
    },
    "gateway": {
        "String": [49, 57, 50, 46, 49, 54, 56, 46, 48, 46, 49, 0, 0, 0, 0, 0]
    }
    }, {
    "name": "以太网 2",
    "ipAddress": {
        "String": [48, 46, 48, 46, 48, 46, 48, 0, 0, 0, 0, 0, 0, 0, 0, 0]
    },
    "ipMask": {
        "String": [48, 46, 48, 46, 48, 46, 48, 0, 0, 0, 0, 0, 0, 0, 0, 0]
    },
    "gateway": {
        "String": [48, 46, 48, 46, 48, 46, 48, 0, 0, 0, 0, 0, 0, 0, 0, 0]
    }
}]
```
### 2.8、 OS采集
**函数名**：GetOSInfo()

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

##### 实现
使用gopsutil包的host子包,host.PlatformInformation()函数会返回所处平台的信息。
##### 返回示例
```json
"Os": {
    "osName": "Microsoft Windows 10 Enterprise",
    "osType": "Standalone Workstation",
    "osVersion": "10.0.19042 Build 19042"
}
```
### 2.9、 Process采集
**函数名**：GetProcessInfo()

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
##### 实现
使用gopsutil包的process子包,用process.Pids()获取进程ID的数组,通过进程ID创建进程:process.NewProcess(pi[i]),使用p.MemoryPercent()函数获得进程占用资源的百分比,使用 p.Name()获得进程名字。
##### 返回示例

```json
"Process": [{
    "ProcessId": "0",
    "ProcessName": "[System Process]",
    "ProcessMem": 0
}, {
    "ProcessId": "4",
    "ProcessName": "System",
    "ProcessMem": 0
}, {
    "ProcessId": "124",
    "ProcessName": "Registry",
    "ProcessMem": 0
}, {
    "ProcessId": "472",
    "ProcessName": "smss.exe",
    "ProcessMem": 0
}, {
    "ProcessId": "504",
    "ProcessName": "fontdrvhost.exe",
    "ProcessMem": 0
}, {
    "ProcessId": "520",
    "ProcessName": "fontdrvhost.exe",
    "ProcessMem": 0
}, {
    "ProcessId": "548",
    "ProcessName": "svchost.exe",
    "ProcessMem": 0
}]
```
### 2.10、 Service采集
**函数名**：GetServiceInfo()

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

##### 实现
使用os包的exec子包,使用其中的exec.Command()方法代替cmd运行""Sc query state=all""命令导出服务信息(这里使用utils工具包里的convert解决中文乱码的问题),再使用Split和Replace等方法调整数据格式,获得详细的服务信息;
##### 返回示例
```json
"Service": [{
    "ServiceName": "AdobeUpdateService",
    "DisplayName": "AdobeUpdateService",
    "Type": "10WIN32_OWN_PROCESS",
    "State": "4RUNNING",
    "Win32ExitCode": "0(0x0)",
    "ServiceExitCode": "0(0x0)",
    "CheckPoint": "0x0",
    "WaitHint": "0x0"
}, {
    "ServiceName": "AGMService",
    "DisplayName": "AdobeGenuineMonitorService",
    "Type": "10WIN32_OWN_PROCESS",
    "State": "4RUNNING",
    "Win32ExitCode": "0(0x0)",
    "ServiceExitCode": "0(0x0)",
    "CheckPoint": "0x0",
    "WaitHint": "0x0"
}, {
    "ServiceName": "AGSService",
    "DisplayName": "AdobeGenuineSoftwareIntegrityService",
    "Type": "10WIN32_OWN_PROCESS",
    "State": "4RUNNING",
    "Win32ExitCode": "0(0x0)",
    "ServiceExitCode": "0(0x0)",
    "CheckPoint": "0x0",
    "WaitHint": "0x0"
}]
```