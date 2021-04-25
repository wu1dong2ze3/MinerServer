package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"reflect"
	"sync"
	"syscall"
)

const dbname string = "miner.db"

type DBM struct{}

var dbm *DBM
var onceDBM sync.Once

func GetInstance() *DBM {
	onceDBM.Do(func() {
		dbm = &DBM{}
	})
	return dbm
}

var instance *gorm.DB = nil
var err error = nil

func (dbm DBM) DB() *gorm.DB {
	if instance == nil {
		bFirst := false
		dbfile, ferr := os.OpenFile(dbname, syscall.O_RDWR, 0755)
		if ferr != nil {
			bFirst = true
		} else {
			bFirst = false
			err = dbfile.Close()
		}
		instance, err = gorm.Open("sqlite3", dbname)
		if err != nil {
			log.Fatalln("Failure to create database! exist!")
		}
		//添加数据库表名前缀
		gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
			return "miner_" + defaultTableName
		}
		if bFirst {
			err = firstStart(instance)
			if err != nil {
				log.Fatalln(err, ":The first attempt to create the database failed! exist!")
			}
		}
		//TODO test
		{
			//instance.Close()
			//os.Remove(dbname)
		}
	}
	return instance
}

func (dbm DBM) Close() {
	if instance != nil {
		err = instance.Close()
		if err != nil {
			log.Println("datebase close error!")
		}
		instance = nil
	}
}

func firstStart(db *gorm.DB) error {
	allModels := []Table{&User{}, &TestData{}}
	var err error = nil
	for _, v := range allModels {
		if err = CreateTable(db, v); err != nil {
			return err
		}
		raw := v.getDefaultData()
		if raw != nil {
			for _, v := range raw {
				log.Println("getDefaultData:", v, reflect.TypeOf(v))
				if err = db.Create(v).Error; err != nil {
					return err
				}
			}
		}
	}
	return err
}
