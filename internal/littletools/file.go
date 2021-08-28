package gokupkg

import (
	"fmt"
	"os"
	"log"
	"time"
)

/* 文件操作 */

// os.O_TRUNC 覆盖写入，不加golang则默认追加写入
func WriteToFile(fileName string, content string) error {
   f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
   if err != nil {
      fmt.Println("file create failed. err: " + err.Error())
   } else {
      // offset
      //os.Truncate(filename, 0) //clear
      n, _ := f.Seek(0, os.SEEK_END)
      _, err = f.WriteAt([]byte(content), n)
      fmt.Println("write succeed!")
      defer f.Close()
   }
return err
}

func CreateFile(fileName string){
	f, err := os.Create(fileName)
	fmt.Println("create file success.")
	defer f.Close()
	if err != nil{
		fmt.Println("open file fail")
		return
	}
}

func RemoveFile(fileName string) {
	err := os.Remove(fileName)
	if err != nil {
	log.Fatal(err)
	}
}


// 读取文件大小等信息

func GetFileSize(fileName string) {
    fi,err:= os.Stat(fileName)
    if err == nil {
        fmt.Println("file size is ",fi.Size(),err)
    }
}


func GetFileInfos(fileName string) {
    fi,err:= os.Stat(fileName)
    if err == nil {
        fmt.Println("name:",fi.Name())
        fmt.Println("size:",fi.Size())
        fmt.Println("is dir:",fi.IsDir())
        fmt.Println("mode::",fi.Mode())
        fmt.Println("modTime:",fi.ModTime())
    }
}
