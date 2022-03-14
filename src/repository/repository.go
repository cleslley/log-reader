package repository

import (
	"log-reader/src/db"
	"log-reader/src/model"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const LogsCollection = "logsdb"

type Log interface {
	Insert(log *model.Log) error
	GetAll() (logs []*model.Log, err error)
	DeleteAll() error
}

type logs struct {
	c *mgo.Collection
}

func NewLogRepository(conn db.Connection) Log {
	return &logs{c: conn.DB().C(LogsCollection)}
}

func (l *logs) Insert(log *model.Log) error {
	return l.c.Insert(log)
}

func (l *logs) GetAll() (logs []*model.Log, err error) {
	err = l.c.Find(bson.M{}).All(&logs)
	return logs, err
}

func (r *logs) DeleteAll() error {
	return r.c.DropCollection()
}
