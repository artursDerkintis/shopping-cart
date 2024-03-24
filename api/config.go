package api

import "fmt"

func GetDatabaseURI() string {
	return fmt.Sprintf("mongodb+srv://artursderkintis:%s@cluster0.93gdem2.mongodb.net/", db_password)
}
