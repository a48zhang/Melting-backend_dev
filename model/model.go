package model

import (
	"main/model/db"
)

// GetSth is used to get something from database.
// The "something" must fulfil interface "sth", which has method "TableName" and "GetKey".
func GetSth[T db.Sth](value T) T {
	pk, id := db.GetKey(value)
	db.DB.Find(&value, pk+" = ?", id)
	return value
}

func UpdateSth[T db.Sth](value T) error {
	pk, id := db.GetKey(value)
	result := db.DB.Table(value.TableName()).Updates(value).Where(pk+" = ?", id)
	return result.Error
}

func CreateSth[T db.Sth](value T) (error, int) {
	var x = new(T) // <- Used to fix nil pointer panic.
	*x = value     //	Don't touch it.
	result := db.DB.Table(value.TableName()).Create(x)
	_, id := db.GetKey(*x)
	return result.Error, id
}

func DeleteSth[T db.Sth](value T) error {
	result := db.DB.Table(value.TableName()).Delete(&value)
	return result.Error
}

func GetManySth[T db.Sth](value T) ([]T, int) {
	pk, id := db.GetKey(value)
	data := make([]T, 100)
	result := db.DB.Table(value.TableName()).Where(pk+" = ?", id).Scan(&data)
	return data, int(result.RowsAffected)
}

func DeleteProposal(value db.ProposalInfo) error {
	pk, id := db.GetKey(value)
	result := db.DB.Table(value.TableName()).Where(pk+" = ?", id).Update("uid", 0)
	return result.Error
}

func GetProposals(uid int) ([]db.ProposalInfo, int) {
	data := make([]db.ProposalInfo, 100)
	result := db.DB.Table(db.TableNameProposalInfo).Where("uid = ?", uid).Scan(&data)
	return data[0:int(result.RowsAffected)], int(result.RowsAffected)
}

func GetFromUsers(value db.User) db.User {
	db.DB.Table(db.TableNameUser).Find(&value, "nick_name = ?", value.NickName)
	return value
}

func GetGames(value db.Game) []db.Game {
	data := make([]db.Game, 100)
	db.DB.Table(value.TableName()).Find(&data,
		"venue LIKE ? or time LIKE ? or crowd LIKE ?",
		value.Venue, value.Time, value.Crowd,
	).Limit(99)
	return data
}
