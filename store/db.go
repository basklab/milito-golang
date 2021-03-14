package store

import "milito-golang/entity"

type Db struct {
	internal entity.GameState
}

func NewDb(internal entity.GameState) *Db {
	return &Db{internal: internal}
}

func (db *Db) Internal() entity.GameState {
	return db.internal
}

func (db *Db) SetInternal(internal entity.GameState) {
	db.internal = internal
}
