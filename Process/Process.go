package Process

import (
	"awesomeProject/Data"
	"github.com/shirou/gopsutil/process"
	"strconv"
)

/*
	采集进程信息
	ProcessId 		进程号
	ProcessName 	进程名
	ProcessMenm 	占用资源
*/
func GetProcessInfo() []Data.Process {
	pi, _ := process.Pids()
	ProcessArr := []Data.Process{}
	Process := Data.Process{}
	for i := 0; i < len(pi); i++ {
		p, _ := process.NewProcess(pi[i])
		strPid := strconv.FormatInt(int64(pi[i]), 10)
		pm, _ := p.MemoryPercent()
		pn, _ := p.Name()
		Process.ProcessId = strPid
		Process.ProcessName = pn
		Process.ProcessMem = pm

		ProcessArr = append(ProcessArr, Process)
	}
	return ProcessArr
}
