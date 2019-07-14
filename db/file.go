package db

import (
	mydb "filestore_server/db/mysql"
	"fmt"
)

// 文件上传完成，保存数据
func OnFileUploadFinished(filehash string, filename string, filesize int64, fileaddr string) bool {
	stmt, err := mydb.DBConn().Prepare(
		"insert ignore into tbl_file(`file_sha1`,`file_name`,`file_size`, `file_addr`, `status`)" +
			" values(?,?,?,?,1)")

	if err != nil {
		fmt.Println("Failed to prepare statement, err:" + err.Error())
		return false
	}

	defer stmt.Close()

	ret, err := stmt.Exec(filehash, filename, filesize, fileaddr)
	if err != nil {
		fmt.Println("Failed to exec, err:" + err.Error())
		return false
	}
	// 获取影响的数据
	if rf, err := ret.RowsAffected(); nil == err {
		if rf <= 0 {
			fmt.Printf("file with hash: %s has been uploaded before", filehash)
		}
		return true
	}
	return false

}
