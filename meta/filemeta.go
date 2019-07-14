package meta

import (
	mydb "filestore_server/db"
)

type FileMeta struct {
	FileSha1 string
	FileName string
	FileSize int64
	Location string
	UploadAt string
}

var fileMetas map[string]FileMeta

func init() {
	fileMetas = make(map[string]FileMeta)
}

// 新增或更新文件元信息
func UploadateFileMate(fmeta FileMeta) {
	fileMetas[fmeta.FileSha1] = fmeta
}

// 通过sha1获取文件对象
func GetFileMeta(fileSha1 string) FileMeta {
	return fileMetas[fileSha1]
}

// 删除元数据
func RemoveFileMeta(fileSha1 string) {
	delete(fileMetas, fileSha1)
}

// 新增和更新元数据到数据库中
func UpdateFileMetaDB(fmeta FileMeta) bool {
	return mydb.OnFileUploadFinished(fmeta.FileSha1, fmeta.FileName, fmeta.FileSize, fmeta.Location)
}
