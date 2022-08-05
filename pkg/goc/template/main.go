package template

var Main = `
package main

import (
	"github.com/ixfan/gofan/pkg/web"
)

func main(){
	web.Default().
		InitConfig().
		InitDataBase().
		InitRoute().
		InitJob().
		Start()
}
`
