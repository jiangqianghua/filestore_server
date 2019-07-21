package main

import (
	"filestore_server/handler"
	"fmt"
	"net/http"
)

func main() {
	// 静态文件处理
	http.Handle("/static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	//http://127.0.0.1:8080/file/upload
	http.HandleFunc("/file/upload", handler.UploadHandler)
	http.HandleFunc("/file/upload/suc", handler.UploadSucHandler)
	//http://127.0.0.1:8080/file/meta?filehash=a0cef7662ea8880e4a6c2792557428b5c2159816
	http.HandleFunc("/file/meta", handler.GetFileMetaHandler)
	//http://127.0.0.1:8080/file/download?filehash=a0cef7662ea8880e4a6c2792557428b5c2159816
	http.HandleFunc("/file/download", handler.DownloadHandler)
	http.HandleFunc("/file/update", handler.FileMetaUpdateHandler)
	http.HandleFunc("/file/delete", handler.FileDeleteHandler)
	http.HandleFunc("/file/query", handler.FileQueryHandler)
	http.HandleFunc("/file/fastupload", handler.HTTPInterceptor(handler.TryFastUploadHandler))
	//http://127.0.0.1:8080/user/signup
	http.HandleFunc("/user/signup", handler.SignupHandler)
	//http://127.0.0.1:8080/user/signin
	http.HandleFunc("/user/signin", handler.SigninHandler)
	//http.HandleFunc("/user/info", handler.UserInfoHandler)
	// 添加HttpInterceptor拦截器
	http.HandleFunc("/user/info", handler.HTTPInterceptor(handler.UserInfoHandler))

	// 分块上传接口
	http.HandleFunc("/file/mpupload/init", handler.HTTPInterceptor(handler.InitialMultippartUploadHandler))
	http.HandleFunc("/file/mpupload/uppart", handler.HTTPInterceptor(handler.UploadPartHandler))
	http.HandleFunc("/file/mpupload/complete", handler.HTTPInterceptor(handler.CompleteUploadHandler))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Failed to start server, err:%s", err.Error())
	}
}
