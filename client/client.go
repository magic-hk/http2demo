package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"github.com/custhk/http2demo/server/srvpolicy"
	"github.com/custhk/http2demo/client/pushhandler"
	"github.com/custhk/http2demo/resource"
	"golang.org/x/net/http2"
)

const (
	_CertFile = "../pem/cert.pem"
)

//NewPushHandlerSupportTransport set transport push support handler
func NewPushHandlerSupportTransport(pushHandler http2.PushHandler) *http2.Transport {
	// Create a pool with the server certificate since it is not signed
	// by a known CA
	caCert, err := ioutil.ReadFile(_CertFile)
	if err != nil {
		log.Fatalf("Reading server certificate: %s", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Create TLS configuration with the certificate of the server
	tlsConfig := &tls.Config{
		RootCAs: caCertPool,
	}

	tr := &http2.Transport{
		TLSClientConfig: tlsConfig,
	}
	// set push handler ,that's why client can support push handler
	tr.PushHandler = pushHandler
	return tr
}

// NewClient transport binding
func NewClient(tr *http2.Transport) *http.Client {
	client := http.Client{}
	client.Transport = tr
	return &client
}

// NewPushHandlerSupportClient new push handler support client
func NewPushHandlerSupportClient(pushHandler http2.PushHandler) *http.Client {
	return NewClient(NewPushHandlerSupportTransport(pushHandler))
}



// NoPushSrvClient to test NoPusher
func NoPushSrvClient() {
	pushHandler := pushhandler.NewDefaultPushHandler()
	client := NewPushHandlerSupportClient(pushHandler)
	var mpdURL string = "https://localhost:8080/image/0.jpg"
	mpdReq, err := http.NewRequest(http.MethodGet, mpdURL, nil)
	if err != nil {
		log.Printf("创建请求失败: %s", err)
		return
	}
	// 发送初始请求
	mpdResp, err := client.Do(mpdReq)
	if err != nil {
		log.Printf("发送请求失败: %s", err)
		return
	}
	resource.SaveResByURLPath(mpdReq.URL.Path, mpdResp)
	time.Sleep(20 * time.Second)
}

// PushSpecSrvClient to test SpecifiedPusher
func PushSpecSrvClient() {
	pushHandler := pushhandler.NewDefaultPushHandler()
	client := NewPushHandlerSupportClient(pushHandler)
	var mpdURL string = "https://localhost:8080/image/0.jpg"
	mpdReq, err := http.NewRequest(http.MethodGet, mpdURL, nil)
	mpdReq.Header.Set(srvpolicy.HEADERKEY, srvpolicy.PushSpePolicy)
	mpdReq.Header.Set(srvpolicy.GetSpecPromiseHeaderKey(), "1.jpeg,2.jpeg,3.jpg")
	if err != nil {
		log.Printf("创建请求失败: %s", err)
		return
	}
	// 发送初始请求
	mpdResp, err := client.Do(mpdReq)
	if err != nil {
		log.Printf("发送请resource求失败: %s", err)
		return
	}
	resource.SaveResByURLPath(mpdReq.URL.Path, mpdResp)
	time.Sleep(20 * time.Second)
}
func main() {
	//NoPushSrvClient()
	PushSpecSrvClient()
}