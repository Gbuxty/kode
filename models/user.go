package models

import (
    "golang.org/x/crypto/bcrypt"
)

type User struct {
    ID       int    `json:"id"`
    Username string `json:"username"`
    Password string `json:"-"`
}

func (u *User) HashPassword(password string) error {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    if err != nil {
        return err
    }
    u.Password = string(bytes)
    return nil
}

func (u *User) CheckPassword(providedPassword string) error {
    err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(providedPassword))
    if err != nil {
        return err
    }
    return nil
}


