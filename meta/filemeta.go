package meta

//FileMeta 文件結構體
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

// 新增/更新 FileMeta結構體
func UpdateFileMeta(file FileMeta) {
	fileMetas[file.FileSha1] = file
}

// 通過 Sha1值 獲取對應 FileMeta 結構體
func GetFileMeta(fileSha1 string) FileMeta {
	return fileMetas[fileSha1]
}