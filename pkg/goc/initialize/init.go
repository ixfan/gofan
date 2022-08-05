package initialize

import (
	"github.com/gookit/goutil/fsutil"
	"github.com/ixfan/gofan/pkg/global"
	"github.com/ixfan/gofan/pkg/goc/template"
	"gopkg.in/yaml.v2"
	"os"
	"strings"
)

type Config struct {
	Mysql struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DataBase string `yaml:"database"`
	} `yaml:"mysql"`
	Redis struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Password string `yaml:"password"`
		Db       string `yaml:"db"`
	} `yaml:"redis"`
}

func Make() {
	//初始化目录
	_ = os.Mkdir("./config", os.ModePerm)
	_ = os.Mkdir("./internal", os.ModePerm)
	_ = os.Mkdir("./internal/api", os.ModePerm)
	_ = os.Mkdir("./internal/api/controller", os.ModePerm)
	_ = os.Mkdir("./internal/api/route", os.ModePerm)
	_ = os.Mkdir("./internal/application", os.ModePerm)
	_ = os.Mkdir("./internal/domain", os.ModePerm)
	_ = os.Mkdir("./internal/domain/entity", os.ModePerm)
	_ = os.Mkdir("./internal/domain/aggregate", os.ModePerm)
	_ = os.Mkdir("./internal/domain/interfaces", os.ModePerm)
	_ = os.Mkdir("./internal/domain/event", os.ModePerm)
	_ = os.Mkdir("./internal/domain/event/publish", os.ModePerm)
	_ = os.Mkdir("./internal/domain/event/subscribe", os.ModePerm)
	_ = os.Mkdir("./internal/infrastructure", os.ModePerm)
	_ = os.Mkdir("./internal/infrastructure/model", os.ModePerm)
	_ = os.Mkdir("./internal/infrastructure/dto", os.ModePerm)
	_ = os.Mkdir("./internal/infrastructure/repository", os.ModePerm)
	_ = os.Mkdir("./internal/infrastructure/dao", os.ModePerm)
	//初始化配置文件
	configFile := global.ConfigPath
	if !fsutil.FileExists(configFile) {
		cBytes, _ := yaml.Marshal(Config{})
		_ = os.WriteFile(configFile, cBytes, os.ModePerm)
	}
	//初始化main
	mainFile := "./main.go"
	if !fsutil.FileExists(mainFile) {
		_ = os.WriteFile(mainFile, []byte(strings.TrimSpace(template.Main)), os.ModePerm)
	}
}
