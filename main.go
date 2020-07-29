package main

import (
	"fmt"

	"github.com/eren-dhoheiri/api-pq/config"
)

func main() {
	fmt.Println("Hello")

	db, err := config.GetPostgresDB()

	if err != nil {
		fmt.Println(err)
	}

	err = db.Ping()

	if err != nil {
		fmt.Println(err)
	}

}
