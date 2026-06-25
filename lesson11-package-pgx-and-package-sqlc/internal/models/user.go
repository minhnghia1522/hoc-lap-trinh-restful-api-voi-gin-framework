package models

type User struct {
	Id        int    `json:"id"`
	Uuid      string `json:"uuid"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}
