package users

import (
	"context"
	"gorm.io/gorm"
	"log"
)

// User -  a representation of the users of the wallet engine
type User struct {
	gorm.Model `json:"-"`
	Username   string `json:"username"`  // username for the user
	Email      string `json:"email"`     // email address for the user
	IsActive   bool   `json:"is_active"` // status of the user, true means active
}

type UserStore interface {
	CreateUser(context.Context, *User) error
	GetUserByID(context.Context, int64) (User, error)
	UpdateUser(context.Context, User, uint) error
	DeleteUser(context.Context, uint) error
}

// UserService is the blueprint for the user logic
type UserService struct {
	Store UserStore
}

// NewService creates a new service
func NewService(store UserStore) UserService {
	return UserService{
		Store: store,
	}
}

func (u *UserService) CreateUser(ctx context.Context, user *User) error {
	if err := u.Store.CreateUser(ctx, user); err != nil {
		log.Printf("Error creating user: %v", err)
		return err
	}
	return nil
}

func (u *UserService) GetUserByID(ctx context.Context, id int64) (User, error) {
	user, err := u.Store.GetUserByID(ctx, id)
	if err != nil {
		log.Printf("Error fetching user with ID %v: %v", id, err)
		return user, err
	}
	return user, nil
}

func (u *UserService) UpdateUser(ctx context.Context, user User, id uint) error {
	if err := u.Store.UpdateUser(ctx, user, id); err != nil {
		log.Printf("Error updating user: %v", err)
		return err
	}

	return nil
}

func (u *UserService) DeleteUser(ctx context.Context, id uint) error {
	if err := u.Store.DeleteUser(ctx, id); err != nil {
		log.Printf("Error updating user: %v", err)
		return err
	}

	return nil
}
