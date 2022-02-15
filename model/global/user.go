package global

type User struct {
	ID   string `gorm:"column:id"`
	Name string `gorm:"column:name"`
	Age  int    `gorm:"column:age"`
}

func (m User) GetID() string {
	return m.ID
}
