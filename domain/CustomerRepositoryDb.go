package domain

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDb struct {
}

func (d CustomerRepositoryDb) FindAll([]Customer, error) {
	client, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	FindAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
	rows, err := client.Query(FindAllSql)
	if err != nil {
		log.Println("Error while querying customer table" + err.Error())
	}
}
