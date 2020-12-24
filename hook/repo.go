package hook

import (
	"github.com/globalsign/mgo"
)

const (
	// DBNAME - Database name
	DBNAME = "GIN-HOOK"
	// DBTABLE todo table
	DBTABLE = "hook"
)

type irepo interface {
	createHookRepo(hook Hook) (Hook, error)
}
type repo struct {
	DBSession *mgo.Session
	DBName    string
	DBTable   string
}

func (r *repo) createHookRepo(hook Hook) (Hook, error) {
	coll := r.DBSession.DB(r.DBName).C(r.DBTable)
	err := coll.Insert(&hook)
	if err != nil {
		return Hook{}, err
	}
	return hook, nil
}
func newRepoService(dbs *mgo.Session) *repo {
	return &repo{
		DBSession: dbs,
		DBName:    DBNAME,
		DBTable:   DBTABLE,
	}
}
