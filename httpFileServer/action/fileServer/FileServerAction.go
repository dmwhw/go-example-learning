package fileServerAction

import (
	"httpFileServer/assets/web"
	"net/http"

	assetfs "github.com/elazarl/go-bindata-assetfs"
)

//httpFileServer.ServerFile

type HttpFileServerHandler struct {
	HttpHandler http.Handler
}

func (httpFileServerHandler *HttpFileServerHandler) ServerFile(w http.ResponseWriter, req *http.Request) {
	httpFileServerHandler.HttpHandler.ServeHTTP(w, req)
}

func New() *HttpFileServerHandler {
	fs := assetfs.AssetFS{
		Asset:     web.Asset,
		AssetDir:  web.AssetDir,
		AssetInfo: web.AssetInfo,
		Fallback:  "index.html",
		Prefix:    "static"}
	httpFileServer := new(HttpFileServerHandler)
	httpFileServer.HttpHandler = http.FileServer(&fs)
	return httpFileServer
}
