package data

import (
	"database/sql"
	"time"
)

func TestNew(dbPool *sql.DB) Models {
	db = dbPool

	return Models{}
}

var _ UserInterface = UserTest{}

type UserTest struct {
	ID        int
	Email     string
	FirstName string
	LastName  string
	Password  string
	Active    int
	IsAdmin   int
	CreatedAt time.Time
	UpdatedAt time.Time
	Plan      *Plan
}

// Delete implements UserInterface.
func (u UserTest) Delete() error {
	return nil
}

// DeleteByID implements UserInterface.
func (u UserTest) DeleteByID(id int) error {
	return nil
}

// GetAll implements UserInterface.
func (u UserTest) GetAll() ([]*User, error) {
	var users []*User = []*User{
		{
			ID:        0,
			Email:     "admin@exmaple.com",
			FirstName: "First",
			LastName:  "Last",
			Password:  "abc",
			Active:    1,
			IsAdmin:   1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	return users, nil
}

// GetByEmail implements UserInterface.
func (u UserTest) GetByEmail(email string) (*User, error) {
	user := &User{
		ID:        0,
		Email:     "admin@exmaple.com",
		FirstName: "First",
		LastName:  "Last",
		Password:  "abc",
		Active:    1,
		IsAdmin:   1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return user, nil
}

// GetOne implements UserInterface.
func (u UserTest) GetOne(id int) (*User, error) {
	user := &User{
		ID:        0,
		Email:     "admin@exmaple.com",
		FirstName: "First",
		LastName:  "Last",
		Password:  "abc",
		Active:    1,
		IsAdmin:   1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return user, nil
}

// Insert implements UserInterface.
func (u UserTest) Insert(user User) (int, error) {
	return 0, nil
}

// PasswordMatches implements UserInterface.
func (u UserTest) PasswordMatches(plainText string) (bool, error) {
	return true, nil
}

// ResetPassword implements UserInterface.
func (u UserTest) ResetPassword(password string) error {
	return nil
}

// Update implements UserInterface.
func (u UserTest) Update() error {
	return nil
}
