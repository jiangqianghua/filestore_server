package handler

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

// 上传
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UploadHandler....")
	if r.Method == "GET" {
		// 返回上传的html
		data, err := ioutil.ReadFile("./static/view/index.html")
		if err != nil {
			io.WriteString(w, "internel server error")
			return
		}
		// 返回index.html页面
		io.WriteString(w, string(data))
	} else if r.Method == "POST" {
		// 接受文件流及存储到本地目录
		file, head, err := r.FormFile("file")
		if err != nil {
			fmt.Println("Failed to get data, err:%s", err.Error())
			return
		}

		defer file.Close()

		// 创建本地文件
		newFile, err := os.Create("./tmp/" + head.Filename)
		if err != nil {
			fmt.Println("Failed to create file, err:%s", err.Error())
			return
		}
		defer newFile.Close()

		// 把上传的文件拷贝到创建的文件中
		_, err = io.Copy(newFile, file)
		if err != nil {
			fmt.Println("Failed to save data into file, err:%s", err.Error())
			return
		}

		// 返回消息
		http.Redirect(w, r, "/file/upload/suc", http.StatusFound)
	}
}

func UploadSucHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Upload finished!")
}
