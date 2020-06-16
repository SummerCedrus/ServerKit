package main
//dbo test
import (
	//."dbkit"
	."github/SummerCedrus/ServerKit/misc"
	"github/SummerCedrus/ServerKit/protocol"
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

type DBItem struct {
	ID int32 `bson:"id,omitempty"`
	Amount int32 `bson:"amount,omitempty"`
	Name string `bson:"name,omitempty"`
}
const(
	DEFAULT_MGO_DB = "192.168.1.165:27017"
)

func main(){
	InitLog("run", "dbtest")
	InitLog("error", "dbtest")
	//hosts := []string{DEFAULT_MGO_DB}
	//dbo := NewDbo("item", MONOTONIC, hosts, "item","","")
	//docs := make(map[string]interface{})
	//docs["ID"] = 3001
	//docs["Amount"] = 10
	//result := make([]bson.M,0)
	//handle := func(m bson.M) error{
	//	result = append(result, m)
	//	return nil
	//}
	//ret := dbo.FindAll("", "item_0", bson.M{"ID":3001}, handle)
	//
	//if nil != ret{
	//	fmt.Println("Upsert failed")
	//	misc.Errorf("Upsert failed")
	//	return
	//}else{
	//	fmt.Println("Upsert success")
	//	misc.Logf("Upsert success")
	//}
	//
	//misc.Log(result)
	item := DBItem{
		ID:3001,
		Amount:12,
		Name:"item_3001",
	}
	out := bson.M{}
	err := Struct2M(item, out)
	if nil != err{
		fmt.Println("error")
		return
	}
	fmt.Println("Struct2M:",out)
	newItem := protocol.ItemInfo{}
	err = M2Struct(out, &newItem)

	if nil != err{
		fmt.Println("error")
		return
	}
	fmt.Println("M2Struct:",newItem.String())
}