/*
 * Define database operations of users
 */

package models

import "github.com/noahzyl/gin-ranking/dao"

// User struct is consistent with the user table in database
type User struct {
	Id       int // Primary key, assigned by MySQL automatically
	Username string
	Password string
}

// Set the table name which gorm will use
func (User) TableName() string {
	// gorm will use "user" as the specific table name when functions in models/user.go are executed
	return "user"
}

// Add/Create a new user
func AddUser(username string, password string) (int, error) {
	user := User{Username: username, Password: password}
	err := dao.DB.Create(&user).Error
	return user.Id, err
}

// Update a user's name
func UpdateUserName(id int, newUsername string) error {
	err := dao.DB.Model(&User{}).Where("id = ?", id).Update("username", newUsername).Error
	return err
}

// Update a user's password
func UpdateUserPassword(id int, newPassword string) error {
	err := dao.DB.Model(&User{}).Where("id = ?", id).Update("password", newPassword).Error
	return err
}

// Delete a user
func DeleteUser(id int) error {
	err := dao.DB.Delete(&User{}, id).Error
	return err
}

// Get a user in database by id
func GetUser(id int) (User, error) {
	var user User
	err := dao.DB.Where("id = ?", id).First(&user).Error // If something is wrong, get an error
	return user, err
}

// Get users by username
func GetUserByUsername(username string) ([]User, error) {
	var users []User
	err := dao.DB.Where("username = ?", username).Find(&users).Error
	return users, err
}
