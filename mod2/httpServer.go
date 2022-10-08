package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
)

//var l *zap.Logger

const httpPort = 9999
const version = "1.1"

func main() {

	http.HandleFunc("/", httpAccessFunc)

	http.HandleFunc("/healthz", healthFunc)

	err := http.ListenAndServe(":"+strconv.Itoa(httpPort), nil)

	if err != nil {
		log.Fatal(err.Error())
	}
}

func healthFunc(w http.ResponseWriter, r *http.Request) {
	HealthCode := "200"
	w.Write([]byte(HealthCode))
}

func httpAccessFunc(w http.ResponseWriter, r *http.Request) {

	if len(r.Header) > 0 {
		for k, v := range r.Header {
			log.Printf("%s=%s", k, v[0])
			w.Header().Set(k, v[0])
		}
	}

	r.ParseForm() //解析所有请求数据
	if len(r.Form) > 0 {
		for k, v := range r.Form {
			log.Printf("%s=%s", k, v[0])
		}
	}

	os.Setenv("VERSION", version) //设置环境值的值

	name := os.Getenv("VERSION")
	log.Printf("VERSION Env: ", name)

	//获取IP，打印出来
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		fmt.Println("err:", err)
	}

	if net.ParseIP(ip) != nil {
		fmt.Println("ip=", ip)
		log.Println(ip)
	}

	fmt.Println("code=", http.StatusOK)
	log.Println(http.StatusOK)

	//response响应
	w.WriteHeader(http.StatusOK)

	w.Write([]byte("Server Access,Success!"))
}
