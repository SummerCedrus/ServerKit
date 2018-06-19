package dbkit

import (
	"labix.org/v2/mgo"
	"fmt"
)
const(
	DEFAULT_MGO_DB = "127.0.0.1:27017"
)

const(
	EVENTUAL = 0
	STRONG = 1
	MONOTONIC = 2
)

type MgoDBO struct {
	session *mgo.Session
}
func NewSession() *MgoDBO{
	session, _ := connectDB(DEFAULT_MGO_DB)

	dbo := &MgoDBO{
		session:session,
	}

	return dbo
}
func NewSessionWithURL(addr string) *MgoDBO{
	session, _ := connectDB(addr)

	dbo := &MgoDBO{
		session:session,
	}

	return dbo
}

func connectDB(addr string) (*mgo.Session, error){
	session, err := mgo.Dial(addr)
	if nil != err{
		msg := fmt.Sprintf("conncet addr[%s] failed", addr)
		panic(msg)
		return nil, err
	}
	return session, err
}
func (dbo *MgoDBO) SetMode(mode int, refresh bool) {
	m := mgo.Strong
	switch mode {
	case EVENTUAL:
		m = mgo.Eventual
	case MONOTONIC:
		m = mgo.Monotonic
	}
	dbo.session.SetMode(m, refresh)
}

func (dbo *MgoDBO) Insert(){
	dbo.session.Clone()
}