package models

import (
	"BetterApp/internal/users"
	"context"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Email    string `gorm:"unique;not null"`
	IsActive bool   `gorm:"not null"`
}

func (d *Database) CreateUser(ctx context.Context, user *users.User) error {
	newUser := &User{
		Username: user.Username,
		Email:    user.Email,
		IsActive: false,
	}

	if err := d.Client.WithContext(ctx).Create(newUser).Error; err != nil {
		return err
	}

	return nil
}

// GetUserByID returns the user with a specified id
func (d *Database) GetUserByID(ctx context.Context, id int64) (users.User, error) {
	user := users.User{}
	if err := d.Client.WithContext(ctx).Where("id = ?", id).First(&user).Error; err != nil {
		return users.User(User{}), err
	}
	return users.User(User{
		Username: user.Username,
		Email:    user.Email,
		IsActive: user.IsActive,
	}), nil
}

// UpdateUser updates an existing user in the database
func (d *Database) UpdateUser(ctx context.Context, updatedUser users.User, id uint) error {
	// Check if the user with the specified ID exists
	var existingUser User
	if err := d.Client.WithContext(ctx).Where("id = ?", id).First(&existingUser).Error; err != nil {
		return err
	}

	// Update the fields of the existing user with the new values
	existingUser.Username = updatedUser.Username
	existingUser.Email = updatedUser.Email
	existingUser.IsActive = updatedUser.IsActive

	// Save the updated user back to the database
	if err := d.Client.WithContext(ctx).Save(&existingUser).Error; err != nil {
		return err
	}

	return nil
}

// DeleteUser deletes a user from the database by their ID

func (d *Database) DeleteUser(ctx context.Context, id uint) error {
	// Check if the user with the specified ID exists
	var existingUser User
	if err := d.Client.WithContext(ctx).Where("id = ?", id).First(&existingUser).Error; err != nil {
		return err
	}

	// Delete the user from the database
	if err := d.Client.WithContext(ctx).Delete(&existingUser).Error; err != nil {
		return err
	}

	return nil
}
