package database

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

type Table interface {
	getDefaultData() []interface{}
	getTableStruct() interface{}
}
type TableBase struct {
	//unUse int `sql:"-"`
}

func CreateTable(db *gorm.DB, table interface{}) error {
	fmt.Println("CreateTable", table)
	db.DropTable(table)
	db.AutoMigrate(table)
	if !db.HasTable(table) {
		return errors.New("Failure to create table!")
	}
	return nil
}

func (t TableBase) getDefaultData() []interface{} {
	log.Println("TableBase getDefaultData")
	return nil
}
