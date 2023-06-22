package gplus

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

type Condition struct {
	Group       string
	ColumnName  string
	Op          string
	ColumnValue any
}

var operators = []string{"!~<=", "!~>=", "~<=", "~>=", "!?=", "!^=", "!~=", "?=", "^=", "~=", "!=", ">=", "<=", "=", ">", "<"}

func BuildQuery[T any](queryParams url.Values) *QueryCond[T] {

	conditionMap, gcond := buildConditionMap(queryParams)

	queryCondMap := buildQueryCondMap[T](conditionMap)

	// 如果没有分组条件，直接返回默认的查询条件
	if len(gcond) == 0 {
		return queryCondMap["default"]
	}

	return buildGroupQuery[T](gcond, queryCondMap)
}

func buildConditionMap(queryParams url.Values) (map[string][]*Condition, string) {
	var maps = make(map[string][]*Condition)
	var gcond string
	for key, values := range queryParams {
		if key == "q" {
			for _, value := range values {
				var currentOperator string
				for _, op := range operators {
					if strings.Contains(value, op) {
						currentOperator = op
						break
					}
				}
				params := strings.SplitN(value, currentOperator, 2)
				if len(params) == 2 {
					condition := &Condition{}
					groups := strings.Split(params[0], ".")
					var groupName string
					var columnName string
					// 如果不包含组，默认分为同一个组
					if len(groups) == 1 {
						groupName = "default"
						columnName = groups[0]
					} else if len(groups) == 2 {
						groupName = groups[0]
						columnName = groups[1]
					}
					condition.Group = groupName
					condition.ColumnName = columnName
					condition.Op = currentOperator
					condition.ColumnValue = params[1]
					conditions, ok := maps[groupName]
					if ok {
						conditions = append(conditions, condition)
					} else {
						conditions = []*Condition{condition}
					}
					maps[groupName] = conditions
				}
			}
		} else if key == "gcond" {
			gcond = values[0]
		}
	}
	return maps, gcond
}

func buildQueryCondMap[T any](maps map[string][]*Condition) map[string]*QueryCond[T] {
	var queryMaps = make(map[string]*QueryCond[T])
	columnTypeMap := getColumnTypeMap[T]()
	for key, conditions := range maps {
		query, _ := NewQuery[T]()
		for _, condition := range conditions {
			name := condition.ColumnName
			op := condition.Op
			value := condition.ColumnValue
			switch op {
			case "=":
				if strings.ToLower(fmt.Sprintf("%s", value)) == "null" {
					query.IsNull(name)
				} else {
					query.Eq(name, convert(columnTypeMap, name, value))
				}
			case "!=":
				if strings.ToLower(fmt.Sprintf("%s", value)) == "null" {
					query.IsNotNull(name)
				} else {
					query.Ne(name, convert(columnTypeMap, name, value))
				}
			case ">":
				query.Gt(name, convert(columnTypeMap, name, value))
			case ">=":
				query.Ge(name, convert(columnTypeMap, name, value))
			case "<":
				query.Lt(name, convert(columnTypeMap, name, value))
			case "<=":
				query.Le(name, convert(columnTypeMap, name, value))
			case "?=":
				values := strings.Split(fmt.Sprintf("%s", value), ",")
				var queryValues []any
				for _, v := range values {
					queryValues = append(queryValues, convert(columnTypeMap, name, v))
				}
				query.In(name, queryValues)
			case "!?=":
				query.NotIn(name, convert(columnTypeMap, name, value))
			case "^=":
				values := strings.Split(fmt.Sprintf("%s", value), ",")
				if len(values) == 2 {
					query.Between(name, convert(columnTypeMap, name, values[0]), convert(columnTypeMap, name, values[1]))
				}
			case "!^=":
				values := strings.Split(fmt.Sprintf("%s", value), ",")
				if len(values) == 2 {
					query.NotBetween(name, convert(columnTypeMap, name, values[0]), convert(columnTypeMap, name, values[1]))
				}
			case "~=":
				query.Like(name, convert(columnTypeMap, name, value))
			case "!~=":
				query.NotLike(name, convert(columnTypeMap, name, value))
			case "~<=":
				query.LikeLeft(name, convert(columnTypeMap, name, value))
			case "~>=":
				query.LikeRight(name, convert(columnTypeMap, name, value))
			case "!~<=":
				query.NotLikeLeft(name, convert(columnTypeMap, name, value))
			case "!~>=":
				query.NotLikeRight(name, convert(columnTypeMap, name, value))
			}
		}
		queryMaps[key] = query
	}
	return queryMaps
}

func convert(columnTypeMap map[string]reflect.Type, name string, value any) any {
	columnType, ok := columnTypeMap[name]
	if ok {
		switch columnType.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			atoi, err := strconv.Atoi(fmt.Sprintf("%s", value))
			if err == nil {
				value = atoi
			}
		}
	}
	return value
}

func buildGroupQuery[T any](gcond string, queryMaps map[string]*QueryCond[T]) *QueryCond[T] {
	query, _ := NewQuery[T]()
	var tempQuerys []*QueryCond[T]
	tempQuerys = append(tempQuerys, query)
	for i, char := range gcond {
		str := string(char)
		tempQuery := tempQuerys[len(tempQuerys)-1]
		// 如果是 左括号 开头，则代表需要嵌套查询
		if str == "(" && i != len(gcond)-1 {
			if i != 0 && string(gcond[i-1]) == "|" {
				tempQuery.Or(func(q *QueryCond[T]) {
					paramQuery, isOk := queryMaps[string(gcond[i+1])]
					if isOk {
						q.queryExpressions = append(q.queryExpressions, paramQuery.queryExpressions...)
						tempQuerys = append(tempQuerys, q)
					}
				})
				continue
			} else {
				tempQuery.And(func(q *QueryCond[T]) {
					paramQuery, isOk := queryMaps[string(gcond[i+1])]
					if isOk {
						q.queryExpressions = append(q.queryExpressions, paramQuery.queryExpressions...)
						tempQuerys = append(tempQuerys, q)
					}
				})
			}
			continue
		}

		// 如果当前为 | ,而且不是最后一个字符，而且下一个字符不是 ( ,则为 or
		if str == "|" && i != len(gcond)-1 {
			paramQuery, isOk := queryMaps[string(gcond[i+1])]
			if isOk {
				tempQuery.Or().queryExpressions = append(tempQuery.queryExpressions, paramQuery.queryExpressions...)
				tempQuery.last = paramQuery.queryExpressions[len(paramQuery.queryExpressions)-1]
			}
			continue
		}

		if str == "*" && i != len(gcond)-1 {
			paramQuery, isOk := queryMaps[string(gcond[i+1])]
			if isOk {
				tempQuery.And()
				tempQuery.queryExpressions = append(tempQuery.queryExpressions, paramQuery.queryExpressions...)
				tempQuery.last = paramQuery.queryExpressions[len(paramQuery.queryExpressions)-1]
			}
			continue
		}

		if str == ")" {
			// 删除最后一个query对象
			tempQuerys = tempQuerys[:len(tempQuerys)-1]
		}
	}
	return query
}

func getColumnTypeMap[T any]() map[string]reflect.Type {
	var columnNameMap = make(map[string]reflect.Type)
	typeOf := reflect.TypeOf((*T)(nil)).Elem()
	for i := 0; i < typeOf.NumField(); i++ {
		field := typeOf.Field(i)
		columnName := parseColumnName(field)
		columnNameMap[columnName] = field.Type
	}
	return columnNameMap
}
