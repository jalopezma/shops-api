package common

import (
	"log"

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
func ConnectMongo(url string, db string) *mgo.Session {
	var err error
	// to improve this
	log.Printf("[ConnectMongo] url: %q", url)

	mgoSession = mgoData{}
	mgoSession.dialInfo, err = mgo.ParseURL(url)
	if err != nil {
		log.Printf("[ConnectMongo] Couldn't parse mongodb url %q\n%v", url, err)
	}

	mgoSession.session, err = mgo.Dial(url)
	if err != nil {
		log.Fatal("[ConnectMongo] Failed to start the Mongo session")
	}

	mgoSession.db = mgoSession.session.DB(db)
	//log.Printf("\n%+v\n%+v", mgoSession.dialInfo, mgoSession.db)
	return mgoSession.session.Clone()
}
