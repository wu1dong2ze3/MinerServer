package database

import (
	"example.com/m/v2/database/utils"
	"fmt"
	"log"
)

type User struct {
	TableBase
	Id   int
	Name string
	Pwd  string
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

func UpdateUser(name string, pwd string, old string) error {
	user := NowUser()

	if !utils.IsLetterAndNumber(user.Name) ||
		!utils.IsLetterAndNumber(user.Pwd) ||
		!utils.Len5_12(user.Name) ||
		!utils.Len5_12(user.Pwd) ||
		user.Pwd != old {
		return UserPwdError
	}
	user.Name = name
	user.Pwd = pwd
	return GetInstance().DB().Save(user).Error
}

func Verify(name string, pwd string) error {
	if err := GetInstance().DB().Find(&User{}, "Name = ? and Pwd = ? ", name, pwd).Error; err != nil {
		log.Println("Verify:", name, pwd)
		return UserPwdInvalid
	}
	return nil
}
