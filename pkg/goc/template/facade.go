package template

var Facade = `
package facade

import (
	"github.com/ixfan/gofan/pkg/web"
	"{{goModule}}/internal/domain/interfaces"
	"{{goModule}}/internal/infrastructure/dao"
	"{{goModule}}/internal/infrastructure/repository"
)

func {{className}}Repository(context *web.Context) interfaces.{{className}}Repository {
	return repository.New{{className}}Repository(context)
}

func {{className}}Dao(context *web.Context) interfaces.{{className}}Dao {
	return dao.New{{className}}Dao(context)
}
`
