package main

import (
	"fmt"
	"lynzips/src/public"
	"lynzips/src/zipfunc"
	"os"
	"path/filepath"
)

func getFilelist(path string) []public.FlieList {
	var fliemeslice []public.FlieList //var
	err := filepath.Walk(path, func(paths string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		fliemeslice = append(fliemeslice, public.FlieList{Flienames: filepath.Base(paths), Filepaths: paths})
		//println(paths)
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
		return nil
	} else {
		// for _, infos := range fliemeslice {
		// 	//println(infos.flienames)
		// 	//println(infos.filepaths)
		// }

	}
	// lynZip(fliemeslice)
	return fliemeslice
}

func main() {
	//返回所给路径的绝对路径

	path, _ := filepath.Abs("./1.txt")
	fmt.Println(path)
	fslices := getFilelist("D:\\GOPATHS\\src\\lynzips\\zips") //
	zipfunc.LynZip(fslices)
	//unzip("file.zip", "D:\\GOPATHS\\src\\lynzips\\") //./"

}
