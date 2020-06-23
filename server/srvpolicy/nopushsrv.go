package srvpolicy

import (
	"io/ioutil"
	"net/http"
	"github.com/custhk/http2demo/resource"
	
)
// NoPushSrv default service policy
type NoPushSrv struct {
}

//Handle 处理push 请求
func (srv *NoPushSrv) handle(w http.ResponseWriter, r *http.Request) {
	//解析URL，获取请求资源
	fileInfo := resource.ParseURL(r.URL.Path)
	if fileInfo != nil {
		file, err := ioutil.ReadFile(fileInfo.GetLocalPath())
		if err != nil {
			panic(err)
		}
		w.Header().Add("Content-Type", "application/octet-stream")
		w.Header().Add("Access-Control-Allow-Origin", "*")	
		w.Write(file)
	}
}