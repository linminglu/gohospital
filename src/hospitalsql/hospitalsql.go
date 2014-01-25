package hospitalsql

import (
	"container/list"
	"database/sql"
	"github.com/go-sql-driver/mysql"
)

const dbmanager string = "root"
const dbpasswd string = "5256120"
const dbip string = "192.168.1.110"
const dbconnectbaseCount int = 10
const dbconnectbaseMax int = 30

var gGameList *list

func OnstartServer() {
	gGameList = list.New()
}

func GetInstance() {
	if gGameList == nil {
		OnstartServer()
	}
	if gGameList.Len() == 0 {
		for i := 0; i < dbconnectbaseCount; i++ {
			db, err := sql.Open("mysql", "root:5256120@tcp(192.168.1.110:3306)/hospital?charset=utf8")
			if err == nil {
				dbstuct := new(DbConnectStruct)
				dbstuct.SetDb(db)
				dbstuct.SetIsOpen(false)

			}
		}
	}

}

func GetPoolOpenDatabasePoint() *DbConnectStruct {
	return nil
}

func CheckIsHaveUserById(UserId int) bool {
	return false
}
