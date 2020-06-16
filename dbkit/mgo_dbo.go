package dbkit

import (
	"gopkg.in/mgo.v2"
	"fmt"
	"github/SummerCedrus/ServerKit/misc"
	"gopkg.in/mgo.v2/bson"
	"errors"
)

const(
	EVENTUAL = 0
	STRONG = 1
	MONOTONIC = 2
)

type MgoDBO struct {
	Name string
	Info *mgo.DialInfo
	Session *mgo.Session
}

func NewDbo(SessionName string, mode int, hosts []string, databaseName string, username string, password string) *MgoDBO{
	info := &mgo.DialInfo{
		Addrs:hosts,
		Database:databaseName,
		Username:username,
		Password:password,
	}
	session, _ := connectDB(info)

	dbo := &MgoDBO{
		Name:SessionName,
		Info:info,
		Session:session,
	}

	SessionMgr.Add(SessionName, dbo)
	dbo.SetMode(mode, true)
	return dbo
}

func connectDB(info *mgo.DialInfo) (*mgo.Session, error){
	session, err := mgo.DialWithInfo(info)
	if nil != err{
		msg := fmt.Sprintf("conncet addr[%s] failed", info.Addrs)
		panic(msg)
		return nil, err
	}
	return session, err
}
func (dbo *MgoDBO) SetMode(mode int, refresh bool) {
	if !dbo.checkSession(){
		return
	}
	m := mgo.Strong
	switch mode {
	case EVENTUAL:
		m = mgo.Eventual
	case MONOTONIC:
		m = mgo.Monotonic
	}
	dbo.Session.SetMode(m, refresh)
}

func (dbo *MgoDBO) Insert(dbName string, colName string, docs interface{}) bool{
	if !dbo.checkSession(){
		return false
	}
	session := dbo.Session.Clone()
	defer session.Close()
	err := session.DB(dbName).C(colName).Insert(docs)
	if nil != err{
		misc.Errorf("Insert [%s] [%s] error[%s]", dbName, colName, err.Error())
		return false
	}

	return true
}
//Update One Record Match The Condtion
func (dbo *MgoDBO) UpdateOne(dbName string, colName string, cond,docs interface{}) bool{
	if !dbo.checkSession(){
		return false
	}
	session := dbo.Session.Clone()
	defer session.Close()
	err := session.DB(dbName).C(colName).Update(cond,bson.M{"$set":docs})
	if nil != err{
		misc.Errorf("UpdateOne [%s] [%s] error[%s]", dbName, colName, err.Error())
		return false
	}

	return true
}

//Update All Record Match The Condtion
func (dbo *MgoDBO) UpdateAll(dbName string, colName string, cond, docs interface{}) bool{
	if !dbo.checkSession(){
		return false
	}
	session := dbo.Session.Clone()
	defer session.Close()
	chgInfo ,err := session.DB(dbName).C(colName).UpdateAll(cond,bson.M{"$set":docs})
	if nil != err{
		misc.Errorf("UpdateAll [%s] [%s] error[%s]", dbName, colName, err.Error())
		return false
	}
	misc.Logf("UpdateAll Updated[%d]Removed[%d]Matched[%d]", chgInfo.Updated, chgInfo.Removed, chgInfo.Matched)
	return true
}
//Remove One Record Match The Condtion
func (dbo *MgoDBO) Delete(dbName string, colName string, cond interface{}) bool{
	if !dbo.checkSession(){
		return false
	}
	session := dbo.Session.Clone()
	defer session.Close()
	err := session.DB(dbName).C(colName).Remove(cond)
	if nil != err{
		misc.Errorf("Delete [%s] [%s] error[%s]", dbName, colName, err.Error())
		return false
	}

	return true
}

//Remove All Record Match The Condtion
func (dbo *MgoDBO) DeleteAll(dbName string, colName string, cond interface{}) bool{
	if !dbo.checkSession(){
		return false
	}
	session := dbo.Session.Clone()
	defer session.Close()
	chgInfo, err := session.DB(dbName).C(colName).RemoveAll(cond)
	if nil != err{
		misc.Errorf("Update [%s] [%s] error[%s]", dbName, colName, err.Error())
		return false
	}

	misc.Logf("DeleteAll Updated[%d]Removed[%d]Matched[%d]", chgInfo.Updated, chgInfo.Removed, chgInfo.Matched)

	return true
}

func (dbo *MgoDBO) Upsert(dbName string, colName string, cond, docs interface{}) bool{
	if !dbo.checkSession(){
		return false
	}
	session := dbo.Session.Clone()
	defer session.Close()
	chgInfo, err := session.DB(dbName).C(colName).Upsert(cond, bson.M{"$set":docs})
	if nil != err{
		misc.Errorf("Update [%s] [%s] error[%s]", dbName, colName, err.Error())
		return false
	}
	misc.Logf("Upsert Updated[%d]Removed[%d]Matched[%d]", chgInfo.Updated, chgInfo.Removed, chgInfo.Matched)
	return true
}

func (dbo *MgoDBO) FindOne(dbName string, colName string, cond interface{}, result bson.M) error{
	if !dbo.checkSession(){
		return errors.New("Find failed")
	}
	session := dbo.Session.Clone()
	defer session.Close()
	q := session.DB(dbName).C(colName).Find(cond)

	return q.One(result)
}

func (dbo *MgoDBO) FindAll(dbName string, colName string, cond interface{}, handle func(m bson.M) error) error{
	if !dbo.checkSession(){
		return errors.New("Session is invaild")
	}
	session := dbo.Session.Clone()
	defer session.Close()
	q := session.DB(dbName).C(colName).Find(cond)
	iter := q.Iter()
	b := make(bson.M)
	for true == iter.Next(b){
		err := handle(b)
		if nil != err{
			return err
		}
	}

	return nil
}
func (dbo *MgoDBO)checkSession() bool{
	if nil  == dbo.Session{
		misc.Errorf("[%s]Session is invaild!!!", dbo.Name)
		return false
	}
	return true
}