package global

type UserValue struct {
	ID   string `alias:"" bson:"_id" db:"_id"`
	Type int64
}

func (u UserValue) GetID() string {
	return u.ID
}
