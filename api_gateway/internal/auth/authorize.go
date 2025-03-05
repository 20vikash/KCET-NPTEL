package authorize

import "github.com/alexedwards/scs/v2"

type Authorize struct {
	Session *scs.SessionManager
}

func (a *Authorize) IsStudent() {

}

func (a *Authorize) IsTeacher() {

}

func (a *Authorize) IsAdmin() {

}
