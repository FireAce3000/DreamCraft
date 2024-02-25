package Services

import (
	"DreamCraft/Models"
	"crypto/rand"
    "crypto/sha256"
    "encoding/hex"
	"fmt"
)

var userSlice []Models.User

// CRUD

// Create
func Create(name string, password string) {

	// Create id and check the max id
	var maxId int = 0

	for _, user := range userSlice {
		if user.Id > maxId {
			maxId = user.Id
		}
	}

	var id int = maxId + 1

	// Salt in []byte
	salt, _ := CreateSalt()

	// Hash SHA-265
	hash, _ := hashPassword(password, salt)

	// Add newUser
	newUser := Models.NewUser(id, name, password, salt, hash)

	// Add a newUser to Slice
	userSlice = append(userSlice, *newUser)
}

// Read
func Read() []Models.User {
	return userSlice
}

func CreateSalt() ([]byte, error) {
	// Salt lenght
    var saltLength int = 32

    // Create a slice, to save the salt
    var salt []byte = make([]byte, saltLength)

    // Create a crypto salt
	_, err := rand.Read(salt)
    if err != nil {
        fmt.Println("ERROR:", err)
    }

	// Convert bytes salt to hex String
    // saltString := hex.EncodeToString(salt)
	// fmt.Println("Generatic Salt:", saltString)
	return salt, nil
}

// Hashen from password and salt
func hashPassword(password string, salt []byte) (string, error) {
    // Password + salt
    var saltedPassword []byte = append([]byte(password), salt...)

    // Create SHA-256-Hash string
    hash := sha256.Sum256(saltedPassword)

    // Konvertiere den Hash in eine hexadezimale Zeichenkette
    hashedPassword := hex.EncodeToString(hash[:])

    return hashedPassword, nil
}