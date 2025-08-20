package user

type User struct {
	ID    int
	Name  string
	Email string
}

func New(id int, name string, email string) User {
	return User{
		ID:    id,
		Name:  name,
		Email: email,
	}
}
