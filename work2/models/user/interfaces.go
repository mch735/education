package user

type Repository interface {
	Save(user User) error
	FindByID(id string) (User, error)
	FindAll() []User
	DeleteByID(id string) error
	FilterFunc(fun func(user User) bool) []User
}
