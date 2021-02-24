package models

type QuerySQLResult struct {
	Rows []interface{} // contents depend on used sql query
}
