package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func healthz (w http.ResponseWriter, req *http.Request) {

	// 1、接收客户端 request，并将 request 中带的 header 写入 response header
	for name, headers := range req.Header {
		for _, h := range headers {
			w.Header().Set(name, h)
		}
	}

	// 2、读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	version := os.Getenv("VERSION")
	w.Header().Add("VERSION", version)

	// 3、Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
	log.Printf("%v %s", req.RemoteAddr, http.StatusOK)

	// 4、当访问 localhost/healthz 时，应返回 200
	io.WriteString(w, "200")
}

func main () {
	http.HandleFunc("/healthz", healthz)

	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Fatal("Server fail to start!")
		return
	}
}