package handler

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
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

         newFile, err := os.Create("./temp/" + head.Filename)
	     if err != nil {
	     	fmt.Printf("Failed to create file, err:%s\n", err.Error())
	     	return
		 }
         defer newFile.Close()

         _, err = io.Copy(newFile, file)
         if err != nil {
         	fmt.Printf("Failed to save file, err:%s\n", err.Error())
         	return
		 }

         http.Redirect(w, r, "/file/upload/success", http.StatusFound)
	}
}

func UploadSuccessHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Upload Finished!")
}