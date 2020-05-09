package dao

import (
	"fmt"
	"github.com/danieloluwadare/dmessanger/oldstructure/migrations"
	"strings"
)

type ExampleRowStruct struct {
	Col1        string `json:"firstName"`
	Col2        string `json:"lastName"`
	Col3         string `json:"age"`

}

//BulkInsert template to do bulk insert still going to undergo refactoring
//Implement and test
func BulkInsert(unsavedRows []*ExampleRowStruct) error {

	db := migrations.GetDB()
	tx:=db.Begin()

	valueStrings := make([]string, 0, len(unsavedRows))
	valueArgs := make([]interface{}, 0, len(unsavedRows) * 3)
	for _, post := range unsavedRows {
		valueStrings = append(valueStrings, "(?, ?, ?)")
		valueArgs = append(valueArgs, post.Col1)
		valueArgs = append(valueArgs, post.Col2)
		valueArgs = append(valueArgs, post.Col3)
	}
	stmt := fmt.Sprintf("INSERT INTO my_sample_table (column1, column2, column3) VALUES %s",
		strings.Join(valueStrings, ","))


	err := tx.Exec(stmt, valueArgs...).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
