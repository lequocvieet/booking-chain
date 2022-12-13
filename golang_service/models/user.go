package models

import (
	"fmt"
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func Santize(data string) string {
	data = html.EscapeString(strings.TrimSpace(data))
	return data
}

func NewUser(userName string, email string, password string) User {
	u := User{
		Email:    email,
		UserName: userName,
		Password: password,
	}
	return u
}

func (m *Model) FindUserByUserName(username string) (User, error) {
	var user User
	result := m.DB.Where("user_name = ?", username).First(&user)
	if result.Error != nil {
		return user, fmt.Errorf("userName not exist")
	}
	return user, nil
}

func (m *Model) FindUserByUserNameAndEmail(username string, email string) (User, error) {
	var user User
	err := m.DB.
		Where("user_name = ? and email=?", username, email).
		First(&user).Error
	return user, err
}

func (m *Model) SaveUser(user User) error {
	m.DB.Save(&user)
	return fmt.Errorf("Save user success!")
}
