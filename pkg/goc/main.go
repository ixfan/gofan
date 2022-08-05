package main

import (
	"flag"
	"fmt"
	"github.com/ixfan/gofan/pkg/goc/generate"
	"github.com/ixfan/gofan/pkg/goc/initialize"
	"github.com/ixfan/gofan/pkg/web"
)

func main() {
	var module, table, command string
	flag.StringVar(&command, "c", "", "请输入运行命令")
	flag.StringVar(&module, "m", "", "请输入模块")
	flag.StringVar(&table, "t", "", "请输入table表名")
	flag.Parse()
	switch command {
	case "init":
		initialize.Make()
	case "domain":
		(&web.Engine{}).InitConfig().InitDataBase()
		_ = generate.Domain(module, table)
	case "repository":
		_ = generate.Repository(table)
	case "service":
		(&web.Engine{}).InitConfig().InitDataBase()
		_ = generate.Service(module, table)
	case "api":
		(&web.Engine{}).InitConfig().InitDataBase()
		_ = generate.Api(module, table)
	case "help":
		fmt.Println("goc -c init 初始化项目")
		fmt.Println("goc -c domain -m module -t tableName 生成结构体")
		fmt.Println("goc -c repository -m module -t tableName 生成仓储")
		fmt.Println("goc -c service -m module -t tableName 生成service")
		fmt.Println("goc -c api -m module -t tableName 生成控制器和路由")
		fmt.Println("goc -c help 帮助")
		return
	case "":
		(&web.Engine{}).InitConfig().InitDataBase()
		_ = generate.Domain(module, table)
		_ = generate.Api(module, table)
		_ = generate.Service(module, table)
		_ = generate.Repository(table)
	}
	fmt.Println("命令执行完成")
}
