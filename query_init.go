package ormx

import (
	"io"
	"reflect"
	"strings"
	"text/template"
)

func QueryInit(v any, wr io.Writer) {
	typeV := reflect.TypeOf(v)
	if typeV.Kind() != reflect.Struct {
		panic("model type must struct")
	}

	name := typeV.Name()
	path := typeV.PkgPath()
	split := strings.Split(path, "/")
	typeName := typeV.Name()
	if len(split) > 1 {
		typeName = split[len(split)-1] + "." + typeName
	}

	data := map[string]string{}
	data["queryName"] = strings.ToLower(name[:1]) + name[1:]
	data["typeName"] = typeName

	t := template.New("")

	_, err := t.Parse(querTemplate)
	if err != nil {
		panic(err)
	}

	t.Execute(wr, data)

	num := typeV.NumField()

	fields := make([]reflect.StructField, 0)
	for i := 0; i < num; i++ {
		field := typeV.Field(i)
		fields = append(fields, field)
	}

	for i := 0; i < num; i++ {
		field := typeV.Field(i)
		fieldInfo, _ := fieldHandler(field, fields)
		if fieldInfo.FieldType == 1 {
			switch field.Type.Kind() {
			case reflect.String:
				genFunc(name, fieldInfo, stringTemplate, wr)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				genFunc(name, fieldInfo, numTemplate, wr)
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				genFunc(name, fieldInfo, numTemplate, wr)
			}
		}
		if fieldInfo.FieldType == 2 {
			genFunc(name, fieldInfo, stringTemplate, wr)
		}

	}

}

func fieldHandler(f reflect.StructField, fields []reflect.StructField) (FieldName, []string) {
	fieldName := f.Name
	typeName := ""
	pkgPaths := make([]string, 0)
	fieldType := 0
	if f.Type.Kind().String() == f.Type.Name() { // 基础类型
		typeName = f.Type.Name()
		fieldType = 1
	} else { // 自定义类型
		path := f.Type.PkgPath()
		split := strings.Split(path, "/")
		typeName = split[len(split)-1] + "." + f.Type.Name()
		pkgPaths = append(pkgPaths, path)
	}
	return FieldName{FieldType: fieldType, fieldName: fieldName, typeName: typeName}, pkgPaths
}

type FieldName struct {
	FieldType   int // 0 omit 1 基础类型 2 struct
	fieldName   string
	typeName    string
	preloadName string
}

func genFunc(queryName string, field FieldName, textTemplate string, wr io.Writer) {
	data := map[string]string{}

	data["repoName"] = "_" + strings.ToLower(queryName[:1]) + queryName[1:]
	data["fieldName"] = field.fieldName
	data["tableFieldName"] = snakeString(field.fieldName)
	data["paramName"] = strings.ToLower(field.fieldName[:1]) + field.fieldName[1:]
	if field.fieldName == "ID" {
		data["paramName"] = "id"
		data["tableFieldName"] = "id"
	}
	data["typeName"] = field.typeName
	t := template.New("")

	_, err := t.Parse(textTemplate)
	if err != nil {
		panic(err)
	}

	t.Execute(wr, data)
}

const querTemplate = `
package repo

type _{{.queryName}}Query[Query any] struct {
	ormx.IDBExec[Query]
	parent *Query
}

type {{.queryName}}Query struct {
	*_{{.queryName}}Query[{{.queryName}}Query]
	*ormx.DBXQuery[{{.typeName}}, {{.queryName}}Query]
}

type {{.queryName}}Update struct {
	*_{{.queryName}}Query[{{.queryName}}Update]
	*ormx.DBXUpdate[{{.typeName}}, {{.queryName}}Update]
}

type {{.queryName}}Delete struct {
	*_{{.queryName}}Query[{{.queryName}}Delete]
	*ormx.DBXDelete[{{.typeName}}, {{.queryName}}Delete]
}

func New{{.queryName}}Query(ctx context.Context) *{{.queryName}}Query {
	query := &{{.queryName}}Query{}
	query._{{.queryName}}Query = &_{{.queryName}}Query[{{.queryName}}Query]{}
	query.DBXQuery = ormx.NewDBXQuery[{{.typeName}}, {{.queryName}}Query](ctx, query)
	query.IDBExec = query.DBXQuery
	query.parent = query
	return query
}

func New{{.queryName}}Update(ctx context.Context) *{{.queryName}}Update {
	update := &{{.queryName}}Update{}
	update._{{.queryName}}Query = &_{{.queryName}}Query[{{.queryName}}Update]{}
	update.DBXUpdate = ormx.NewDBXUpdate[{{.typeName}}, {{.queryName}}Update](ctx, update)
	update.IDBExec = update.DBXUpdate
	update.parent = update
	return update
}

func New{{.queryName}}Delete(ctx context.Context) *{{.queryName}}Delete {
	del := &{{.queryName}}Delete{}
	del._{{.queryName}}Query = &_{{.queryName}}Query[{{.queryName}}Delete]{}
	del.DBXDelete = ormx.NewDBXDelete[{{.typeName}}, {{.queryName}}Delete](ctx, del)
	del.IDBExec = del.DBXDelete
	del.parent = del
	return del
}

func New{{.queryName}}Create(ctx context.Context) *ormx.DBXCreate[{{.typeName}}] {
	return ormx.NewDBXCreate[{{.typeName}}](ctx)
}


`

const stringTemplate = `

func (us *{{.repoName}}Query[Query]) {{.fieldName}}({{.paramName}} {{.typeName}}) *Query {
	us.IDBExec.Where("{{.tableFieldName}} = ?", {{.paramName}})
	return us.parent
}

func (us *{{.repoName}}Query[Query]) {{.fieldName}}GT({{.paramName}} {{.typeName}}) *Query {
	us.IDBExec.Where("{{.tableFieldName}} > ?", {{.paramName}})
	return us.parent
}

func (us *{{.repoName}}Query[Query]) {{.fieldName}}LT({{.paramName}} {{.typeName}}) *Query {
	us.IDBExec.Where("{{.tableFieldName}} < ?", {{.paramName}})
	return us.parent
}


func (us *{{.repoName}}Query[Query]) {{.fieldName}}LTE({{.paramName}} {{.typeName}}) *Query {
	us.IDBExec.Where("{{.tableFieldName}} <= ?", {{.paramName}})
	return us.parent
}

func (us *{{.repoName}}Query[Query]) {{.fieldName}}GTE({{.paramName}} {{.typeName}}) *Query {
	us.IDBExec.Where("{{.tableFieldName}} >= ?", {{.paramName}})
	return us.parent
}

func (us *{{.repoName}}Query[Query]) {{.fieldName}}Like({{.paramName}} {{.typeName}}) *Query {
	us.IDBExec.Where("{{.tableFieldName}} like ?", "%"+{{.paramName}}+"%")
	return us.parent
}

func (us *{{.repoName}}Query[Query]) {{.fieldName}}In({{.paramName}}s []{{.typeName}}) *Query {
	us.IDBExec.Where("{{.tableFieldName}} in ?", {{.paramName}}s)
	return us.parent
}

func (us *{{.repoName}}Query[Query]) {{.fieldName}}NEq({{.paramName}} {{.typeName}}) *Query {
	us.IDBExec.Where("{{.tableFieldName}} != ?", {{.paramName}})
	return us.parent
}

func (us *{{.repoName}}Query[Query]) {{.fieldName}}NotIn({{.paramName}}s []{{.typeName}}) *Query {
	us.IDBExec.Where("{{.tableFieldName}} not in ?", {{.paramName}}s)
	return us.parent
}

`

const numTemplate = `

func (us *{{.repoName}}Query[Query]) {{.fieldName}}({{.paramName}} {{.typeName}}) *Query {
	us.IDBExec.Where("{{.tableFieldName}} = ?", {{.paramName}})
	return us.parent
}

func (us *{{.repoName}}Query[Query]) {{.fieldName}}GT({{.paramName}} {{.typeName}}) *Query {
	us.IDBExec.Where("{{.tableFieldName}} > ?", {{.paramName}})
	return us.parent
}

func (us *{{.repoName}}Query[Query]) {{.fieldName}}LT({{.paramName}} {{.typeName}}) *Query {
	us.IDBExec.Where("{{.tableFieldName}} < ?", {{.paramName}})
	return us.parent
}


func (us *{{.repoName}}Query[Query]) {{.fieldName}}LTE({{.paramName}} {{.typeName}}) *Query {
	us.IDBExec.Where("{{.tableFieldName}} <= ?", {{.paramName}})
	return us.parent
}

func (us *{{.repoName}}Query[Query]) {{.fieldName}}GTE({{.paramName}} {{.typeName}}) *Query {
	us.IDBExec.Where("{{.tableFieldName}} >= ?", {{.paramName}})
	return us.parent
}

func (us *{{.repoName}}Query[Query]) {{.fieldName}}In({{.paramName}}s []{{.typeName}}) *Query {
	us.IDBExec.Where("{{.tableFieldName}} in ?", {{.paramName}}s)
	return us.parent
}

func (us *{{.repoName}}Query[Query]) {{.fieldName}}NEq({{.paramName}} {{.typeName}}) *Query {
	us.IDBExec.Where("{{.tableFieldName}} != ?", {{.paramName}})
	return us.parent
}

func (us *{{.repoName}}Query[Query]) {{.fieldName}}NotIn({{.paramName}}s []{{.typeName}}) *Query {
	us.IDBExec.Where("{{.tableFieldName}} not in ?", {{.paramName}}s)
	return us.parent
}

`

func snakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		// or通过ASCII码进行大小写的转化
		// 65-90（A-Z），97-122（a-z）
		//判断如果字母为大写的A-Z就在前面拼接一个_
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	//ToLower把大写字母统一转小写
	return strings.ToLower(string(data[:]))
}
