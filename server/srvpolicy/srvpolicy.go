package srvpolicy

import "net/http"

const (
	// NoPushPolicy default service policy
	NoPushPolicy string = ""
	// PushSpePolicy provide push service refer to req header 
	PushSpePolicy string = "specified"
)
// ISrvPolicy  server policy interface
type ISrvPolicy interface {
	handle(w http.ResponseWriter, r *http.Request)
}
// HEADERKEY policy name in header
const HEADERKEY = "srvPolicyName"

var handleMap map[string]ISrvPolicy

//初始化方法，构建一个处理map，绑定处理策略
func init() {
	handleMap = make(map[string]ISrvPolicy)
	handleMap[NoPushPolicy] = &NoPushSrv{}
	handleMap[PushSpePolicy] = &PushSpecSrv{}
}
func init() {

}
// Dispatch service policy
func Dispatch(w http.ResponseWriter, r *http.Request) {
	policyName := r.Header.Get(HEADERKEY)
	//在header中利用HEADERKEY获取客户端希望服务器采用的策略，默认情况下采用nopush策略
	nowPolicy, ok := handleMap[policyName]
	if ok {
		nowPolicy.handle(w, r)
	}
}