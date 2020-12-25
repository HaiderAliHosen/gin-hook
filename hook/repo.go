package hook

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

const (
	// DBNAME - Database name
	DBNAME = "GIN-HOOK"
	// DBTABLE todo table
	DBTABLE = "hook"
)

type irepo interface {
	createHookRepo(hook Hook) (Hook, error)
	readAllHookRepo() ([]Hook, error)
	readSingleHookRepo(hookID string) (Hook, error)
	deleteSingleHookRepo(hookID string) error
	updateHookRepo(hookID string, hook Hook) (Hook, error)
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

func (r *repo) readAllHookRepo() ([]Hook, error) {
	var hook []Hook
	coll := r.DBSession.DB(r.DBName).C(r.DBTable)
	err := coll.Find(bson.M{}).All(&hook)
	if err != nil {
		return []Hook{}, err
	}
	return hook, nil
}

func (r *repo) readSingleHookRepo(hookID string) (Hook, error) {
	var hook Hook
	coll := r.DBSession.DB(r.DBName).C(r.DBTable)
	err := coll.Find(bson.M{"_id": hookID}).One(&hook)
	if err != nil {
		return Hook{}, err
	}
	return hook, nil
}

func (r *repo) deleteSingleHookRepo(hookID string) error {
	coll := r.DBSession.DB(r.DBName).C(r.DBTable)
	err := coll.Remove(bson.M{"_id": hookID})
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) updateHookRepo(hookID string, hook Hook) (Hook, error) {
	hook.ID = hookID
	coll := r.DBSession.DB(r.DBName).C(r.DBTable)
	selector := bson.M{"_id": hookID}
	err := coll.Update(selector, bson.M{"$set": hook})
	if err != nil {
		return Hook{}, err
	}
	updateHook, err := r.readSingleHookRepo(hookID)
	if err != nil {
		return Hook{}, err
	}
	return updateHook, nil
}

func newRepoService(dbs *mgo.Session) *repo {
	return &repo{
		DBSession: dbs,
		DBName:    DBNAME,
		DBTable:   DBTABLE,
	}
}
