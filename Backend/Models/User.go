package Models

// Klasse with 5 fields
type User struct {
	Id       int
	Name     string
	Password string
	Salt	 []byte
	Hash 	 string
}

// Constructor
func NewUser(id int, username string, password string, salt []byte, hash string) *User {
    return &User{
        Id:       id,
		Name:     username,
		Password: password,
		Salt:     salt,
		Hash: 	  hash,
    }
}