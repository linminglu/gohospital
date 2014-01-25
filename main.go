package main

import (
	"container/list"
	"fmt"
	"hospitalsql"
)

func main() {
	items := list.New()
	for i := 0; i < 9; i++ {
		temp := new(hospitalsql.DbConnectStruct)
		temp.SetIsOpen(i % 2)
		items.PushBack(temp)
	}
	var i int = 0
	for e := items.Front(); nil != e; e = e.Next() {
		//fmt.Println("[", i,"]: ",e.Value.(int))
		fmt.Printf("[%d]: %d\n", i, e.Value.(hospitalsql.DbConnectStruct).GetIsOpen())
		i += 1
	}
}
