package db

// If you have a new struct,
// please implement:
// func GetKey()
// interface Sth

func GetKey(value sth) (string, int) {
	switch value.(type) {
	case User:
		return "uid", int(value.(User).UID)
	case ProposalInfo:
		return "info_id", int(value.(ProposalInfo).InfoID)
	case Tag:
		return "tid", int(value.(Tag).Tid)
	case Question:
		return "qid", int(value.(Question).Qid)
	case Template:
		return "temid", int(value.(Template).Temid)
	case Game:
		return "gameid", int(value.(Game).Gameid)
	}
	return "", 0
}

type sth interface {
	TableName() string
}

type Sth interface {
	User | Template | ProposalInfo | Tag | Question | Game
	TableName() string
}
