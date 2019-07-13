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
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Failed to start server, err:%s", err.Error())
	}
}
