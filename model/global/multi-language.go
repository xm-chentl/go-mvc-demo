package global

type EnumItem struct {
	Key   string `bson:"key"`
	Value string ``
}

type MultiLanguage struct {
	ID    string     `bson:"_id"` // Key
	Items []EnumItem `bson:"items"`
}
