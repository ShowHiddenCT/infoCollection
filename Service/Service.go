package Service

import (
	"awesomeProject/Data"
	"awesomeProject/utils"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

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

func GetServiceInfo() []Data.Service {
	var Services = []Data.Service{}
	// 使用go语言运行windows命令“Sc query state=all”，获得所有的所有服务信息
	cmd := exec.Command("Sc", "query", "state=all")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("combined out:\n%s\n", string(out))
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	// 转换编码
	var s string = utils.ConvertByte2String(out, utils.Charset("GB18030"))

	// 调整格式，清除空格、换行等字符
	ServiceArr := strings.Split(s, "SERVICE_NAME:")
	for i := 1; i < len(ServiceArr); i++ {
		result := strings.Split(ServiceArr[i], "\n")
		result[0] = strings.Replace(result[0], " ", "", -1)
		result[0] = strings.Replace(result[0], "\r", "", -1)
		service := Data.Service{
			ServiceName: result[0],
		}
		for j := 1; j < len(result); j++ {
			result[j] = strings.Replace(result[j], "\t", "", -1)
			result[j] = strings.Replace(result[j], " ", "", -1)
			result[j] = strings.Replace(result[j], "\r", "", -1)
			result[j] = strings.Replace(result[j], "-", "", -1)
			if result[j] == "" {
				continue
			}
			result2 := strings.Split(result[j], ":")
			if len(result2) < 2 {
				continue
			}
			switch result2[0] {
			case "DISPLAY_NAME":
				service.DisplayName = result2[1]
			case "TYPE":
				service.Type = result2[1]
			case "STATE":
				service.State = result2[1]
			case "WIN32_EXIT_CODE":
				service.Win32ExitCode = result2[1]
			case "SERVICE_EXIT_CODE":
				service.ServiceExitCode = result2[1]
			case "CHECKPOINT":
				service.CheckPoint = result2[1]
			case "WAIT_HINT":
				service.WaitHint = result2[1]
			}
		}
		Services = append(Services, service)
	}
	return Services
}

//func main() {
//	se := GetServiceInfo()
//	fmt.Println(se)
//}
