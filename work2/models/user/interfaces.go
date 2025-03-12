package user

type Repository interface {
	Save(user User) error
	FindByID(id int) (User, error)
	FindAll() []User
	DeleteByID(id int) error
	FilterFunc(fun func(user User) bool) []User
}
