package model

import "main/model/db"

func DeleteProposal(value db.ProposalInfo) error {
	result := db.DB.Table(value.TableName()).Where("uid = ?", value.UID).Update("uid", 0)
	return result.Error
}

func GetFromUsers(value db.User) db.User {
	db.DB.Table(db.TableNameUser).Find(&value, "nick_name = ?", value.NickName)
	return value
}

func GetGames(value db.Game) []db.Game {
	data := make([]db.Game, MaxSliceCapacity)
	db.DB.Table(value.TableName()).Find(&data,
		"venue LIKE ? or time LIKE ? or crowd LIKE ?",
		value.Venue, value.Time, value.Crowd,
	).Limit(99)
	return data
}

func GetTemplate(name string) db.Template {
	data := db.Template{Name: name}
	db.DB.Table(db.TableNameTemplate).Find(&data)
	return data
}
