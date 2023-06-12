package gplux

type SqlSegment interface {
	getSqlSegment() string
}

type columnPointer struct {
	column any
}

func (cp *columnPointer) getSqlSegment() string {
	return ""
}

type sqlKeyword struct {
	keyword any
}

func (sk *sqlKeyword) getSqlSegment() string {
	return ""
}

type columnValue struct {
	value any
}

func (cv *columnValue) getSqlSegment() string {
	return ""
}
