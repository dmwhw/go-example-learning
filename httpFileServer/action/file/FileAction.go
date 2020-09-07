package fileAction

import (
	"fmt"
	"net/http"
)

func UploadFile(w http.ResponseWriter, req *http.Request) {
	fmt.Println("upload...")
}

func DownloadFile(w http.ResponseWriter, req *http.Request) {
	fmt.Println("down...")
}

func ListFile(w http.ResponseWriter, req *http.Request) {
	fmt.Println("list...")
}
