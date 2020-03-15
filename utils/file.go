package utils

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path"
)

// 解压文件
// compress -- 带解压的文件
// dist -- 加压到目标路径,nill 代表当前路径"./"
//func UnzipTarGz(compress string, dist string) bool {
//	fmt.Printf("unzip compress : %v to %v \n", compress, dist)
//	err := os.MkdirAll(dist, os.ModePerm)
//	if err != nil {
//		panic(err)
//	}
//	// compress read
//	//fr, err := os.Open(compress)
//	//if err != nil {
//	//	panic(err)
//	//}
//	//defer fr.Close()
//	//// gzip read
//	//gr, err := gzip.NewReader(fr)
//	//if err != nil {
//	//	panic(err)
//	//}
//	//defer gr.Close()
//	//// tar read
//	//tr := tar.NewReader(gr)
//	//// 读取文件
//	//for {
//	//	h, err := tr.Next()
//	//	if err == io.EOF {
//	//		break
//	//	}
//	//	if err != nil {
//	//		panic(err)
//	//	}
//	//	// 显示文件
//	//	fmt.Println(h.Name)
//	//	// 打开文件
//	//	fw, err := os.OpenFile(dist + h.Name, os.O_CREATE | os.O_WRONLY, 0644/*os.FileMode(h.Mode)*/)
//	//	if err != nil {
//	//		panic(err)
//	//	}
//	//	defer fw.Close()
//	//	// 写文件
//	//	_, err = io.Copy(fw, tr)
//	//	if err != nil {
//	//		panic(err)
//	//	}
//	//}
//	//fmt.Println("un tar.gz ok")
//	return true
//}


// 解压文件
// srcFilePath -- 带解压的文件
// destDirPath -- 加压到目标路径,nill 代表当前路径"./"
func UnzipTarGz(srcFilePath string, destDirPath string) error {
	fmt.Printf("unzip compress : %v to %v \n", srcFilePath, destDirPath)
	os.Mkdir(destDirPath, os.ModePerm)
	fr, err := os.Open(srcFilePath)
	if err != nil {
		return err
	}
	defer fr.Close()


	gr, err := gzip.NewReader(fr)
	if err != nil {
		return err
	}
	defer gr.Close()
	tr := tar.NewReader(gr)


	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}


		if hdr.Typeflag != tar.TypeDir {
			os.MkdirAll(destDirPath+"/"+path.Dir(hdr.Name), os.ModePerm)
			fw, _ := os.OpenFile(destDirPath+"/"+hdr.Name, os.O_CREATE|os.O_WRONLY, os.FileMode(hdr.Mode))
			if err != nil {
				return err
			}
			_, err = io.Copy(fw, tr)
			if err != nil {
				return err
			}
		}
	}
	return nil
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

func DeleteFile(file string) bool  {
	err := os.Remove(file)
	if err != nil {  // 删除失败
		return false
	} else { // 删除成功
		return true
	}
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
