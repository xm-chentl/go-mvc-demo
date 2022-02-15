package query

import (
	"fmt"

	"github.com/xm-chentl/go-mvc-demo/graphql/views"
	"github.com/xm-chentl/go-mvc-demo/model/global"
)

func (q *Query) GetUserName() (string, error) {
	name := "chentenglong"
	return name, nil
}

func (q *Query) UserList() (*[]views.ResponseUser, error) {
	var entries []global.User
	if err := q.DbFactory.Db().Query(global.User{}).ToArray(&entries); err != nil {
		return nil, err
	}

	fmt.Println(entries)
	var respEntries []views.ResponseUser
	for _, entry := range entries {
		respEntries = append(respEntries, views.ResponseUser{
			RID:   entry.ID,
			RName: entry.Name,
		})
	}

	return &respEntries, nil
}
