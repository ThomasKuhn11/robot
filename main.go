package main

import (
	"database/sql"
	"fmt"

	"go-gin-api/business"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db := createSQLConnection()

	addNewTemperature(db)

	defer db.Close()

	//temp := business.GetTemperatures() //business.getTempearture
	//fmt.Println(temp)

	//ctx := context.Background()

	//insertIntoDatabase(ctx, temp) //internal.insertIntoDatabase

}

func createSQLConnection() *sql.DB {

	db, err := sql.Open("mysql", "root:my-secret-pw@tcp(127.0.0.1:3308)/temp_db")

	if err != nil {
		panic(err.Error())
	}

	//getTemperature(db)

	//defer db.Close()

	fmt.Println("Go MySQL connected")

	return db
}

func getTemperature(db *sql.DB) {

	type Value struct {
		ID        int    `json:"tempId"`
		Temp      string `json:"temperature"`
		PeriodDay string `json:"period of day"`
		DateSite  string `json:"date_time"`
	}

	result, err := db.Query(fmt.Sprintf("Select * from new_table where tempId= '%v'", 2))
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var value Value

		err := result.Scan(&value.ID, &value.Temp, &value.PeriodDay, &value.DateSite)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(value.ID, value.Temp, value.PeriodDay, value.DateSite)
	}

}

func addNewTemperature(db *sql.DB) {

	temp := business.GetTemperatures().EveningT
	tempWithoutDegree := temp[0:2]

	insert, err := db.Query(fmt.Sprintf("INSERT INTO new_table values (11, %s, 'Evening', current_date()); ", tempWithoutDegree))

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

	fmt.Println("Line INSERTED")

}
