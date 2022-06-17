package schema

import (
	"go/ast"
	"go_learning/7days-golang/GeeORM/dialect"
	"reflect"
)

// -------------------------------------------
// @file          : schema.go
// @author        : binshow
// @time          : 2022/6/16 10:29 AM
// @description   :
// -------------------------------------------

// Schema 完成对象 和 表的转换，给定一个任意的对象，转换成关系型数据库中的表结构
// 数据库中创建一张表需要以下的要素：

// 表名(table name) —— 结构体名(struct name)
// 字段名和字段类型 —— 成员变量和类型。
// 额外的约束条件(例如非空、主键等) —— 成员变量的Tag（Go 语言通过 Tag 实现，Java、Python 等语言通过注解实现）

// Field represents a column of database
type Field struct {
	Name 	string		// 字段名
	Type    string		// 类型
	Tag     string		// 约束条件Tag
}

// Schema represents a table of database
type Schema struct {
	Model 	interface{}				// 被映射的对象 Model
	Name    string					// 表名 Name
	Fields  []*Field				// 字段 Fields
	FieldNames []string				// 包含所有的字段名
	fieldMap   map[string]*Field	// 记录字段名和 Field的 映射关系
}

func (schema *Schema) GetField(name string) *Field {
	return schema.fieldMap[name]
}

// Parse 将任意对象解析为 Schema 实例
func Parse(dest interface{}, d dialect.Dialect) *Schema {
	//reflect.TypeOf()	返回入参的类型
	//reflect.ValueOf() 返回入参的值
	// 因为这里设计的入参是一个 对象的指针，所以需要通过 Indirect 获取指针指向的实例
	modelType := reflect.Indirect(reflect.ValueOf(dest)).Type()
	schema := &Schema{
		Model:      dest,
		Name:       modelType.Name(),
		fieldMap:   make(map[string]*Field),
	}

	for i := 0; i < modelType.NumField(); i++ {
		p := modelType.Field(i)
		if !p.Anonymous && ast.IsExported(p.Name) {
			field := &Field{
				Name: p.Name,	// 字段名
				Type: d.DataTypeOf(reflect.Indirect(reflect.New(p.Type))),	// 字段类型
			}

			if v, ok := p.Tag.Lookup("geeorm"); ok {
				field.Tag = v
			}

			schema.Fields = append(schema.Fields , field)
			schema.FieldNames = append(schema.FieldNames, p.Name)
			schema.fieldMap[p.Name] = field

		}
	}
	return schema
}


func (schema *Schema) RecordValues(dest interface{}) []interface{} {
	destValue := reflect.Indirect(reflect.ValueOf(dest))
	var fieldValues []interface{}
	for _, field := range schema.Fields {
		fieldValues = append(fieldValues, destValue.FieldByName(field.Name).Interface())
	}
	return fieldValues
}