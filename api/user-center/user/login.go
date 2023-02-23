package user

import "fmt"

type LoginAPI struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

func (a LoginAPI) Call() (res interface{}, err error) {
	fmt.Println(a.Account, " | ", a.Password)
	return
}
