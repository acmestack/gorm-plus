package gplus

import (
	"gorm.io/gorm/schema"
	"reflect"
	"sync"
)

// 缓存项目中所有实体字段名，储存格式：key为字段指针值，value为字段名
// 通过缓存实体的字段名，方便gplus通过字段指针获取到对应的字段名
var columnNameCache sync.Map

// 缓存实体对象，主要给NewQuery方法返回使用
var modelInstanceCache sync.Map

// Cache 缓存实体对象所有的字段名
func Cache(model any, namingStrategy ...schema.Namer) {
	valueOf := reflect.ValueOf(model).Elem()
	typeOf := reflect.TypeOf(model).Elem()

	for i := 0; i < valueOf.NumField(); i++ {
		field := typeOf.Field(i)
		// 如果当前实体嵌入了其他实体，同样需要缓存它的字段名
		if field.Anonymous {
			// 如果存在多重嵌套，通过递归方式获取他们的字段名
			subFieldMap := getSubFieldColumnNameMap(valueOf, field)
			for key, value := range subFieldMap {
				columnNameCache.Store(key, value)
			}
		} else {
			// 获取对象字段指针值
			pointer := valueOf.Field(i).Addr().Pointer()
			name := parseColumnName(field, namingStrategy...)
			columnNameCache.Store(pointer, name)
		}
	}

	// 缓存对象
	modelTypeStr := reflect.TypeOf(model).Elem().String()
	modelInstanceCache.Store(modelTypeStr, model)
}

// 递归获取嵌套字段名
func getSubFieldColumnNameMap(valueOf reflect.Value, field reflect.StructField) map[uintptr]string {
	result := make(map[uintptr]string)

	modelType := field.Type
	if modelType.Kind() == reflect.Ptr {
		modelType = modelType.Elem()
	}
	for j := 0; j < modelType.NumField(); j++ {
		subField := modelType.Field(j)
		if subField.Anonymous {
			nestedFields := getSubFieldColumnNameMap(valueOf, subField)
			for key, value := range nestedFields {
				result[key] = value
			}
		} else {
			pointer := valueOf.FieldByName(modelType.Field(j).Name).Addr().Pointer()
			name := parseColumnName(modelType.Field(j))
			result[pointer] = name
		}
	}

	return result
}

// 获取字段名称
func parseColumnName(field reflect.StructField, namingStrategy ...schema.Namer) string {
	tagSetting := schema.ParseTagSetting(field.Tag.Get("gorm"), ";")
	name, ok := tagSetting["COLUMN"]
	if ok {
		return name
	}

	if len(namingStrategy) > 0 {
		return namingStrategy[0].ColumnName("", field.Name)
	}

	strategy := schema.NamingStrategy{}
	return strategy.ColumnName("", field.Name)
}
