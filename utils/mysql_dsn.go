package utils

import "fmt"

// MysqlDSN generates MySQL connection string
func MysqlDSN(DBName, Host, Port, User, Password string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4&loc=Local", User, Password, Host, Port, DBName)

}
