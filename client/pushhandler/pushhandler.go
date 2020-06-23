package pushhandler

import (
	"log"
	"time"

	"github.com/custhk/http2demo/resource"
	"golang.org/x/net/http2"
)

// DefaultPushHandler default push handler
type DefaultPushHandler struct {
}

// NewDefaultPushHandler reutrn new push handler
func NewDefaultPushHandler() *DefaultPushHandler {
	return &DefaultPushHandler{}
}

// HandlePush func to deal with server push
// push处理方法
func (ph *DefaultPushHandler) HandlePush(r *http2.PushedRequest) {
	handleWrite := make(chan struct{})
	// promise request
	promise := r.Promise
	// 开启线程，处理push回来的文件
	go func() {
		defer close(handleWrite)
		if promise == nil {
			log.Printf("promise not received")
			return
		}
		
		// parse to get fileInfo
		// 解析获取push文件信息
		fileInfo := resource.ParseURL(promise.URL.Path)
		if fileInfo != nil {
			//等待响应
			push, pushErr := r.ReadResponse(r.Promise.Context())
			//receiveTimeStamp := time.Now().UnixNano() / 1e6

			if pushErr != nil {
				log.Printf("push error = %v; want %v", pushErr, nil)
			}
			if push == nil {
				log.Printf("push not received")
			} else {
				//从响应中读取push回来的文件，并更新文件信息属性
				fileInfo = resource.SaveRes(fileInfo, push)
				if fileInfo != nil {
					log.Printf("save  push file = %q\n", promise.URL.Path)
					log.Printf("push file size= %v\n", fileInfo.GetDataSize())
				}

			}
		}

	}()
	//超时取消，如果在5秒钟之内，任务没有完成，则取消当前push 文件处理任务
	select {
	case <-handleWrite:
	case <-time.After(5 * time.Second):
		//case <-time.After(1 * time.Nanosecond):
		r.Cancel()
		log.Printf("-------cancel push file = %q\n-----------", promise.URL.Path)
	}

}