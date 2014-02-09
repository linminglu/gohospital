package hospitalsql

import (
	"container/list"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"sync"
	"time"
)

const dbmanager string = "linminglu"
const dbpasswd string = "5256120"
const dbip string = "db4free.net"
const dbconnectbaseCount int = 10
const dbconnectbaseMax int = 30
const dbname string = "dlinminglu"
const dbconnectaddstep int = 10

var gGameList *list.List // = nilß
var mutex = &sync.Mutex{}

func OnstartServer() {
	gGameList = list.New()
}
func OnServerDown() {
	for e := gGameList.Front(); e != nil; e = e.Next() {

	}
}
func appendOnSqlDb() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", dbmanager, dbpasswd, dbip, dbname))
	if err == nil {
		defer db.Close()
		dbstuct := new(DbConnectStruct)
		dbstuct.SetDb(db)
		dbstuct.SetDbStatus(0)
		gGameList.PushFront(dbstuct)

	} else {
		fmt.Printf("data base error %s\n", err.Error())
	}
}
func GetInstance() {
	fmt.Printf("aaaaaaaa")
	if gGameList == nil {
		OnstartServer()
		fmt.Print("create gamelist ok\n")
	}
	if gGameList.Len() == 0 {
		for i := 0; i < dbconnectbaseCount; i++ {
			fmt.Print("append  one  sqldb1\n")
			appendOnSqlDb()
			fmt.Print("append  one  sqldb2 ok\n")
		}
		go startTimerToCloseDatabase()
	}

}

func GetPoolOpenDatabasePoint() (*DbConnectStruct, error) {
	//找到状态为1的 db指针 然后给 使用者

	if gGameList.Len() >= dbconnectbaseMax {
		return nil, errors.New("连接池不够用了")
	} else {
		//如果连接池<4个的话  就删除。
		var tem int = 1
		for ; tem < dbconnectaddstep; tem++ {
			appendOnSqlDb()
		}

	}
	for e := gGameList.Front(); e != nil; e = e.Next() {
		if e.Value.(*DbConnectStruct).GetDbStatus() == 0 {
			mutex.Lock()
			result := e.Value.(*DbConnectStruct)
			gGameList.MoveToBack(e)
			mutex.Unlock()
			return result, nil
		}
	}
	return nil, errors.New("连接池不够用了")

}

func testTimer1() {
	fmt.Print("start time")
	var isopenandunsedcout int = 0
	mutex.Lock()
	for e := gGameList.Front(); e != nil; e = e.Next() {
		if e.Value.(*DbConnectStruct).GetDbStatus() == 0 {
			isopenandunsedcout++
		}
	}
	if isopenandunsedcout > 10 && gGameList.Len() > dbconnectbaseCount {
		//执行关闭数据库并且删除
		var temp int = 0
		for e := gGameList.Front(); e != nil; e = e.Next() {
			if e.Value.(*DbConnectStruct).GetDbStatus() == 0 {
				temp++
				if temp > dbconnectaddstep {
					break
				}
				// tempData := e.Value.(*DbConnectStruct)
				// Delete(tempData)
				gGameList.Remove(e)
			}
		}
	}
	mutex.Unlock()
}

func startTimerToCloseDatabase() {
	//每隔1分钟 检查 连接池的连接无用 如果超过的话 关闭连接池
	fmt.Print("start time")
	timer1 := time.NewTicker(60 * time.Second)
	for {
		select {
		case <-timer1.C:
			testTimer1()
		}
	}
}
func GetPoolDbUseOk(db *DbConnectStruct) {
	db.SetDbStatus(0)
	for e := gGameList.Front(); e != nil; e = e.Next() {
		if e.Value == db {
			mutex.Lock()
			gGameList.MoveToFront(e)
			mutex.Unlock()
			return
		}
	}
}
