package main

import (
	. "dbkit"
	"fmt"
	"misc"
)

const(
	DEFAULT_MGO_DB = "192.168.1.165:27017"
)

func main(){
	misc.InitLog("run", "dbtest")
	misc.InitLog("error", "dbtest")
	hosts := []string{DEFAULT_MGO_DB}
	dbo := NewDbo("item", MONOTONIC, hosts, "item","","")
	docs := make(map[string]interface{})
	docs["ID"] = 3001
	docs["Amount"] = 10
	ret := dbo.Insert("", "item_0",docs)
	if !ret{
		fmt.Println("insert failed")
		misc.Errorf("insert failed")
		return
	}
	fmt.Println("insert success")
	misc.Logf("insert success")
}