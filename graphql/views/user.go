package views

type RequestUser struct {
	Name string
}

type ResponseUser struct {
	RID   string
	RName string
}

func (r ResponseUser) Id() *string {
	return &r.RID
}

func (r ResponseUser) Name() *string {
	return &r.RName
}
