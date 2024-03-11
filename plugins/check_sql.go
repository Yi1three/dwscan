package plugins

import (
	"regexp"
)

var DBMSErrors = map[string][]*regexp.Regexp{
	"MySQL": {
		regexp.MustCompile(`SQL syntax.*MySQL`),
		regexp.MustCompile(`Warning.*mysql_.*`),
		regexp.MustCompile(`valid MySQL result`),
		regexp.MustCompile(`MySqlClient\.`),
	},
	"PostgreSQL": {
		regexp.MustCompile(`PostgreSQL.*ERROR`),
		regexp.MustCompile(`Warning.*\Wpg_.*`),
		regexp.MustCompile(`valid PostgreSQL result`),
		regexp.MustCompile(`Npgsql\.`),
	},
	"Microsoft SQL Server": {
		regexp.MustCompile(`Driver.* SQL[\-\_\ ]*Server`),
		regexp.MustCompile(`OLE DB.* SQL Server`),
		regexp.MustCompile(`(\W|\A)SQL Server.*Driver`),
		regexp.MustCompile(`Warning.*mssql_.*`),
		regexp.MustCompile(`(\W|\A)SQL Server.*[0-9a-fA-F]{8}`),
		regexp.MustCompile(`(?s)Exception.*\WSystem\.Data\.SqlClient\.`),
		regexp.MustCompile(`(?s)Exception.*\WRoadhouse\.Cms\.`),
	},
	"Microsoft Access": {
		regexp.MustCompile(`Microsoft Access Driver`),
		regexp.MustCompile(`JET Database Engine`),
		regexp.MustCompile(`Access Database Engine`),
	},
	"Oracle": {
		regexp.MustCompile(`\bORA-[0-9][0-9][0-9][0-9]`),
		regexp.MustCompile(`Oracle error`),
		regexp.MustCompile(`Oracle.*Driver`),
		regexp.MustCompile(`Warning.*\Woci_.*`),
		regexp.MustCompile(`Warning.*\Wora_.*`),
	},
	"IBM DB2": {
		regexp.MustCompile(`CLI Driver.*DB2`),
		regexp.MustCompile(`DB2 SQL error`),
		regexp.MustCompile(`\bdb2_\w+\(`),
	},
	"SQLite": {
		regexp.MustCompile(`SQLite/JDBCDriver`),
		regexp.MustCompile(`SQLite.Exception`),
		regexp.MustCompile(`System.Data.SQLite.SQLiteException`),
		regexp.MustCompile(`Warning.*sqlite_.*`),
		regexp.MustCompile(`Warning.*SQLite3::`),
		regexp.MustCompile(`\[SQLITE_ERROR\]`),
	},
	"Sybase": {
		regexp.MustCompile(`(?i)Warning.*sybase.*`),
		regexp.MustCompile(`Sybase message`),
		regexp.MustCompile(`Sybase.*Server message.*`),
	},
}

func Checksql() {

}
