package models

type User struct {
	ID       uint   `db:"id"`
	Email    string `db:"email"`
	Password string `db:"password"`
	Phone    string `db:"phone"`
	Address  string `db:"address"`
}
