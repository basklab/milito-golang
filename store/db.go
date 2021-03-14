package store

import "milito-golang/entities"

type Db struct {
	internal entities.GameState
}

func NewDb(internal entities.GameState) *Db {
	return &Db{internal: internal}
}

func (db *Db) Internal() entities.GameState {
	return db.internal
}

func (db *Db) SetInternal(internal entities.GameState) {
	db.internal = internal
}
