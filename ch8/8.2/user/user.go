package user

func New() User {
	return User{}
}

type User struct {
	Authenticated bool
}
