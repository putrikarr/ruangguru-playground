package repository

import (
	"fmt"
	//"log"
	"strconv"

	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/cashier-app/db"
)

type UserRepository struct {
	db db.DB
}

func NewUserRepository(db db.DB) UserRepository {
	return UserRepository{db}
}

func (u *UserRepository) LoadOrCreate() ([]User, error) {
	//return []User{}, nil // TODO: replace this
	records, err := u.db.Load("users")
	if err != nil {
		records = [][]string{
			{"username", "password", "loggedin", "role"},
		}
		if err := u.db.Save("users", records); err != nil {
			return nil, err
		}
	}

	result := make([]User, 0)
	for i := 1; i < len(records); i++ {
		loggedIn, err := strconv.ParseBool(records[i][2])
		if err != nil {
			return nil, err
		}

		user := User{
			Username: records[i][0],
			Password: records[i][1],
			Loggedin: loggedIn,
		}
		result = append(result, user)
	}
	return result, nil
}

func (u *UserRepository) SelectAll() ([]User, error) {
	//return []User{}, nil // TODO: replace this
	records, _ := u.db.Load("users")
	result := make([]User, 0)
	var isLoggedIn bool
	for i := 1; i < len(records); i++ {
		if records[i][2] == "false" {
			isLoggedIn = false
		} else {
			isLoggedIn = true
		}
		user := User{
			Username: records[i][0],
			Password: records[i][1],
			Loggedin: isLoggedIn,
			Role:     records[i][3],
		}
		result = append(result, user)
	}
	return result, nil

	// users, err := u.db.SelectAll()
	// if err != nil {
	// 	return nil, err
	// }
	// return users, nil
}

func (u UserRepository) Login(username string, password string) (*string, error) {
	//return nil, nil // TODO: replace this
	records, _ := u.db.Load("users")
	updated := [][]string{
		{"username", "password", "loggedin", "role"},
	}

	var isUsernameAndPasswordFound bool
	var result *string
	for i := 1; i < len(records); i++ {
		records[i][2] = "false"
		if records[i][0] == username && records[i][1] == password {
			records[i][2] = "true"
			isUsernameAndPasswordFound = true
			result = &records[i][3]
		}
		updated = append(updated, []string{
			records[i][0],
			records[i][1],
			records[i][2],
			records[i][3],
		})
	}

	if isUsernameAndPasswordFound == false {
		return nil, fmt.Errorf("Username and password not found")
	}
	err := u.db.Save("users", updated)
	if err != nil {
		return nil, err
	}
	return result, nil

	// if err := u.LogoutAll(); err != nil {
	// 	return nil, err
	// }

	// users, err := u.SelectAll()
	// if err != nil {
	// 	return nil, err
	// }
	// for _, user := range users {
	// 	if user.Username == username && user.Password == password {
	// 		u.changeStatus(username, true)
	// 		return &username, nil
	// 	}
	// }
	// return nil, fmt.Errorf("Login Failed")
}

func (u *UserRepository) Save(users []User) error {
	return nil // TODO: replace this

	// records := [][]string{
	// 	{"username", "password", "loggedin", "role"},
	// }
	// for i := 0; i < len(users); i++ {
	// 	records = append(records, []string{
	// 		users[i].Username,
	// 		users[i].Password,
	// 		strconv.FormatBool(users[i].LoggedIn),
	// 		users[i].Role,
	// 	})
	// }
	// return u.db.Save("users", records)
}

func (u *UserRepository) GetUserRole(username string) (*string, error) {
	// TODO: answer here
	records, _ := u.db.Load("users")
	for i := 1; i < len(records); i++ {
		if records[i][0] == username {
			return &records[i][3], nil
		}
	}
	return nil, fmt.Errorf("User not found")

	// users, err := u.SelectAll()
	// if err != nil {
	// 	return nil, err
	// }

	// for _, user := range users {
	// 	if user.Username == username {
	// 		return &user.Role, nil
	// 	}
	// }
	// return nil, fmt.Errorf("Failed to get user role")
}
