package db

import (
	"os"
    "time"
    "gopkg.in/mgo.v2"
)

type MongoCon struct{
	session *mgo.Session
}
func NewConnection(host string, dbName string) (conn *MongoCon) {
	info := &mgo.DialInfo{
		Addrs: []string{host},
		Timeout: 60 * time.Second,
		Database: dbName,
		Username: os.Getenv("Mongo_user"),
		Password: os.Getenv("Mongo_password"),
	}
	session, err := mgo.DialWithInfo(info)
	if err != nil {
			panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
    conn = &MongoCon{session}
    return conn
}
func (conn *MongoCon) Use(dbName, tableName string) (collection *mgo.Collection) {
    // This returns method that interacts with a specific collection and table
    return conn.session.DB(dbName).C(tableName)
}

// Close handles closing a database connection
func (conn *MongoCon) Close() {
    // This closes the connection
    conn.session.Close()
    return
}