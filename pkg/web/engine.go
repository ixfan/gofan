package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ixfan/gofan/pkg/database/orm"
	"github.com/ixfan/gofan/pkg/global"
	"github.com/ixfan/gofan/pkg/web/filters"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Engine struct {
	*gin.Engine
}

var server *Engine

// Default 默认实例
func Default() *Engine {
	if server == nil {
		engine := gin.Default()
		engine.Use(filters.Cors(), filters.Transaction())
		server = &Engine{
			engine,
		}
	}
	return server
}

// InitConfig 初始化配置
func (engine *Engine) InitConfig(configFiles ...string) *Engine {
	configFile := global.ConfigPath
	if len(configFiles) > 0 {
		configFile = configFiles[0]
	}
	config := make(map[string]map[string]string)
	mBytes, _ := ioutil.ReadFile(configFile)
	_ = yaml.Unmarshal(mBytes, &config)
	if len(config) > 0 {
		for index, item := range config {
			if len(item) > 0 {
				for key, value := range item {
					_ = os.Setenv(index+"."+key, value)
				}
			}
		}
	}
	return engine
}

// InitDataBase 初始化数据库
func (engine *Engine) InitDataBase() *Engine {
	orm.InitGorm()
	return engine
}

// InitJob 初始化任务
func (engine *Engine) InitJob(handler ...func()) *Engine {
	if len(handler) > 0 {
		for _, item := range handler {
			go item()
		}
	}
	return engine
}

// InitRoute 初始化路由
func (engine *Engine) InitRoute(handler ...func()) *Engine {
	if len(handler) > 0 {
		for _, item := range handler {
			item()
		}
	}
	return engine
}

// Middleware 使用全局中间件
func (engine *Engine) Middleware(handlers ...gin.HandlerFunc) *Engine {
	if len(handlers) > 0 {
		for _, item := range handlers {
			engine.Engine.Use(item)
		}
	}
	return engine
}

// Start 启动服务
func (engine *Engine) Start() {
	port := "8080"
	if os.Getenv("server.port") != "" {
		port = os.Getenv("server.port")
	}
	Default().Static("/static", "./static")
	fmt.Println("http start at :" + port)
	_ = Default().Run(":" + port)
}
