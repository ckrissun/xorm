package main

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
	. "github.com/lunny/xorm"
)

type User struct {
	Id   int64
	Name string
}

type LoginInfo struct {
	Id     int64
	IP     string
	UserId int64
	// timestamp should be updated by database, so only allow get from db
	TimeStamp string `xorm:"<-"`
	// assume
	Nonuse int `xorm:"->"`
}

func main() {
	f := "singleMapping.db"
	os.Remove(f)

	Orm, err := NewEngine("sqlite3", f)
	if err != nil {
		fmt.Println(err)
		return
	}
	Orm.ShowSQL = true
	err = Orm.CreateTables(&User{}, &LoginInfo{})
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = Orm.Insert(&User{1, "xlw"}, &LoginInfo{1, "127.0.0.1", 1, "", 23})
	if err != nil {
		fmt.Println(err)
		return
	}

	info := LoginInfo{}
	_, err = Orm.Id(1).Get(&info)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(info)
}
