package model

import (
	"main/model/db"
)

const MaxSliceCapacity = 100

type Sth interface {
	db.User | db.Template | db.ProposalInfo | db.Tag | db.Question | db.Game
	TableName() string
}

// GetSth is used to get something from database.
// The "something" must fulfil interface "sth", which has method "TableName" and "GetKey".
func GetSth[T Sth](value T) T {
	db.DB.Find(&value)
	return value
}

// UpdateSth The primary key must not be empty
func UpdateSth[T Sth](value T) error {
	result := db.DB.Updates(value)
	return result.Error
}

// CreateSth The primary key must be 0.
func CreateSth[T Sth](value T) (error, T) {
	var x = new(T) // <- Used to fix nil pointer panic.
	*x = value     //	Don't touch it.
	result := db.DB.Create(x)
	return result.Error, *x
}

// DeleteSth The primary key must not be empty
func DeleteSth[T Sth](value T) error {
	result := db.DB.Delete(&value)
	return result.Error
}

// GetManySth Set all the conditions you need to value
// return data slice, number
func GetManySth[T Sth](value T) ([]T, int) {
	data := make([]T, MaxSliceCapacity)
	result := db.DB.Find(&data, value)
	return data[0:int(result.RowsAffected)], int(result.RowsAffected)
}
