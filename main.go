package main

import (
	"filestore_server/handler"
	"fmt"
	"net/http"
)

func main() {
	//http://127.0.0.1:8080/file/upload
	http.HandleFunc("/file/upload", handler.UploadHandler)
	http.HandleFunc("/file/upload/suc", handler.UploadSucHandler)
	//http://127.0.0.1:8080/file/meta?filehash=a0cef7662ea8880e4a6c2792557428b5c2159816
	http.HandleFunc("/file/meta", handler.GetFileMetaHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Failed to start server, err:%s", err.Error())
	}
}
