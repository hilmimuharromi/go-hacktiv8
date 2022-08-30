package service

type User struct {
	Nama string
}

type dbUser struct {
	db []*User
}

type UserIface interface {
	Register(u *User) string
	GetUser() []*User
}

func NewUserService(db []*User) UserIface {
	return &dbUser{
		db: db,
	}
}

func (u *dbUser) Register(user *User) string {
	u.db = append(u.db, user)
	return user.Nama
}

func (u *dbUser) GetUser() []*User {
	return u.db
}
