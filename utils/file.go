package utils

import (
	"fmt"
	"os"
)

// 解压文件
// file -- 带解压的文件
// dist -- 加压到目标路径,nill 代表当前路径"./"
func UnzipTarGz(file string, dist string) bool {
	fmt.Printf("unzip file : %v to %v \n", file, dist)
	err := os.MkdirAll(dist, os.ModePerm)
	if err != nil {
		panic(err)
	}
	// file read
	//fr, err := os.Open(file)
	//if err != nil {
	//	panic(err)
	//}
	//defer fr.Close()
	//// gzip read
	//gr, err := gzip.NewReader(fr)
	//if err != nil {
	//	panic(err)
	//}
	//defer gr.Close()
	//// tar read
	//tr := tar.NewReader(gr)
	//// 读取文件
	//for {
	//	h, err := tr.Next()
	//	if err == io.EOF {
	//		break
	//	}
	//	if err != nil {
	//		panic(err)
	//	}
	//	// 显示文件
	//	fmt.Println(h.Name)
	//	// 打开文件
	//	fw, err := os.OpenFile(dist + h.Name, os.O_CREATE | os.O_WRONLY, 0644/*os.FileMode(h.Mode)*/)
	//	if err != nil {
	//		panic(err)
	//	}
	//	defer fw.Close()
	//	// 写文件
	//	_, err = io.Copy(fw, tr)
	//	if err != nil {
	//		panic(err)
	//	}
	//}
	//fmt.Println("un tar.gz ok")
	return true
}

// 移动文件夹
// source  -- 源头
// dist    -- 目标路径
func MoveDocument(source string, dist string) bool {

	return true
}

//判断文件夹是否存在
func IsFileExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		fmt.Println(err)
		return false
	}
	return true
}

//func MkdirAll(path string){
//	//递归创建文件夹
//	err := os.MkdirAll("./test/1/2", os.ModePerm)
//	if err != nil{
//		fmt.Println(err)
//		return
//	}
//	dirs := []string{"./test/1", "./test/2", "./test/1.txt"}
//	for _, v := range dirs{
//		if isExist(v){
//			fmt.Printf("%s is exist!", v)
//		}else{
//			fmt.Printf("%s is not exist!", v)
//		}
//	}
//}
