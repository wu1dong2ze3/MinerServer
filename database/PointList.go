package database

import (
	"example.com/m/v2/errs"
	"fmt"
	"log"
	"sync"
)

/**
扩展表最多保存8台矿机的采样点
*/
type PointList struct {
	TableBase
	Id         int     `gorm:"primary_key"`
	Time       int     `gorm:"type:integer(12);not null;"`
	Point1     float64 `gorm:"type:decimal(12,2);not null;"`
	Point2     float64 `gorm:"type:decimal(12,2)"`
	Point3     float64 `gorm:"type:decimal(12,2)"`
	Point4     float64 `gorm:"type:decimal(12,2)"`
	Point5     float64 `gorm:"type:decimal(12,2)"`
	Point6     float64 `gorm:"type:decimal(12,2)"`
	Point7     float64 `gorm:"type:decimal(12,2)"`
	Point8     float64 `gorm:"type:decimal(12,2)"`
	PointTotal float64 `gorm:"type:decimal(12,2);not null;"`
}

func (obj PointList) getDefaultData() []interface{} {
	fmt.Println("User getDefaultData")
	return nil
}
func (obj PointList) getTableStruct() interface{} {
	return &PointList{}
}

var lock = sync.Mutex{}

func SavePoints(time int, points *[]float64) *errs.CodeError {
	log.Println("SavePoints:", time, *points)
	lock.Lock()
	defer lock.Unlock()
	if points == nil || len(*points) < 1 || len(*points) > 8 {
		return PointDatasError.AddByString("SavePoints len(points)==0 || len(points)>8 ")
	}
	points8 := make([]float64, 8)
	total := 0.0
	for i, v := range points8 {
		total = v + total
		if i < len(*points) {
			points8[i] = (*points)[i]
			total = total + points8[i]
		} else {
			points8[i] = 0
		}
	}
	pts := &PointList{TableBase{}, 0, time, points8[0],
		points8[1], points8[2], points8[3], points8[4], points8[5],
		points8[6], points8[7], total}

	if err := GetInstance().DB().Create(pts).Error; err != nil {
		return PointDatasError.Add(err)
	}
	return nil
}

func QueryPoints(sinceTime int) (*[]PointList, error) {
	lock.Lock()
	defer lock.Unlock()
	lists := make([]PointList, 0)
	if err := GetInstance().DB().Where("time >= ?", sinceTime).Find(&lists).Error; err != nil {
		return nil, PointDatasError.AddByString("Find error! ")
	}
	return &lists, nil
}

func DelPoints(untilTime int) (*[]PointList, error) {
	lock.Lock()
	defer lock.Unlock()
	lists := make([]PointList, 0)
	if err := GetInstance().DB().Where("time <= ?", untilTime).Delete(&PointList{}).Error; err != nil {
		return nil, PointDatasError.AddByString("Find error! ")
	}
	return &lists, nil
}
