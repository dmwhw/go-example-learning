package main

//https://blog.csdn.net/idwtwt/article/details/81180460

import (
	"fmt"
	fileAction "httpFileServer/action/file"
	fileServerAction "httpFileServer/action/fileServer"

	"httpFileServer/filter"
	"httpFileServer/handler/errorHandler"
	"net/http"
	"os"
)

func main() {
	fmt.Println(os.Args)
	////////////////////////////////File Server start/////////////////////////////////////

	//http.Handle("/", http.FileServer(&fs)) //to FileServerfolder file without auth
	fileServerFilter := filter.NewFilterRouter()
	fileServerHandler := new(filter.FilterHandler)
	fileServerHandler.Action = fileServerAction.New().ServerFile
	fileServerFilter.AddHanlder("^/(.*)$", fileServerHandler)
	////////////////////////////////File Server end/////////////////////////////////////

	////////////////////////////////API start/////////////////////////////////////
	downloadFileHandler := new(filter.FilterHandler)
	downloadFileHandler.Action = fileAction.DownloadFile

	uploadFileHandler := new(filter.FilterHandler)
	uploadFileHandler.Action = fileAction.UploadFile

	listFileHandler := new(filter.FilterHandler)
	listFileHandler.Action = fileAction.ListFile

	httpAPIFilter := filter.NewFilterRouter()
	httpAPIFilter.GlobalErrorHandler = errorHandler.Handle

	httpAPIFilter.AddHanlder("/api/file/down", downloadFileHandler)
	httpAPIFilter.AddHanlder("/api/file/upload", uploadFileHandler)
	httpAPIFilter.AddHanlder("/api/file/list", listFileHandler)
	////////////////////////////////API end/////////////////////////////////////

	///AddHttp routing
	http.HandleFunc("/", fileServerFilter.Routing) // provide download validate

	http.HandleFunc("/api/", httpAPIFilter.Routing) // provide download validate

	http.ListenAndServe(":8080", nil)
}
