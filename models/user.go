package models

type UserProfile struct {
	Lastname   string `json:"lastname"`
	Firstname  string `json:"firstname"`
	Midllename string `json:"midllename"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
	Payment    string `json:"payment"`
}
