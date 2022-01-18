package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode string
	HttpPort string
	JwtKey string

	DB string
	DBHost string
	DBPost string
	DBUser string
	DBPassWord string
	DBName string

	AccessKey string
	SecretKey string
	Bucket string
	QiniuSever string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径：",err)
	}

	LoadServer(file)
	LoadData(file)
	LoadQiniu(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
	JwtKey = file.Section("server").Key("JwtKey").MustString("89js82js72dfasfsda")
}

func LoadData(file *ini.File) {
	DB = file.Section("database").Key("DB").MustString("mysql")
	DBHost = file.Section("database").Key("DBHost").MustString("localhost")
	DBPost = file.Section("database").Key("DBPost").MustString("3306")
	DBUser = file.Section("database").Key("DBUser").MustString("root")
	DBPassWord = file.Section("database").Key("DBPassWord").MustString("2468.shiSHI")
	DBName = file.Section("database").Key("DBName").MustString("gin_blog")
}

func LoadQiniu(file *ini.File) {
	AccessKey = file.Section("qiniu").Key("AccessKey").String()
	SecretKey = file.Section("qiniu").Key("SecretKey").String()
	Bucket = file.Section("qiniu").Key("Bucket").String()
	QiniuSever = file.Section("qiniu").Key("QiniuSever").String()
}