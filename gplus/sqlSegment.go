package gplus

type SqlSegment interface {
	getSqlSegment() string
}

type columnPointer struct {
	column any
}

func (cp *columnPointer) getSqlSegment() string {
	return getColumnName(cp.column)
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
