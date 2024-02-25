package Services

import (
	"DreamCraft/Models"
)

var userSlice []Models.User

// CRUD

// Create
func Create(name string, password string) {

	// Create id and check the max id
	maxId := 0

	for _, user := range userSlice {
		if user.Id > maxId {
			maxId = user.Id
		}
	}

	id := maxId + 1

	// Add newUser
	newUser := Models.NewUser(id, name, password)

	// Add a newUser to Slice
	userSlice = append(userSlice, *newUser)
}

// Read
func Read() []Models.User {
	return userSlice
}
