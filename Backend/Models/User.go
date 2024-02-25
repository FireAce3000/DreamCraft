package Models

// Klasse with 3 fields
type User struct {
	Id       int
	Name     string
	Password string
}

// Constructor
func NewUser(id int, username string, password string) *User {
    return &User{
        Id:       id,
		Name:     username,
		Password: password,
    }
}