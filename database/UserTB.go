package database

import (
	"example.com/m/v2/database/utils"
	"fmt"
	"log"
)

type User struct {
	TableBase
	Id   int    `gorm:"primary_key"`
	Name string `gorm:"type:varchar(256);not null;"`
	Pwd  string `gorm:"type:varchar(256);not null;"`
}

func (obj User) getDefaultData() []interface{} {
	fmt.Println("User getDefaultData")
	return []interface{}{&User{TableBase{}, 0, "admin", "admin"}}
}
func (obj User) getTableStruct() interface{} {
	return &User{}
}

func NowUser() *User {
	user := User{}
	err := GetInstance().DB().Find(&user, "Id = ?", 1).Error
	if err != nil {
		log.Fatalln("The current user cannot be null！Exist!", err)
	}
	return &user
}

func UpdateUser(pwd string, old string) (*User, error) {
	user := NowUser()
	if !utils.IsLetterAndNumber(user.Pwd) ||
		!utils.Len5_12(user.Pwd) ||
		user.Pwd != old {
		log.Println("UpdateUser check:", pwd, old, UserPwdError)
		return user, UserPwdError
	}
	user.Pwd = pwd
	if err := GetInstance().DB().Model(user).Update("Pwd", pwd).Error; err != nil {
		log.Println("UpdateUser Save:", pwd, old, user, UserPwdError)
		return nil, err
	}
	return user, nil
}

func Verify(name string, pwd string) error {
	if err := GetInstance().DB().Find(&User{}, "Name = ? and Pwd = ? ", name, pwd).Error; err != nil {
		log.Println("Verify:", name, pwd)
		return UserPwdInvalid
	}
	return nil
}
