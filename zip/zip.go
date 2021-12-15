package zip

import (
	"archive/zip"
	"fmt"
	"github.com/tearsforyears/go-tool/support"
	"io"
	"io/ioutil"
	"os"
)

func zZip(dir string, fileName string) {
	//获取源文件列表
	f, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
	}
	fzip, _ := os.Create(fileName)
	w := zip.NewWriter(fzip)
	defer w.Close()
	for _, file := range f {
		fw, _ := w.Create(file.Name())
		filecontent, err := ioutil.ReadFile(dir + file.Name())
		if err != nil {
			fmt.Println(err)
		}
		n, err := fw.Write(filecontent)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(n)
	}
}
func DeZip(target string, dir string) {
	os.Mkdir(dir, 0777) //创建一个目录

	rc, err := zip.OpenReader(target) //读取zip文件
	if err != nil {
		fmt.Println(err)
	}
	defer rc.Close()
	for _, file := range rc.File {
		println(file.Name)
		rc, _ := file.Open()
		f, _ := support.CreateDirFile(dir + file.Name)
		io.Copy(f, rc)
	}
}
