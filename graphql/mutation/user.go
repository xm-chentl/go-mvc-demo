package mutation

import (
	"context"
	"fmt"
	"time"

	"github.com/xm-chentl/go-mvc-demo/graphql/views"
)

func (m *Mutation) CreateUser(ctx context.Context, args struct {
	InputArgs views.RequestUser
}) (res *views.ResponseUser, err error) {
	unix := time.Now().Unix()
	res = &views.ResponseUser{
		RID:   fmt.Sprintf("%d", unix),
		RName: fmt.Sprintf("%s 排队 %d", args.InputArgs.Name, unix),
	}

	return
}
