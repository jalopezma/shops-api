package common

import (
	"log"

	"gopkg.in/mgo.v2"
)

var mgoSession *mgo.Session

// GetMongoSession Creates a new session if mgoSession is nil i.e there is no active mongo session.
// If there is an active mongo session it will return a Clone
func GetMongoSession() *mgo.Session {
	if mgoSession == nil {
		var err error
		// to improve this
		mgoSession, err = mgo.Dial(mongo_conn_str)
		if err != nil {
			log.Fatal("Failed to start the Mongo session")
		}
	}
	return mgoSession.Clone()
}
