package storage

type persistence interface {
	GetUser()
	PostUser()
	DeleteUser()

}
