package Application

import (
	"awesomeProject/Data"
	"golang.org/x/sys/windows/registry"
	"strconv"
	"sync"
)

/*
	采集应用信息
	DisplayIcon Icon位置
	DisplayName 应用名
	DisplayVersion 版本
	InstallLocation 下载地址
	Publisher 发行者
	UninstallString 卸载命令
*/

func QueryWindowsExe() []Data.Application {
	ApplicationSli := []Data.Application{}
	queryKey := func(w *sync.WaitGroup, startKey registry.Key, res *[]Data.Application) {
		defer w.Done()
		var queryPath string
		var query64Path string
		if strconv.IntSize == 64 {
			query64Path = "Software\\WOW6432Node\\Microsoft\\Windows\\CurrentVersion\\Uninstall"
			queryPath = "Software\\Microsoft\\Windows\\CurrentVersion\\Uninstall"
			kQuery64, err := registry.OpenKey(startKey, query64Path, registry.READ)
			if err != nil {
				return
			}
			keyNames, err := kQuery64.ReadSubKeyNames(0)
			if err != nil {
				return
			}
			//查询出query64Path下面的程序详情，并且添加到Application
			ApplicationSli = getApplication(startKey, keyNames, query64Path)
			*res = append(*res, ApplicationSli...)
		} else {
			queryPath = "Software\\Microsoft\\Windows\\CurrentVersion\\Uninstall"
		}
		k, err1 := registry.OpenKey(startKey, queryPath, registry.READ)
		if err1 != nil {
			return
		}
		// 读取所有子项
		keyNames, err1 := k.ReadSubKeyNames(0)
		if err1 != nil {
			return
		}
		*res = append(*res, getApplication(startKey, keyNames, queryPath)...)
	}
	res := []Data.Application{}
	waitGroup := new(sync.WaitGroup)
	waitGroup.Add(1)
	go queryKey(waitGroup, registry.LOCAL_MACHINE, &res)
	waitGroup.Wait()
	return res
}

// 获得软件各项信息
func getApplication(startKey registry.Key, appName []string, path string) []Data.Application {
	Applications := []Data.Application{}
	for _, value := range appName {
		kQuery64Details, err := registry.OpenKey(startKey, path+"\\"+value, registry.READ)
		if err != nil {
			continue
		}
		displayIcon, _, err := kQuery64Details.GetStringValue("DisplayIcon")
		displayName, v, err := kQuery64Details.GetStringValue("DisplayName")
		displayVersion, _, err := kQuery64Details.GetStringValue("DisplayVersion")
		installLocation, _, err := kQuery64Details.GetStringValue("InstallLocation")
		publisher, _, err := kQuery64Details.GetStringValue("Publisher")
		uninstallString, _, err := kQuery64Details.GetStringValue("UninstallString")
		if v == 0 {
			continue
		}
		softDetails := Data.Application{displayIcon, displayName, displayVersion, installLocation, publisher, uninstallString}
		Applications = append(Applications, softDetails)
	}
	return Applications
}

//
//func main() {
//	soft := QueryWindowsExe()
//	for i := 0; i < len(soft); i++ {
//		fmt.Println("ICON:%s,\n Name:%s,\n version:%s,\n Location:%s,\n publisher：%s,\n UninstallString:%s\n---------------------", soft[i].DisplayIcon, soft[i].DisplayName, soft[i].DisplayVersion, soft[i].InstallLocation, soft[i].Publisher, soft[i].UninstallString)
//	}
//}
