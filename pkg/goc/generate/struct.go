package generate

import (
	"github.com/dave/jennifer/jen"
	"github.com/gookit/goutil/arrutil"
	"github.com/gookit/goutil/strutil"
)

type GenStruct struct {
	packageName string
	name        string
	comment     string
	ignores     []string
	labelTag    bool
}

func NewGenStruct(packageName string) *GenStruct {
	return &GenStruct{
		packageName: packageName,
	}
}

func (gen *GenStruct) EnableLabelTag() *GenStruct {
	gen.labelTag = true
	return gen
}

func (gen *GenStruct) SetName(name string) *GenStruct {
	gen.name = name
	return gen
}

func (gen *GenStruct) SetComment(comment string) *GenStruct {
	gen.comment = comment
	return gen
}

func (gen *GenStruct) SetIgnores(data []string) *GenStruct {
	gen.ignores = data
	return gen
}

func (gen *GenStruct) SetTable(table Table) *GenStruct {
	gen.name = strutil.CamelCase(table.Name)
	gen.comment = table.Comment
	return gen
}

func (gen *GenStruct) GenerateColumn(columns []Column) string {
	code := jen.NewFile(gen.packageName)
	code.Comment(strutil.UpperFirst(gen.name) + " " + gen.comment)
	code.Type().Id(strutil.UpperFirst(gen.name)).StructFunc(func(group *jen.Group) {
		for _, column := range columns {
			if arrutil.InStrings(column.Field, gen.ignores) {
				continue
			}
			gen.addColumnFiled(group, column)
		}
	})
	return code.GoString()
}

func (gen *GenStruct) GenerateQueryOne(columns []Column) string {
	code := jen.NewFile(gen.packageName)
	code.Comment(strutil.UpperFirst(gen.name) + " " + gen.comment)
	code.Type().Id(strutil.UpperFirst(gen.name)).StructFunc(func(group *jen.Group) {
		for _, column := range columns {
			if column.Key != "PRI" {
				continue
			}
			gen.addColumnFiled(group, column)
		}
	})
	return code.GoString()
}

func (gen *GenStruct) GenerateUpdate(columns []Column) string {
	code := jen.Comment("// Update 更新数据").Op("\n")
	code.Func().Call(jen.Id(gen.name).Add(jen.Op("*")).Id(strutil.UpperFirst(gen.name))).Id("Update").Call(jen.Id("options").Map(jen.String()).Interface()).BlockFunc(func(group *jen.Group) {
		for _, col := range columns {
			if arrutil.InStrings(col.Field, gen.ignores) {
				continue
			}
			fieldType := GetFiledType(col)
			varName := strutil.CamelCase(col.Field)
			if arrutil.InStrings(varName, []string{"type", "func", "interface", "struct", "if", "else", "for", "range"}) {
				varName = "_" + varName
			}
			group.Comment(col.Comment)
			group.If(jen.Id(varName)).Op(",").Op("ok").Op(":=").Id("options").Types(jen.Lit(strutil.UpperFirst(strutil.CamelCase(col.Field)))).Op(";").Op("ok").BlockFunc(func(group *jen.Group) {
				group.Id(gen.name + "." + col.UpperCamelCaseField()).Op("=").Id(varName).Assert(jen.Op(fieldType))
			})
		}
	})
	return "\n" + code.GoString()
}

func (gen *GenStruct) GenerateList(columns []Column) string {
	code := jen.NewFile(gen.packageName)
	jen.Comment(strutil.UpperFirst(gen.name) + " " + gen.comment)
	code.Type().Id(strutil.UpperFirst(gen.name)).StructFunc(func(group *jen.Group) {
		group.Id("Page").Int().Tag(map[string]string{
			"json": "page",
		}).Comment("页码")
		group.Id("PageSize").Int().Tag(map[string]string{
			"json": "pageSize",
		}).Comment("每页行数")
	})
	return code.GoString()
}

func (gen *GenStruct) addColumnFiled(group *jen.Group, column Column) {
	fieldType := GetFiledType(column)
	jsonTag := strutil.Camel(column.Field)
	if fieldType == "int64" {
		jsonTag = jsonTag + ",string"
	}
	tags := map[string]string{
		"json": jsonTag,
	}
	if gen.labelTag {
		tags["label"] = column.Comment
	}
	group.Comment(column.Comment)
	if fieldType == "*time.Time" {
		group.Id(column.UpperCamelCaseField()).Op("*").Qual("time", "Time").Tag(tags)
	} else {
		group.Id(column.UpperCamelCaseField()).Op(fieldType).Tag(tags)
	}
}
