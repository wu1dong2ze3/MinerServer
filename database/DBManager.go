package database

import (
	"example.com/m/v2/shell"
	"example.com/m/v2/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"reflect"
	"sync"
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
		path := getDBFilePath()
		log.Println("DB path=", path)
		bFirst := false
		dbfile, ferr := os.OpenFile(path, os.O_RDONLY, 0755)
		if ferr != nil {
			bFirst = true
			log.Println("DB open path is ", ferr)
		} else {
			bFirst = false
			err = dbfile.Close()
		}
		instance, err = gorm.Open("sqlite3", path)
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

func (dbm DBM) Delete() {
	instance.Close()
	shell.RM.Params("-f " + getDBFilePath()).Exec()
}

func (dbm DBM) Close() {
	if instance != nil {
		err = instance.Close()
		if err != nil {
			log.Println("datebase close error!", err)
		}
		instance = nil
	}
}

func firstStart(db *gorm.DB) error {
	log.Println("DB firstStart")
	allModels := []Table{&User{}, &TestData{}, &PointList{}}
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

func getDBFilePath() string {
	path := utils.NowPath()
	if path != "" {
		path = path + "/" + dbname
	} else {
		path = dbname
	}
	return path
}
