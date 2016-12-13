package mongo

import (
	"gopkg.in/mgo.v2"
	"sync"
	"time"
)

var (
	globalMS *mgo.Session
	mux sync.RWMutex
)

const (
	DEFAULT_MGO_TIMEOUT=15
)
func init(){
	sess,err:=mgo.Dial("192.168.52.128:27017")
	if(err!=nil){
		panic(err)
	}
	sess.SetMode(mgo.Strong,true)
	sess.SetSocketTimeout(DEFAULT_MGO_TIMEOUT*time.Second)
	sess.SetCursorTimeout(0)
	globalMS=sess
}

func Connect(dataBase,cname string)(*mgo.Session,*mgo.Collection){
	ms:=globalMS.Copy()
	co:=ms.DB(dataBase).C(cname)
	return ms,co
}