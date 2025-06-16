package utils

import "fmt"

func PGDSN(DBName, Host, Port, User, Password string) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", Host, Port, User, Password, DBName)
}
