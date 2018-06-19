package common

import (
	"crypto/tls"
	"log"
	"net"
	"time"

	"gopkg.in/mgo.v2"
)

type mgoData struct {
	session  *mgo.Session
	dialInfo *mgo.DialInfo
	db       *mgo.Database
}

var mgoSession mgoData

// GetMongoSession Creates a new session if mgoSession is nil i.e there is no active mongo session.
// If there is an active mongo session it will return a Clone
func GetMongoSession() (session *mgo.Session) {
	if mgoSession.session != nil {
		session = mgoSession.session.Clone()
	}
	return
}

// ConnectMongo - Connects mongo and returns a clone of the session
func ConnectMongo(hosts []string, dbName string, user string, password string, replicaSet string, ssl bool, authSource string) *mgo.Session {

	var err error

	mgoSession = mgoData{}
	mgoSession.dialInfo = &mgo.DialInfo{
		Addrs:          hosts,
		Database:       dbName,
		Username:       user,
		Password:       password,
		ReplicaSetName: replicaSet,
		Source:         authSource,
		DialServer: func(addr *mgo.ServerAddr) (net.Conn, error) {
			return tls.Dial("tcp", addr.String(), &tls.Config{})
		},
		Timeout: time.Second * 10,
	}

	mgoSession.session, err = mgo.DialWithInfo(mgoSession.dialInfo)
	if err != nil {
		log.Fatalf("[ConnectMongo] Failed to start the Mongo session: %v", err.Error())
	}

	mgoSession.db = mgoSession.session.DB(dbName)
	return mgoSession.session.Clone()
}
