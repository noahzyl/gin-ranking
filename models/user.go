/*
 * Define database operations of users
 */

package models

import "github.com/noahzyl/gin-ranking/dao"

type User struct {
	Id       int
	Username string
}

// Set the table name which gorm will use
func (User) TableName() string {
	// gorm will use "user" as the specific table name when functions in models/user.go are executed
	return "user"
}

// Get a user in database by id
func GetUser(id int) (User, error) {
	var user User
	err := dao.DB.Where("id = ?", id).First(&user).Error // If something is wrong, get an error
	return user, err
}
