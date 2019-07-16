package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"storage/meta"
	"storage/util"
	"time"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//返回上傳的html頁面
		view, err := ioutil.ReadFile("./static/view/index.html")
		if err != nil {
            io.WriteString(w, "internal server error")
            return
		}
		io.WriteString(w, string(view))
	} else if r.Method == "POST" {
         //接收文件流及存儲到本地目錄
         file, head, err := r.FormFile("file")
         if err != nil {
         	 fmt.Printf("Failed to get file, err:%s\n", err.Error())
		     return
         }
         defer file.Close()

         fileMeta := meta.FileMeta{
         	FileName: head.Filename,
         	Location: "./temp/" + head.Filename,
         	UploadAt: time.Now().Format("2006-01-02 15:04:05"),
		 }

         newFile, err := os.Create(fileMeta.Location)
	     if err != nil {
	     	fmt.Printf("Failed to create file, err:%s\n", err.Error())
	     	return
		 }
         defer newFile.Close()

         fileMeta.FileSize, err = io.Copy(newFile, file)
         if err != nil {
         	fmt.Printf("Failed to save file, err:%s\n", err.Error())
         	return
		 }

         // 檔案內容回到起始位置
         newFile.Seek(0,0)
         fileMeta.FileSha1 = util.FileSha1(newFile)

         meta.UpdateFileMeta(fileMeta)

         fmt.Print(fileMeta)

         http.Redirect(w, r, "/file/upload/success", http.StatusFound)
	}
}

func UploadSuccessHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Upload Finished!")
}

func GetFileMetaHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	filehash := r.Form["filehash"][0]
	file := meta.GetFileMeta(filehash)
	data, err := json.Marshal(file)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(data)
}