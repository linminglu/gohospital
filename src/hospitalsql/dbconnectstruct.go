package hospitalsql

import (
	"database/sql"
)

/**
 * 定义Person类
 */
type DbConnectStruct struct {
	IsOpen bool
	Db     *sql.DB
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

func (this *DbConnectStruct) SetIsOpen(value bool) {
	this.IsOpen = value
}

func (this *DbConnectStruct) GetIsOpen() bool {
	return this.IsOpen
}
