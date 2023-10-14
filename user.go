package connect

import (
	"fmt"

	user "github.com/go-zoox/connect/user"
	"github.com/go-zoox/zoox"
)

func GetUser(ctx *zoox.Context) (u *user.User, err error) {
	v := ctx.User().Get()
	if v == nil {
		return nil, fmt.Errorf("user not found")
	}

	u, ok := v.(*user.User)
	if !ok {
		return nil, fmt.Errorf("user invalid")
	}

	return u, nil
}
