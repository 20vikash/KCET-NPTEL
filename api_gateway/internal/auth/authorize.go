package authorize

import (
	"context"
	model "gateway/models/auth"

	"github.com/alexedwards/scs/v2"
)

type Authorize struct {
	Session *scs.SessionManager
}

func (a *Authorize) IsRole(ctx context.Context, role string) (*model.UserSession, bool) {
	role_data := a.Session.Get(ctx, "role").(string)
	user_name_data := a.Session.Get(ctx, "user_name").(string)
	id_data := a.Session.Get(ctx, "id").(string)

	if role_data != role {
		return &model.UserSession{}, false
	}

	return &model.UserSession{UserName: user_name_data, Id: id_data, Role: role_data}, true
}
