package models

type User struct {
    ID           string `db:"id"`
    Login        string `db:"login"`
    PasswordHash string `db:"password_hash"`
    Role         string `db:"role"`
}