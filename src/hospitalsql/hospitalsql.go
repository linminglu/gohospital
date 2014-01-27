package hospitalsql

import (
	"container/list"
	"database/sql"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
	//"github.com/go-sql-driver/mysql"
)

const dbmanager string = "root"
const dbpasswd string = "5256120"
const dbip string = "192.168.5.171"
const dbconnectbaseCount int = 10
const dbconnectbaseMax int = 30
const dbname string = "hospital"
const dbconnectaddstep int = 10

var gGameList *list.List // = nilß
var mutex = &sync.Mutex

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

	}
}
func GetInstance() {
	if gGameList == nil {
		OnstartServer()
	}
	if gGameList.Len() == 0 {
		for i := 0; i < dbconnectbaseCount; i++ {
			appendOnSqlDb()
		}
		go startTimerToCloseDatabase()
	}

}

func GetPoolOpenDatabasePoint() (*DbConnectStruct, error) {
	//找到状态为1的 db指针 然后给 使用者
	for e := gGameList.Front(); e != nil; e = e.Next() {
		if e.Value(*DbConnectStruct).GetStatus() == 0 {
			mutex.Lock()
			result := e.Value(*DbConnectStruct)
			gGameList.MoveToBack(e)
			mutex.Unlock()
			return result
		}
	}
	if gGameList.Len() >= dbconnectbaseMax {
		return nil, errors.New("连接池不够用了")
	} else {
		//如果连接池<4个的话  就删除。
		var tem int = 1
		while(tem < dbconnectaddstep)
		{
			appendOnSqlDb()
			tem++
		}

	}
}

func  testTimer1()
{
	var isopenandunsedcout int =0
	mutex.Lock()
	for e := gGameList.Front()	; e!=nil ; e =e.Next() {
		if e.Value(*DbConnectStruct).GetStatus()==0 {
			isopenandunsedcout++;
		}
	}
	if isopenandunsedcout>10 && gGameList.Len()>dbconnectbaseCount {
		//执行关闭数据库并且删除
		var  temp int =0;
		for e := gGameList.Front()	; e!=nil ; e =e.Next() {
		if e.Value(*DbConnectStruct).GetStatus()==0 {
			temp++;
			if temp>dbconnectaddstep {
				break;
			}
			gGameList.Remove(e);
		}
	}
	}
	mutex.Unlock()
}

func  startTimerToCloseDatabase()
{
	//每隔1分钟 检查 连接池的连接无用 如果超过的话 关闭连接池

	timer1 := time.NewTicker(60*time.Second)
	for  {
		select {
		case < -timer1.C:
			testTimer1();

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
