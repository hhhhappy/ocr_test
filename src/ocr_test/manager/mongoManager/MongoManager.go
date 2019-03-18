package mongoManager

import (
	"gopkg.in/mgo.v2"
	"ocr_test/manager/configManager"
	"sync"
)

/*
	MongoDB Manager
*/

var (
	mgoSession *mgo.Session
	dataBase string
	once sync.Once
)

/*
	get the mongodb manager
*/
func GetSession() *mgo.Session {
	once.Do(func() {
		var url string
		conf := configManager.GetConf()

		url = conf.Database.Url

		var err error

		mgoSession, err = mgo.Dial(url)
		if err != nil {
			panic(err)
		}

		dataBase = conf.Database.DbName
	})
	return mgoSession.Clone()
}

/*
	collection operator
*/
func WitchCollection(collection string, s func(*mgo.Collection) error) error {
	session := GetSession()
	defer session.Close()
	c := session.DB(dataBase).C(collection)
	return s(c)
}

/*
	collection operator with return info
*/
func WitchCollectionWithInfo(collection string, s func(*mgo.Collection) (info *mgo.ChangeInfo, err error)) (info *mgo.ChangeInfo, err error) {
	session := GetSession()
	defer session.Close()
	c := session.DB(dataBase).C(collection)
	return s(c)
}
