package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"main/model/db"
	"main/model/mongodb"
)

func DeleteProposal(value db.ProposalInfo) error {
	result := db.DB.Table(value.TableName()).Where("info_id = ?", value.InfoID).Update("uid", 0)
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

func GetGameDetail(gameID int) (bson.D, error) {
	col := mongodb.Client.Database("melting").Collection("gameDetails")
	cursor := col.FindOne(context.TODO(), bson.D{{"gameID", gameID}})
	var result bson.D
	err := cursor.Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
