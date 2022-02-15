package global

type Account struct {
	ID        string
	Code      string
	Phone     string
	CreatedOn int64
}

func (m Account) GetID() string {
	return m.ID
}
