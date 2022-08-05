package config

import (
	"github.com/BurntSushi/toml"
	"os"
)

type tomlConfig struct {
	//Viewer和 SystemConfig 配置文件中涉及到的已经赋值的元素会自动赋值
	Viewer Viewer
	System SystemConfig
}

type Viewer struct {
	Title       string
	Description string
	Logo        string
	Navigation  []string
	Bilibili    string
	Avatar      string
	UserName    string
	UserDesc    string
}

type SystemConfig struct {
	AppName         string
	Version         float32
	CurrentDir      string
	CdnURl          string
	QiniuAccesskey  string
	QiniuSecretkey  string
	Valine          bool
	ValineAppid     string
	ValineAppkey    string
	ValineServerURl string
}

// Cfg 定义一个该类型的全局变量
var Cfg *tomlConfig

//该方法完成配置文件的启动
func init() {
	//程序启动就会执行init方法
	Cfg = new(tomlConfig)

	//配置文件中没有定义的元素值 需要手动赋值
	Cfg.System.AppName = "ms-go-blog"
	Cfg.System.Version = 1.0
	currentDir, _ := os.Getwd()
	Cfg.System.CurrentDir = currentDir

	//解码配置文件 并启动
	_, err := toml.DecodeFile("config/config.toml", &Cfg)
	if err != nil {
		panic(err)
	}
}
