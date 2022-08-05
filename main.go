package main //这里必须是main  一个程序中必须有一个main  出现在main函数之处
import (
	"html/template"
	"log"
	"ms-go-blog/config"
	"ms-go-blog/models"
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

	//indexData数据类型用于距举例
	var indexData IndexData
	indexData.Title = "马神之路go博客"
	indexData.Desc = "现在是入门"
	//template包
	t := template.New("index.html")
	//1.拿到当前路径 （后定义Cfg 就可以将os.Getwd()改成已经记录的 CurrentDir）
	path := config.Cfg.System.CurrentDir
	//解析文件函数ParseFiles
	//访问博客首页模板的时候 因为有多个模板的嵌套 解析文件时 需要将其涉及到的所有模板都进行解析
	home := path + "/template/home.html"
	header := path + "/template/layout/header.html"
	footer := path + "/template/layout/footer.html"
	personal := path + "/template/layout/personal.html"
	pagination := path + "/template/layout/pagination.html"
	post := path + "/template/layout/post-list.html"
	//一个函数解析所有 此函数可以输入多个文件名
	t, _ = t.ParseFiles(path+"/template/index.html", home, header, footer, personal, pagination, post)

	//页面上涉及到的所有数据 必须有定义
	viewer := config.Viewer{
		Title: "马神之路go语言博客",
		//写了一个就发现这里数据逐个赋值很麻烦  要配置数据文件
	}
	var hr = &models.HomeResponse{}
	t.Execute(w, hr)
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
