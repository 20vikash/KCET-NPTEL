package authorize

import (
	"context"

	"github.com/alexedwards/scs/v2"
)

type Authorize struct {
	Session *scs.SessionManager
}

func (a *Authorize) IsAuthenticated(ctx context.Context) bool {
	id_data := a.Session.GetInt(ctx, "id")

	return id_data != 0
}

func (a *Authorize) IsRole(ctx context.Context, role string) bool {
	role_data := a.Session.GetString(ctx, "role")

	return role_data == role
}

func (a *Authorize) GetId(ctx context.Context) int {
	return a.Session.GetInt(ctx, "id")
}

func (a *Authorize) GetUserName(ctx context.Context) string {
	return a.Session.GetString(ctx, "user_name")
}
