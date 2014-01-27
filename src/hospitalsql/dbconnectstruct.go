package hospitalsql

import (
	"database/sql"
)

/**
 * 定义DbConnectStruct类
 */
type DbConnectStruct struct {
	DbStatus int
	Db       *sql.DB
}

/**
 * 为Person类定义方法
 */
func (this *DbConnectStruct) SetDb(db *sql.DB) {
	this.Db = db
}

func (this *DbConnectStruct) GetDb() *sql.DB {
	return this.Db
}

func (this *DbConnectStruct) SetDbStatus(value int) {
	this.DbStatus = value
}

func (this *DbConnectStruct) GetDbStatus() int {
	return this.DbStatus
}
