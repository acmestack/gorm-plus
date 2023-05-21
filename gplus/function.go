package gplus

import "github.com/acmestack/gorm-plus/constants"

func As(columnName any, asName any) string {
	return getColumnName(columnName) + " " + constants.As + " " + getColumnName(asName)
}

func SumAs(columnName any, asName any) string {
	return buildFunction(constants.SUM, getColumnName(columnName), getColumnName(asName))
}

func AvgAs(columnName any, asName any) string {
	return buildFunction(constants.AVG, getColumnName(columnName), getColumnName(asName))
}

func MaxAs(columnName any, asName any) string {
	return buildFunction(constants.MAX, getColumnName(columnName), getColumnName(asName))
}

func MinAs(columnName any, asName any) string {
	return buildFunction(constants.MIN, getColumnName(columnName), getColumnName(asName))
}

func CountAs(columnName any, asName any) string {
	return buildFunction(constants.COUNT, getColumnName(columnName), getColumnName(asName))
}

func buildFunction(function string, columnNameStr string, asNameStr string) string {
	columnNameStr = function + constants.LeftBracket + columnNameStr + constants.RightBracket +
		" " + constants.As + " " + asNameStr
	return columnNameStr
}
