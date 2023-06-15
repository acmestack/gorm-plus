package gplux

type SqlSegment interface {
	getSqlSegment() string
}

type columnPointer struct {
	column any
}

func (cp *columnPointer) getSqlSegment() string {
	// todo 通过反射获取字段名
	return cp.column.(string)
}

type sqlKeyword struct {
	keyword string
}

func (sk *sqlKeyword) getSqlSegment() string {
	return sk.keyword
}

type columnValue struct {
	value any
}

func (cv *columnValue) getSqlSegment() string {
	return ""
}
