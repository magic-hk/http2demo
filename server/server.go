package main

import (
	"github.com/custhk/http2demo/server/srvpolicy"
	"log"
	"net/http"
	"time"

)
const (
	_CertFile string = "../pem/cert.pem"
	_KeyFile string ="../pem/key.pem"
)
func main(){
	srv := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  9000 * time.Second,
		WriteTimeout: 9000 * time.Second,
	}
	http.HandleFunc("/", srvpolicy.Dispatch)
	log.Fatal(srv.ListenAndServeTLS(_CertFile,_KeyFile))
}