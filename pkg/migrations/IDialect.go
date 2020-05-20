package migrations

type IDialect interface {
	GetDataBaseConnection(dialect string, username string, password string, dbName string, dbHost string, dbPort string)
}