package main //这里必须是main  一个程序中必须有一个main  出现在main函数之处
import (
	"encoding/json"
	"log"
	"net/http"
)

type IndexData struct {
	//本身大写是为了说明公共可以访问
	//加上json标识可以改变显示格式
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var indexData IndexData
	indexData.Title = "马神之路go博客"
	indexData.Desc = "现在是入门"
	jsonStr, _ := json.Marshal(indexData)
	//将上述写好的indexData转成的json类型变量写出来
	//json类型 与[]byte 是一个类型吗？
	w.Write(jsonStr)
}

func main() {
	//程序入口 一个项目一个入口
	//web程序  http协议  ip与port
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", index)

	//这个地方赋值过程写在了if条件处
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
