package api

import (
	"com.arturs.shopping-cart/db"
)

func main() {
	database := db.Database{}
	err := database.Connect(GetDatabaseURI())
	if err != nil {
		panic(err)
	}
	defer database.Disconnect()
}
