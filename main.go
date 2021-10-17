package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"reflect"
)


func handleIterceptor(h http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {


		//接收客户端 request，并将 request 中带的 header 写入 response header
		if len(r.Header) > 0 {
			for k,v := range r.Header {
				w.Header().Add(k,v[0])
			}
		}

		//读取当前系统的环境变量中的 VERSION 配置，并写入 response header
		var version string
		version = os.Getenv("VERSION")
		fmt.Println(version)
		w.Header().Add("VERSION",version)
		h(w, r)

		//客户端 IP，HTTP 返回码
		ip,_,_ :=net.SplitHostPort(r.RemoteAddr)
		fmt.Println("client IP:",ip)


		v:=reflect.ValueOf(w)
		codestatus:=v.Elem().FieldByName("status").Int()
		fmt.Println("reponse code:",codestatus)
	}
}


func IndexHandlers(w http.ResponseWriter, r *http.Request){



}


func HealthzHandlers(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(200)
	w.Write([]byte("200"))
}

func main () {
	http.HandleFunc("/", handleIterceptor(IndexHandlers))
	http.HandleFunc("/healthz", handleIterceptor(HealthzHandlers))
	err := http.ListenAndServe("127.0.0.1:80", nil)
	if err != nil {
		fmt.Printf("listen error:[%v]", err.Error())
	}
}
