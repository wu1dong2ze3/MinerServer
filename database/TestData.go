package database

import "fmt"

type TestData struct {
	TableBase
	AAA1 int
	BBB2 string
	CCC3 string
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
