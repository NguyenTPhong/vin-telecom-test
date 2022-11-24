package config

var (
	DbConnStr   = GetString("DB_CONN_STR", "host=telecomdb port=5432 user=dbUser password=dbPassword dbname=telecom sslmode=disable")
	LogFilePath = GetString("LOG_FILE_PATH", "api.log")
	Domain      = GetString("DOMAIN", "127.0.0.1")
)
