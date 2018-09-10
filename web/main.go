package main

import (
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/",handler)
	http.ListenAndServe(":8123",nil)
}
// handler 是http响应的回调函数
func handler(Writer http.ResponseWriter, request *http.Request){
	// 输出hello world和当前url路径
	fmt.Fprintf(Writer, "Hello World,%s!",request.URL.Path[1:])
}