package database

import "fmt"

type TestData struct {
	TableBase
	AAA int
	BBB string
	CCC string
}

func (t TestData) getDefaultData() []interface{} {
	fmt.Println("TableBase getDefaultData")
	return []interface{}{
		TestData{TableBase{}, 11, "BB", "cc"},
		TestData{TableBase{}, 22, "BB22", "cc22"},
	}
}

func (t TestData) getTableStruct() interface{} {
	return &TestData{}
}
