package utils

import "os"

/*
	导出文件的工具类
*/

func WriteFile(file string, jsonByte []byte) {
	fp, err := os.OpenFile(file, os.O_RDWR, 0666)
	if err != nil && os.IsNotExist(err) {
		os.Create(file)
		WriteFile(file, jsonByte)
		return
	}
	defer fp.Close()
	_, err = fp.Write(jsonByte)
	if err != nil {
		panic(err)
	}
}
