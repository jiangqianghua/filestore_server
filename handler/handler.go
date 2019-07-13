package handler

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
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
	}
}
