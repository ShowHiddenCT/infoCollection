package FireWall

import (
	"awesomeProject/Data"
	"awesomeProject/utils"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

/*
	采集防火墙信息
	FireWallState 防火墙状态
	FireWallRules 防火墙规则
*/
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
func GetFireWallInfo() Data.FireWall {
	//  获取防火墙状态
	state := getState()
	// 获取防火墙规则
	rules := getRules()
	// 整合成结构体
	firewall := Data.FireWall{
		FireWallState: state,
		FireWallRules: rules,
	}
	return firewall
}

func getState() Data.FireWallState {

	// 用go语言运行windows命令，展示防火墙状态信息
	cmd := exec.Command("netsh", "firewall", "show", "state")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("combined out:\n%s\n", string(out))
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	// 修改编码，防止乱码
	var s string = utils.ConvertByte2String(out, utils.Charset("GB18030"))
	// 调整获取的信息格式，因为是运行windows命令得到的信息，基本表格格式，下面的代码主要是清除换行、空格的字符，获得需要的数据。
	countSplit := strings.Split(s, "\r\n")
	tempstr := []string{}
	for i := 3; i < 10; i++ {
		countSplit[i] = strings.Replace(countSplit[i], "\t", "", -1)
		countSplit[i] = strings.Replace(countSplit[i], " ", "", -1)
		countSplit[i] = strings.Replace(countSplit[i], "\r", "", -1)
		result := strings.Split(countSplit[i], "=")
		tempstr = append(tempstr, result[1])
	}
	state := Data.FireWallState{
		tempstr[0],
		tempstr[1],
		tempstr[2],
		tempstr[3],
		tempstr[4],
		tempstr[5],
		tempstr[6],
	}
	return state

}

func getRules() []Data.FireWallRules {
	/*
		获取防火墙规则信息
		和上面获取防火墙状态的思路一致
		1. 用go语言运行windows命令，获取防火墙规则信息。
		2. 转换编码，防止乱码
		3. 调整格式，因为输出的格式基本为表格，需要清除多余的换行、空格等字符
	*/
	var rules []Data.FireWallRules
	var Rule Data.FireWallRules
	// 运行windows命令
	cmd := exec.Command("netsh", "advfirewall", "firewall", "show", "rule", "name=all")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("combined out:\n%s\n", string(out))
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	// 修改编码
	var s string = utils.ConvertByte2String(out, utils.Charset("GB18030"))

	// 调整格式
	countSplit := strings.Split(s, "规则名称:")
	for i := 1; i < len(countSplit); i++ {
		result := strings.Split(countSplit[i], "\n")
		result[0] = strings.Replace(result[0], " ", "", -1)
		result[0] = strings.Replace(result[0], "\r", "", -1)
		Rule = Data.FireWallRules{
			RulerName: result[0],
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
			switch result2[0] {
			case "已启用":
				Rule.Enabled = result2[1]
			case "方向":
				Rule.Direction = result2[1]
			case "配置文件":
				Rule.ConfigFile = result2[1]
			case "分组":
				Rule.Grouping = result2[1]
			case "本地IP":
				Rule.LocalIP = result2[1]
			case "远程IP":
				Rule.LongRangeIP = result2[1]
			case "协议":
				Rule.Agreement = result2[1]
			case "本地端口":
				Rule.LocalPort = result2[1]
			case "远程端口":
				Rule.LongRangePort = result2[1]
			case "边缘遍历":
				Rule.EdgeTraversal = result2[1]
			case "操作":
				Rule.Operation = result2[1]
			}
		}
		rules = append(rules, Rule)
	}
	return rules
}

//func main() {
//	Info := GetFireWallInfo()
//	fmt.Println(Info)
//}
