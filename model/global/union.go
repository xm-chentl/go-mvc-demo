package global

type Union struct {
	ID   string `column:"id" pk:""`
	Name string `column:"name"`
}

func (u Union) TableName() string {
	return "unions"
}
