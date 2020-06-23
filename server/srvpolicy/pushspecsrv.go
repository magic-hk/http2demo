package srvpolicy

import (
	"io/ioutil"
	"net/http"
	"log"
	"strings"
	"github.com/custhk/http2demo/resource"

)

const (
	_SpecifiedPushList = "promiseList"
)

// GetSpecPromiseHeaderKey for push
func GetSpecPromiseHeaderKey() string {
	return _SpecifiedPushList
}
// PushSpecSrv server provide push
type PushSpecSrv struct {
}

// parserRequest 解析请求
func (srv *PushSpecSrv) parserRequest(r *http.Request) []resource.IResource {
	//在header中利用SpecPromiseHeaderKey来获取push文件列表
	promiseListStr := r.Header.Get(_SpecifiedPushList)
	result := []resource.IResource{}
	//解析URL获取基础请求资源信息
	fileInfo := resource.ParseURL(r.URL.Path)
	if fileInfo != nil {
		result = append(result, fileInfo)
	}
	if promiseListStr == "" {
		return result
	}
	//从promiseList中解析SVC文件列表，并加入到请求资源集中
	result = append(result, resource.ParsePromiseList(promiseListStr)...)
	r.Header.Del(_SpecifiedPushList)
	return result
}

//Handle 处理push 请求
func (srv *PushSpecSrv) handle(w http.ResponseWriter, r *http.Request) {
	//解析请求
	fileInfos := srv.parserRequest(r)
	//读取基础请求资源
	file, err := ioutil.ReadFile(fileInfos[0].GetLocalPath())
	if err != nil {
		panic(err)
	}
	length := len(fileInfos)
	//如果有promiseList,尝试push
	if length > 0 {
		pusher, ok := w.(http.Pusher)
		if ok {
			//遍历push
			for index := 1; index < length; index++ {
				log.Printf("Try to push: %v", fileInfos[index].GetURLPath())
				if err := pusher.Push(fileInfos[index].GetURLPath(), nil); err != nil {
					log.Printf("Failed to push: %v", err)
					if strings.Contains(err.Error(), "recursive") {
						break
					}
				}
			}
		}
	}
	w.Header().Add("Content-Type", "application/octet-stream")
	w.Write(file)
}