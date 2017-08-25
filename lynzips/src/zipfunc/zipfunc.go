package zipfunc

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	// "strconv"
	"lynzips/src/public"
	"time"
)

/*
函数名：LynZip
参数：slices
描述：压缩指定文件列表，用时间戳命名压缩文件名，压缩完成
	  删除源文件
*/
func LynZip(slices []public.FlieList) {
	buf := new(bytes.Buffer)

	w := zip.NewWriter(buf)

	var files = slices

	for _, file := range files {

		fmt.Println(file.Flienames)
		fziplist, err := w.Create(file.Flienames)
		if err != nil {
			log.Fatal(err)
		}

		filecontent, err := ioutil.ReadFile(file.Filepaths)
		if err != nil {
			fmt.Println(err)
		}

		_, err = fziplist.Write([]byte(filecontent)) //file.Body
		if err != nil {
			log.Fatal(err)
		}
	}

	err := w.Close()
	if err != nil {
		log.Fatal(err)
	}

	timestamp := time.Now().Unix()

	fmt.Println(timestamp)

	//格式化为字符串,tm为Time类型

	tm := time.Unix(timestamp, 0)
	strtime := tm.Format("20060102030405") + ".zip"

	fliezip, err := os.OpenFile(strtime, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	defer fliezip.Close()
	buf.WriteTo(fliezip)

	fmt.Println("压缩完成！")

	DeleteFile(files)

}

/*
函数名：UnZip
参数：archive,target (压缩文件名，解压路径)
描述：解压到指定文件夹下
*/
func UnZip(archive, target string) error {
	reader, err := zip.OpenReader(archive)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(target, 0755); err != nil {
		return err
	}

	for _, file := range reader.File {
		path := filepath.Join(target, file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(path, file.Mode())
			continue
		}

		fileReader, err := file.Open()
		if err != nil {
			return err
		}
		defer fileReader.Close()

		targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer targetFile.Close()

		if _, err := io.Copy(targetFile, fileReader); err != nil {
			return err
		}
	}

	return nil
}

/*
函数名：DeleteFile
参数：slices
描述：删除指定文件
*/
func DeleteFile(slices []public.FlieList) {
	var files = slices

	for _, file := range files {
		err := os.Remove(file.Filepaths) //删除文件test.txt
		if err != nil {
			//如果删除失败则输出 file remove Error!
			fmt.Println("file remove Error!")
			//输出错误详细信息
			fmt.Printf("%s\n", err)
		} else {
			//如果删除成功则输出 file remove OK!
			fmt.Printf("%s remove OK!\n", file.Flienames)
		}

	}

}
