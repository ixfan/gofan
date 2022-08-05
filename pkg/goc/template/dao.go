package template

var Dao = `
package dao

import "github.com/ixfan/gofan/pkg/web"

type {{className}}Dao struct {
	context *web.Context
}

func New{{className}}Dao(context *web.Context) *{{className}}Dao {
	return &{{className}}Dao{context: context}
}

`
