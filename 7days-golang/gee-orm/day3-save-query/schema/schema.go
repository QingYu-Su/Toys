package schema

import (
	"geeorm/dialect" // 引入方言包，用于获取字段的数据库数据类型
	"go/ast"         // 提供 Go 语法树操作功能，用于检查字段是否导出
	"reflect"        // 提供反射功能，用于操作类型和值
)

// Field 表示结构体中的一个字段，包含字段的名称、类型和标签。
type Field struct {
	Name string // 字段名称
	Type string // 字段的数据库数据类型
	Tag  string // 字段的自定义标签（如 geeorm 标签）
}

// Schema 表示一个结构体的模式，包含结构体的元信息和字段信息。
type Schema struct {
	Model      interface{}       // 模型的原始结构体
	Name       string            // 模型的名称（结构体名称）
	Fields     []*Field          // 模型的所有字段
	FieldNames []string          // 模型的所有字段名称
	fieldMap   map[string]*Field // 字段名称到字段对象的映射
}

// GetField 根据字段名称获取字段信息。
func (schema *Schema) GetField(name string) *Field {
	return schema.fieldMap[name] // 从字段映射中查找并返回字段
}

// Parse 解析一个结构体，生成对应的 Schema 对象。
func Parse(dest interface{}, d dialect.Dialect) *Schema {
	// 获取结构体的反射类型（去除指针）
	modelType := reflect.Indirect(reflect.ValueOf(dest)).Type()

	// 初始化 Schema 对象
	schema := &Schema{
		Model:    dest,                    // 保存原始结构体
		Name:     modelType.Name(),        // 设置结构体名称
		fieldMap: make(map[string]*Field), // 初始化字段映射
	}

	// 遍历结构体的所有字段
	for i := 0; i < modelType.NumField(); i++ {
		p := modelType.Field(i) // 获取字段信息
		// 跳过匿名字段和未导出字段
		if !p.Anonymous && ast.IsExported(p.Name) {
			// 创建 Field 对象
			field := &Field{
				Name: p.Name, // 字段名称
				// 使用方言获取字段的数据库数据类型
				Type: d.DataTypeOf(reflect.Indirect(reflect.New(p.Type))),
			}
			// 检查是否有 geeorm 标签
			if v, ok := p.Tag.Lookup("geeorm"); ok {
				field.Tag = v // 设置字段标签
			}
			// 将字段信息添加到 Schema 中
			schema.Fields = append(schema.Fields, field)
			schema.FieldNames = append(schema.FieldNames, p.Name)
			schema.fieldMap[p.Name] = field
		}
	}
	return schema // 返回解析后的 Schema 对象
}

// RecordValues 从目标结构体实例中提取字段值，返回一个值的切片。
// 参数 dest 是目标结构体实例，方法会根据 Schema 中定义的字段顺序提取对应的值。
func (schema *Schema) RecordValues(dest interface{}) []interface{} {
	// 获取目标结构体实例的反射值，并通过 reflect.Indirect 获取其实际值（去除指针）
	destValue := reflect.Indirect(reflect.ValueOf(dest))

	// 初始化一个切片，用于存储提取的字段值
	var fieldValues []interface{}

	// 遍历 Schema 中定义的字段
	for _, field := range schema.Fields {
		// 使用反射从目标结构体实例中提取字段值
		fieldValues = append(fieldValues, destValue.FieldByName(field.Name).Interface())
	}
	// 返回提取的字段值切片
	return fieldValues
}
