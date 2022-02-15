package login

import (
	"fmt"

	"github.com/xm-chentl/go-code/dbfactory"
	"github.com/xm-chentl/go-code/guidex"
)

type AccountAPI struct {
	DbFactory dbfactory.IDbFactory `inject:""`
	Guid      guidex.IGenerate     `inject:""`

	Mobile   string
	Password string
}

func (s AccountAPI) Call() (res interface{}, err error) {
	for i := 0; i < 10; i++ {
		fmt.Println(s.Guid.String())
	}

	return
}
